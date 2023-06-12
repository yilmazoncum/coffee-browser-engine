package main

import (
	"flag"
	"log"
	"main/engine"
	"main/parsers"
	"main/rasterizer"
	"net/http"
	"os"
)

func main() {
	htmlFile := flag.String("html", "example.html", "the name of the html to search")
	cssFile := flag.String("css", "example.css", "the name of the css to search")

	// Parse the command-line flags
	flag.Parse()

	html, err := os.ReadFile(*htmlFile)
	if err != nil {
		log.Fatal(err)
	}

	rootNode, err := parsers.ParseHtml(string(html))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("\n~~~~ HTML ~~~~")
	// parsers.PrintTree(rootNode)

	css, err := os.ReadFile(*cssFile)
	if err != nil {
		log.Fatal(err)
	}

	stylesheet, err := parsers.ParseCss(string(css))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("\n~~~~ CSS ~~~~")
	// parsers.PrintStyle(stylesheet)

	renderTree := engine.RenderTree(rootNode, stylesheet)
	// fmt.Println("\n~~~~ Render Tree ~~~~")
	// engine.PrintRenderTree(renderTree)

	layoutTree := engine.LayoutTree(renderTree)
	// fmt.Println("\n~~~~ Layout Tree ~~~~")
	// engine.PrintLayoutTree(layoutTree)

	img := rasterizer.CreateCanvas()
	img = rasterizer.Paint(img, layoutTree)
	rasterizer.CreateImage(img)

	//serve

	imageData, err := os.ReadFile("image.png")
	if err != nil {
		log.Fatal(err)
	}

	// Create an HTTP handler function to serve the image
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		_, err := w.Write(imageData)
		if err != nil {
			log.Println(err)
		}
	})

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
