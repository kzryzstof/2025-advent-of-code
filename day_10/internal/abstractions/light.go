package abstractions

type Light struct {
	isOn bool
}

func NewLight(isOn bool) *Light {
	return &Light{isOn: isOn}
}

func (l *Light) IsOn() bool {
	return l.isOn
}

func (l *Light) Toggle() {
	l.isOn = !l.isOn
}
