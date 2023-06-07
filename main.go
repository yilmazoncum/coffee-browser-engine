package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"main/Parsers"
)

func main() {
	htmlFile := flag.String("html", "example.html", "the name of the html to search")

	// Parse the command-line flags
	flag.Parse()

	// Print the value of the flag
	_html, err := ioutil.ReadFile(*htmlFile)
	if err != nil {
		panic(err)
	}

	rootNode, err := Parsers.ParseHtml(string(_html))
	if err != nil {
		panic(err)
	}
	fmt.Printf("#%v\n", rootNode.Name)
	fmt.Print("\n~~~~~ Tree ~~~~~\n")

	// Parsers.PrintTree(rootNode)
}
