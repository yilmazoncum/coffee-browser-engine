package engine

import (
	"fmt"
	"strconv"
)

func LayoutTree(rn *RenderNode) *LayoutNode {
	root := buildLayoutTree(rn, 0, 0)
	return root
}

func buildLayoutTree(rn *RenderNode, x int, y int) *LayoutNode {
	layoutNode := &LayoutNode{
		RenderNode: rn,
		Box: &Box{
			X: x,
			Y: y,
		},
	}

	layoutNode.Box = BuildBox(rn.Style, layoutNode.Box.X, layoutNode.Box.Y)

	for _, child := range rn.Children {
		renderChild := buildLayoutTree(child, layoutNode.Box.X, layoutNode.Box.Y)
		layoutNode.Children = append(layoutNode.Children, renderChild)
	}

	return layoutNode
}

func BuildBox(style map[string]string, currentX int, currentY int) *Box {
	box := &Box{
		Width:  ExtractNumber(style["width"]),
		Height: ExtractNumber(style["height"]),
		X:      (currentX + CalculateXY(style, true)),
		Y:      (currentY + CalculateXY(style, false)),
	}

	return box
}

func ExtractNumber(str string) int {
	for i := range str {
		if str[i] >= 65 {
			num, _ := strconv.Atoi(str[:i])
			return num
		}
	}
	return 0
}

func CalculateXY(style map[string]string, opt bool) int {
	dict := make(map[string]bool)
	if opt {
		//X coord
		//true is + , false is -
		dict = map[string]bool{"left": true, "right": false, "margin": true, "margin-left": true, "margin-right": false}
	} else {
		//Y coord
		dict = map[string]bool{"top": true, "bottom": false, "margin": true, "margin-top": true, "margin-bottom": false}
	}
	sum := 0

	//O(N) contains algorithm
	set := make(map[string]bool)
	for k := range dict {
		set[k] = true
	}
	for k, v := range style {
		if set[k] {
			if dict[k] {
				sum += ExtractNumber(v)
			} else {
				sum -= ExtractNumber(v)
			}
		}
	}
	//O(N) contains algorithm

	return sum
}

func PrintLayoutTree(node *LayoutNode) {
	if node == nil {
		return
	}

	fmt.Printf("%s: ", node.RenderNode.Name)

	fmt.Println()
	box := node.Box
	fmt.Println("   ", "height", box.Height)
	fmt.Println("   ", "width", box.Width)
	fmt.Println("   ", "x", box.X)
	fmt.Println("   ", "y", box.Y)

	// Print the node's children.
	for _, child := range node.Children {
		fmt.Println()
		PrintLayoutTree(child)
	}
}
