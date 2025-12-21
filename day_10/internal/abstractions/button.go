package abstractions

type Button struct {
	counterIndex int
}

func NewButton(
	counterIndex int,
) *Button {
	return &Button{counterIndex}
}

func (b *Button) Press(
	machine *Machine,
) {
	counter := machine.GetCounter(b.counterIndex)
	counter.Increment()
}
