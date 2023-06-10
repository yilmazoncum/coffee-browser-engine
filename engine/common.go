package engine

type RenderNode struct {
	Name string

	Attributes map[string]string

	Text string

	Style map[string]string

	Children []*RenderNode
}

type Box struct {
	X      int
	Y      int
	Width  int
	Height int
}

type LayoutNode struct {
	Box *Box

	RenderNode *RenderNode

	Children []*LayoutNode
}
