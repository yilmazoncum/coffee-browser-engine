package main

import (
	"flag"
	"log"
	"main/engine"
	"main/parsers"
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

	// fmt.Println("~~~~ HTML ~~~~")
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
	//engine.PrintRenderTree(renderTree)

	layoutTree := engine.LayoutTree(renderTree)
	engine.PrintLayoutTree(layoutTree)

}
