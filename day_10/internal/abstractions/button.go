package abstractions

type Button struct {
	LightIndicator int
}

type ButtonGroup struct {
	Buttons []*Button
}
