package parsers

import (
	"fmt"
	"strings"
	"unicode"
)

func ParseHtml(html string) (*Node, error) {
	var stack Stack
	stack.Push(&Node{
		Name: "root",
	})
	for i := 0; i < len(html); i++ {
		if unicode.IsSpace(rune(html[i])) {
			continue
		}
		//tag?
		if html[i] == '<' {
			//opening tag?
			if html[i+1] == '/' {
				//closing tag!
				child := stack.Pop()
				parent := stack.Top()
				parent.Children = append(parent.Children, child)
			} else {
				node, end := ParseTag(html[i+1:])
				stack = append(stack, node)
				i = i + end
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
		if unicode.IsSpace(rune(html[i])) {
			continue
		}
		if html[i] == '>' {
			attrs = html[:i]
			tagEnd = i + 1
		} else if html[i] == '<' {
			if i-tagEnd > 0 {
				node.Text = strings.TrimSpace(html[tagEnd:i])
			}
			end = i
			break
		}
	}

	//initialize attrs
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
	return &node, end
}

var indentLevel int

func PrintTree(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%s%s: ", indent(0), node.Name)
	indentLevel++
	if len(node.Attributes) != 0 {
		var attrs []string
		for k, v := range node.Attributes {
			attrs = append(attrs, fmt.Sprintf("%s=%s", k, v))
		}
		fmt.Printf("[%s]", strings.Join(attrs, " "))
	}

	if node.Text != "" {
		fmt.Printf("\n%s%s", indent(1), node.Text)
	}

	// Print the node's children.
	for _, child := range node.Children {
		indentLevel += 1
		fmt.Println()
		PrintTree(child)
	}
	indentLevel--
}

func indent(extra int) string {
	return strings.Repeat(" ", indentLevel+extra)
}
