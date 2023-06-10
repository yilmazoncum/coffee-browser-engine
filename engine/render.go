package engine

import (
	"fmt"
	"main/parsers"
	"strings"
)

func RenderTree(dom *parsers.Node, ss *parsers.Stylesheet) *RenderNode {

	root := buildTree(dom, ss)
	return root

}

func buildTree(dom *parsers.Node, ss *parsers.Stylesheet) *RenderNode {
	renderNode := &RenderNode{
		Name:       dom.Name,
		Attributes: dom.Attributes,
		Text:       dom.Text,
	}
	renderNode.Style = StyleMatch(renderNode, ss)

	for _, child := range dom.Children {
		renderChild := buildTree(child, ss)
		renderNode.Children = append(renderNode.Children, renderChild)
	}

	return renderNode
}

func StyleMatch(node *RenderNode, ss *parsers.Stylesheet) map[string]string {
	styles := make(map[string]string)

	for _, rule := range ss.Rules {
		if rule.Selector == node.Name {
			styles = rule.Properties
			break
		}

		if rule.Selector[0] == '.' && node.Attributes["class"] == rule.Selector[1:] {
			styles = rule.Properties
			break
		}

		if rule.Selector[0] == '#' && node.Attributes["id"] == rule.Selector[1:] {
			styles = rule.Properties
			break
		}
	}

	return styles
}

var indentLevel int

func PrintRenderTree(node *RenderNode) {
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

	if len(node.Style) != 0 {
		var styles []string
		for k, v := range node.Style {
			styles = append(styles, fmt.Sprintf("%s=%s", k, v))
		}
		fmt.Printf("[%s]", strings.Join(styles, " "))
	}

	// Print the node's children.
	for _, child := range node.Children {
		indentLevel += 1
		fmt.Println()
		PrintRenderTree(child)
	}
	indentLevel--
}

func indent(extra int) string {
	return strings.Repeat(" ", indentLevel+extra)
}
