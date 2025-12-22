package abstractions

import "slices"

type Node struct {
	name    string
	parents []*Node
	next    []*Node
	/* True if the one of the nodes below has one of the required paths */
	hasRequiredNodes bool
}

func NewNode(
	name string,
	isRequiredNode bool,
) *Node {
	return &Node{name, []*Node{}, []*Node{}, isRequiredNode}
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
	outputNext.parents = append(outputNext.parents, n)

	if outputNext.hasRequiredNodes {
		n.setHasRequiredNodes()
	}
}

func (n *Node) setHasRequiredNodes() {
	n.hasRequiredNodes = true

	if n.parents == nil {
		return
	}

	/* Also indicates that parents have the required nodes */
	for _, parentNode := range n.parents {
		parentNode.setHasRequiredNodes()
	}
}

func (n *Node) CountPathsTo(
	to string,
	requiredNodes []string,
) uint {
	currentCount := uint(0)

	return n.testPaths(to, currentCount, requiredNodes, []string{}, []string{})
}

func (n *Node) testPaths(
	to string,
	currentCount uint,
	requiredNodes []string,
	encounteredNodes []string,
	visitedNodes []string,
) uint {

	Print(visitedNodes, currentCount)

	if !slices.Contains(visitedNodes, n.name) {
		visitedNodes = AddOnce(visitedNodes, n.name)
	} else {
		/* Loop detected */
		return currentCount
	}

	if slices.Contains(requiredNodes, n.name) {
		encounteredNodes = AddOnce(encounteredNodes, n.name)
	}

	if n.name == to {
		if len(requiredNodes) == len(encounteredNodes) {
			return currentCount + 1
		}
		return currentCount
	}

	for _, nextNode := range n.next {
		if !nextNode.hasRequiredNodes {
			continue
		}
		currentCount = nextNode.testPaths(
			to,
			currentCount,
			requiredNodes,
			encounteredNodes,
			visitedNodes,
		)
	}

	return currentCount
}

func (n *Node) CountPathsToBackwards(
	to string,
	requiredNodes []string,
) uint {

	currentCount := uint(0)

	return n.testPathsBackwards(to, currentCount, requiredNodes, []string{}, []string{})
}

func (n *Node) testPathsBackwards(
	to string,
	currentCount uint,
	requiredNodes []string,
	encounteredNodes []string,
	visitedNodes []string,
) uint {

	Print(visitedNodes, currentCount)

	if !slices.Contains(visitedNodes, n.name) {
		visitedNodes = AddOnce(visitedNodes, n.name)
	} else {
		/* Loop detected */
		return currentCount
	}

	if slices.Contains(requiredNodes, n.name) {
		encounteredNodes = AddOnce(encounteredNodes, n.name)
	}

	if n.name == to {
		if len(requiredNodes) == len(encounteredNodes) {
			return currentCount + 1
		}
		return currentCount
	}

	for _, parentNode := range n.parents {
		currentCount = parentNode.testPathsBackwards(
			to,
			currentCount,
			requiredNodes,
			encounteredNodes,
			visitedNodes,
		)
	}

	return currentCount
}
