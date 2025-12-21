package abstractions

type Graph struct {
	nodes       []*Node
	nodesByName map[string]*Node
}

func BuildGraph(
	devices []*Device,
) *Graph {

	graph := &Graph{
		[]*Node{},
		map[string]*Node{},
	}

	for _, device := range devices {
		deviceNode := graph.getNodeByName(device.name)

		if deviceNode == nil {
			deviceNode = graph.createNewNode(device.name)
			graph.addNodeToRoot(deviceNode)
		}
		for _, outputDeviceName := range device.outputs {
			outputNode := graph.getNodeByName(outputDeviceName)

			if outputNode == nil {
				outputNode = graph.createNewNode(outputDeviceName)
			}

			deviceNode.AddNext(outputNode)
		}
	}

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
) *Node {
	newNode := NewNode(deviceName)
	g.nodesByName[deviceName] = newNode
	return newNode
}
