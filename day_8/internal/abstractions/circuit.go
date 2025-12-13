package abstractions

type Circuit struct {
	junctionBoxes []*JunctionBox
}

func NewCircuit(junctionBox *JunctionBox) *Circuit {
	junctionBoxes := make([]*JunctionBox, 0)
	junctionBoxes = append(junctionBoxes, junctionBox)
	return &Circuit{junctionBoxes: junctionBoxes}
}

func (c *Circuit) AddJunctionBox(junctionBox *JunctionBox) {
	c.junctionBoxes = append(c.junctionBoxes, junctionBox)
}

func (c *Circuit) HasSingleJunctionBox() bool {
	return len(c.junctionBoxes) == 1
}

func GetCircuit(
	circuits []*Circuit,
	junctionBox *JunctionBox,
) *Circuit {
	for _, circuit := range circuits {
		if circuit.ContainsJunctionBox(junctionBox) {
			return circuit
		}
	}

	return nil
}

func (c *Circuit) ContainsJunctionBox(
	junctionBox *JunctionBox,
) bool {
	for _, jb := range c.junctionBoxes {
		if jb == junctionBox {
			return true
		}
	}

	return false
}

func (c *Circuit) Count() int {
	return len(c.junctionBoxes)
}

func (c *Circuit) GetJunctionBoxes() []*JunctionBox {
	return c.junctionBoxes
}
