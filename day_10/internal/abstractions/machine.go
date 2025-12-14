package abstractions

type Machine struct {
	lights       []*Light
	buttonGroups []*ButtonGroup
	voltages     []*Voltage
}

func NewMachine(
	lights []*Light,
	buttonGroups []*ButtonGroup,
) *Machine {
	return &Machine{
		lights,
		buttonGroups,
		nil,
	}
}

func (m *Machine) IsOn() bool {
	for _, light := range m.lights {
		if !light.IsOn() {
			return false
		}
	}

	return true
}

func (m *Machine) GetLight(
	number int,
) *Light {
	return m.lights[number]
}

func (m *Machine) PressGroup(
	groupIndex int,
) {
	buttonGroup := m.buttonGroups[groupIndex]

	for _, button := range buttonGroup.Buttons {
		button.Press(m)
	}
}
