package abstractions

type Circuit struct {
	junctionBoxes []*JunctionBox
}

func NewCircuit(
	junctionBox *JunctionBox,
) *Circuit {
	junctionBoxes := make([]*JunctionBox, 0)
	junctionBoxes = append(junctionBoxes, junctionBox)
	return &Circuit{junctionBoxes: junctionBoxes}
}

func (c *Circuit) Add(
	junctionBox *JunctionBox,
) {
	c.junctionBoxes = append(c.junctionBoxes, junctionBox)
}

func (c *Circuit) HasSingleJunctionBox() bool {
	return len(c.junctionBoxes) == 1
}

func (c *Circuit) Contains(
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
	if c == nil {
		return 0
	}
	return len(c.junctionBoxes)
}

func (c *Circuit) Get() []*JunctionBox {
	return c.junctionBoxes
}
