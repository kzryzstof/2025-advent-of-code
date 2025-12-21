package abstractions

type Graph struct {
	nodes []*Node
}

func BuildGraph(
	devices []*Device,
) *Graph {

	rootNode := NewNode([]string{"root"})

	return &Graph{
		[]*Node{rootNode},
	}
}

func (g *Graph) CountPaths(
	from string,
	to string,
) uint {

	return 0
}
