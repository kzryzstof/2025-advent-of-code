package abstractions

type Node struct {
	name []string
	next []*Node
}

func NewNode(
	name []string,
) *Node {
	return &Node{name, []*Node{}}
}
