package abstractions

import "slices"

type Circuits struct {
	circuits []*Circuit
}

func NewCircuits() *Circuits {
	return &Circuits{circuits: make([]*Circuit, 0)}
}

func (c *Circuits) Count() int {
	return len(c.circuits)
}

func (c *Circuits) Add(
	circuit *Circuit,
) {
	c.circuits = append(c.circuits, circuit)
}

func (c *Circuits) Remove(
	circuit *Circuit,
) {
	c.circuits = slices.DeleteFunc(c.circuits, func(c *Circuit) bool {
		return c == circuit
	})
}

func (c *Circuits) Get(
	junctionBox *JunctionBox,
) *Circuit {
	for _, circuit := range c.circuits {
		if circuit.ContainsJunctionBox(junctionBox) {
			return circuit
		}
	}
	return nil
}
