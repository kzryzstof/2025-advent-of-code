package abstractions

type Button struct {
	Light int
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
	light := machine.GetLight(b.Light)
	light.Toggle()
}
