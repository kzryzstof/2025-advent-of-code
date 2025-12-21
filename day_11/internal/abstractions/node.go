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

func (n *Node) CountPathsTo(
	to string,
) uint {
	if n.name == to {
		return 1
	}

	currentCount := uint(0)

	return n.testPaths(to, currentCount)
}

func (n *Node) testPaths(
	to string,
	currentCount uint,
) uint {

	if n.name == to {
		return currentCount + 1
	}

	for _, nextNode := range n.next {
		currentCount = nextNode.testPaths(to, currentCount)
	}

	return currentCount
}
