package abstractions

import (
	"slices"
)

type Node struct {
	name          string
	parents       []*Node
	next          []*Node
	requiredNodes []string
	newPathsCount int64
}

func NewNode(
	name string,
	requiredNodes []string,
) *Node {
	return &Node{
		name,
		[]*Node{},
		[]*Node{},
		requiredNodes,
		-1,
	}
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

	if len(outputNext.requiredNodes) > 0 {
		n.setHasRequiredNodes(outputNext.requiredNodes)
	}
}

func (n *Node) setHasRequiredNodes(
	requiredNodes []string,
) {
	requiredNodeAdded := false

	for _, requiredNode := range requiredNodes {
		if !slices.Contains(n.requiredNodes, requiredNode) {
			requiredNodeAdded = true
			n.requiredNodes = AddOnce(n.requiredNodes, requiredNode)
		}
	}

	if !requiredNodeAdded {
		/*
			Important.
			Either there were no required nodes or the required nodes have already been added.
			In either case, no need to check parents, since they will have the required nodes
		*/
		return
	}

	if n.parents == nil {
		return
	}

	/* Also indicates that parents have the required nodes */
	for _, parentNode := range n.parents {
		parentNode.setHasRequiredNodes(requiredNodes)
	}
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

	if n.newPathsCount != -1 {
		/*
			We have already visited this node before. Let's return the cached value containing the number of paths.
		*/
		return currentCount + uint(n.newPathsCount)
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

	previousCurrentCount := currentCount

	for _, nextNode := range n.next {

		if n.requiredNodes == nil || len(nextNode.requiredNodes) < (len(requiredNodes)-len(encounteredNodes)) {
			/*
				All the required nodes have not been encountered yet, so
				we are selective about the next node to go. If all the required
				nodes have been encountered, we can go to the children.
			*/
			continue
		}

		currentCount = nextNode.testPaths(
			to,
			currentCount,
			requiredNodes,
			encounteredNodes,
		)
	}

	/*
		Important.
		Keep track how many new paths have been added going through all the nodes below.
		Since the nodes are being visited multiple times, again and again, let's cache the value.
	*/
	n.newPathsCount = int64(currentCount) - int64(previousCurrentCount)
	return currentCount
}
