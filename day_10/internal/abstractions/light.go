package abstractions

type Light struct {
	isOn         bool
	initialState bool
}

func NewLight(isOn bool) *Light {
	return &Light{
		isOn:         isOn,
		initialState: isOn,
	}
}

func (l *Light) IsOn() bool {
	return l.isOn
}

func (l *Light) Toggle() {
	l.isOn = !l.isOn
}

func (l *Light) Reset() {
	l.isOn = l.initialState
}
