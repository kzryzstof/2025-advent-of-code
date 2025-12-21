package abstractions

import "slices"

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
	requiredNodes []string,
) uint {

	currentCount := uint(0)

	return n.testPaths(to, currentCount, requiredNodes, []string{})
}

func (n *Node) testPaths(
	to string,
	currentCount uint,
	requiredNodes []string,
	encounteredNodes []string,
) uint {

	if slices.Contains(requiredNodes, n.name) {
		encounteredNodes = AddOnce(encounteredNodes, n.name)
	}

	if n.name == to {
		if len(requiredNodes) == len(encounteredNodes) {
			return currentCount + 1
		}
		return currentCount
	}

	currentEncounteredNodes := make([]string, len(encounteredNodes))
	copy(currentEncounteredNodes, encounteredNodes)

	for _, nextNode := range n.next {
		currentCount = nextNode.testPaths(
			to,
			currentCount,
			requiredNodes,
			currentEncounteredNodes,
		)
	}

	return currentCount
}
