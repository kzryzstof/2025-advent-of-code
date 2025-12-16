package abstractions

type Button struct {
	CounterIndex int
}

type ButtonGroup struct {
	Buttons []*Button
}

func (g *ButtonGroup) Press(
	machine *Machine,
) {
	for _, button := range g.Buttons {
		button.Press(machine)
	}
}

func (b *Button) Press(
	machine *Machine,
) {
	counter := machine.GetCounter(b.CounterIndex)
	counter.Increment()
}
