package engine

type RenderNode struct {
	Name string

	Attributes map[string]string

	Text string

	Style map[string]string

	Children []*RenderNode
}

type Box struct {
	x      int
	y      int
	width  int
	height int
}

type LayoutNode struct {
	Box *Box

	RenderNode *RenderNode

	Children []*LayoutNode
}
