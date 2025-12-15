package abstractions

type Light struct {
	isOn          bool
	expectedState bool
}

func NewLight(isOn bool) *Light {
	return &Light{
		isOn:          false,
		expectedState: isOn,
	}
}

func (l *Light) IsValid() bool {
	return l.isOn == l.expectedState
}

func (l *Light) Toggle() {
	l.isOn = !l.isOn
}

func (l *Light) Close() {
	l.isOn = false
}
