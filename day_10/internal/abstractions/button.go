package abstractions

type Button struct {
	Light int
}

type ButtonGroup struct {
	Buttons []*Button
}

func (b *Button) Press(
	machine *Machine,
) {
	light := machine.GetLight(b.Light)
	light.Toggle()
}
