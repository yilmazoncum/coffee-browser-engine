package Parsers

import (
	"fmt"
	"strings"
)

func ParseHtml(html string) (*Node, error) {
	var stack Stack
	stack.Push(&Node{
		Name: "root",
	})
	for i := 0; i < len(html); i++ {
		fmt.Printf(string(html[i]))
		//tag?
		if html[i] == '<' {
			fmt.Println("is <")
			//opening tag?
			if html[i+1] == '/' {
				//closing tag!
				child := stack.Pop()
				parent := stack.Top()
				parent.Children = append(parent.Children, child)
			} else {
				fmt.Println("opening tag!")
				node, end := ParseTag(html[:i+1])
				fmt.Println(node.Name, end)
				stack = append(stack, node)
				i = i + end

				node.Children = append(node.Children, node)
			}
		}
	}

	return stack[0], nil
}

func ParseTag(html string) (*Node, int) {
	var node Node

	var attrs string
	var tagEnd int
	var end int

	for i := 0; ; i++ {
		if html[i] == '>' {
			attrs = html[:i-1]
			tagEnd = i + 1
		} else if html[i] == '<' {
			attrs = html[:i-1]
			if i-tagEnd > 0 {
				node.Text = html[i:tagEnd]
			}
			end = i
			break
		}
	}

	//initilaze attrs
	attrs = strings.TrimSpace(attrs)
	textArr := strings.Split(attrs, " ")
	node.Name = textArr[0]

	attributes := make(map[string]string)

	for _, part := range textArr[1:] {
		attrArr := strings.SplitN(part, "=", 2)
		if len(attrArr) == 2 {
			attributes[attrArr[0]] = attrArr[1]
		}
	}
	node.Attributes = attributes
	fmt.Print(node.Name)
	return &node, end
}

func PrintTree(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(" ")

	fmt.Printf("%s", node.Name)
	for k, v := range node.Attributes {
		fmt.Printf(" %s=%s", k, v)
	}
	fmt.Printf(" %s", node.Text)
	fmt.Println()

	// Print the node's children.
	for _, child := range node.Children {
		PrintTree(child)
	}
}
