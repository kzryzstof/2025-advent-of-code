package abstractions

type Node struct {
	name string
	next []*Node
}

func NewNode(
	name string,
) *Node {
	return &Node{name, []*Node{}}
}

func (n *Node) FindNodeByName(
	deviceName string,
) *Node {
	if deviceName == n.name {
		return n
	}

	if n.next == nil {
		return nil
	}

	for _, nextNode := range n.next {
		node := nextNode.FindNodeByName(deviceName)

		if node != nil {
			return node
		}
	}

	return nil
}

func (n *Node) AddNext(
	outputNext *Node,
) {
	n.next = append(n.next, outputNext)
}
