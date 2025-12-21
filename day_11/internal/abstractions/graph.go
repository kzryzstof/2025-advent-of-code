package abstractions

type Graph struct {
	nodes       []*Node
	nodesByName map[string]*Node
}

func BuildGraph(
	devices []*Device,
) *Graph {

	rootNode := NewNode("root")
	graph := &Graph{
		[]*Node{rootNode},
		map[string]*Node{"root": rootNode},
	}

	for _, device := range devices {
		deviceNode := graph.getNodeByName(device.name)

		if deviceNode == nil {
			deviceNode = graph.createNewNode(device.name)
			graph.addNodeToRoot(deviceNode)
		}
		for _, outputDevice := range device.outputs {
			outputNode := graph.getNodeByName(outputDevice)

			if outputNode == nil {
				outputNode = graph.createNewNode(device.name)
			}

			deviceNode.AddNext(outputNode)
		}
	}

	return graph
}

func (g *Graph) GetRootNodesCount() uint {
	return uint(len(g.nodes))
}

func (g *Graph) getNodeByName(
	deviceName string,
) *Node {

	/* Finds out if a node already exists */
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
) *Node {
	newNode := NewNode(deviceName)
	g.nodesByName[deviceName] = newNode
	return newNode
}

func (g *Graph) CountPaths(
	from string,
	to string,
) uint {

	return 0
}
