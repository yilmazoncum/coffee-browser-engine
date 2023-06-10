package rasterizer

import "main/engine"

type Queue []*engine.LayoutNode

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		panic("stack is empty")
	}
	*q = (*q)[1:]
}

func (q *Queue) Enqueue(node *engine.LayoutNode) {
	*q = append(*q, node)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Top() *engine.LayoutNode {
	return (*q)[0]
}
