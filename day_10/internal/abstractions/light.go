package abstractions

type Light struct {
	isOn         bool
	initialState bool
}

func NewLight(isOn bool) *Light {
	return &Light{
		isOn:         false,
		initialState: isOn,
	}
}

func (l *Light) IsValid() bool {
	return l.isOn == l.initialState
}

func (l *Light) Toggle() {
	l.isOn = !l.isOn
}

func (l *Light) Reset() {
	l.isOn = false
}
