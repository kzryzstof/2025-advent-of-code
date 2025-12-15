package abstractions

var combinations chan []int

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

func (m *Machine) GetButtonGroupsCount() int {
	return len(m.buttonGroups)
}

func (m *Machine) IsActivated() bool {

	/* The machine is activated if all lights are in their expected states */
	for _, light := range m.lights {
		if !light.IsValid() {
			return false
		}
	}

	return true
}

func (m *Machine) CloseLights() {
	for _, light := range m.lights {
		light.Close()
	}
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
