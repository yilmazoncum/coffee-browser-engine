package engine

type RenderNode struct {
	Name string

	Attributes map[string]string

	Text string

	Style map[string]string

	Children []*RenderNode
}
