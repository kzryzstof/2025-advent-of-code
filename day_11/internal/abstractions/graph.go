package abstractions

import (
	"fmt"
	"slices"
)

type Graph struct {
	nodes       []*Node
	nodesByName map[string]*Node
}

func BuildGraph(
	devices []*Device,
	requiredNodes []string,
) *Graph {

	graph := &Graph{
		[]*Node{},
		map[string]*Node{},
	}

	for deviceIndex, device := range devices {

		fmt.Printf("Processing device '%s' | %d / %d...\r", device.name, deviceIndex, len(devices))

		isRequiredNode := slices.Contains(requiredNodes, device.name)

		deviceNode := graph.getNodeByName(device.name)

		if deviceNode == nil {
			deviceNode = graph.createNewNode(device.name, isRequiredNode)
			graph.addNodeToRoot(deviceNode)
		}
		for _, outputDeviceName := range device.outputs {
			outputNode := graph.getNodeByName(outputDeviceName)

			if outputNode == nil {
				outputNode = graph.createNewNode(outputDeviceName, isRequiredNode)
			}

			deviceNode.AddNext(outputNode)
		}
	}

	fmt.Println()

	return graph
}

func (g *Graph) CountPaths(
	from string,
	to string,
	requiredNodes []string,
) uint {

	fromNode := g.getNodeByName(from)

	return fromNode.CountPathsTo(to, requiredNodes)
}

func (g *Graph) getNodeByName(
	deviceName string,
) *Node {
	node, isNode := g.nodesByName[deviceName]

	if isNode {
		return node
	}

	return nil
}

func (g *Graph) addNodeToRoot(
	node *Node,
) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) createNewNode(
	deviceName string,
	isRequiredNode bool,
) *Node {
	newNode := NewNode(deviceName, isRequiredNode)
	g.nodesByName[deviceName] = newNode
	return newNode
}

func (g *Graph) CountPathsBackwards(
	from string,
	to string,
	requiredNodes []string,
) uint {

	toNode := g.getNodeByName(to)

	return toNode.CountPathsToBackwards(from, requiredNodes)
}
