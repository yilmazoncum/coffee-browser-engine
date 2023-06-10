package parsers

type Node struct {
	Name string

	Attributes map[string]string

	Children []*Node

	Text string
}

type Rule struct {
	Selector   string
	Properties map[string]string
}

type Stylesheet struct {
	Rules []Rule
}

type Stack []*Node

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(node *Node) {
	*s = append(*s, node) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

func (s *Stack) Top() *Node {
	return (*s)[len(*s)-1]
}
