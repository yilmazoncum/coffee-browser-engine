package main

import (
	"flag"
	"main/parsers"
	"os"
)

func main() {
	htmlFile := flag.String("html", "example.html", "the name of the html to search")

	// Parse the command-line flags
	flag.Parse()

	// Print the value of the flag
	html, err := os.ReadFile(*htmlFile)
	if err != nil {
		panic(err)
	}

	rootNode, err := parsers.ParseHtml(string(html))
	if err != nil {
		panic(err)
	}
	parsers.PrintTree(rootNode)
}
