package abstractions

import "fmt"

type Machine struct {
	buttonGroups []*ButtonGroup
	voltages     []*Voltage
	counters     []*Counter
}

/* Each machine has a set to counters (initialized at 0) */

func NewMachine(
	buttonGroups []*ButtonGroup,
	voltages []*Voltage,
) *Machine {
	counters := make([]*Counter, len(voltages))

	for i := range counters {
		counters[i] = NewCounter()
	}

	return &Machine{
		buttonGroups,
		voltages,
		counters,
	}
}

func (m *Machine) GetButtonGroups() []*ButtonGroup {
	return m.buttonGroups
}

func (m *Machine) GetVoltages() []*Voltage {
	return m.voltages
}

func (m *Machine) GetButtonGroupsCount() int {
	return len(m.buttonGroups)
}

func (m *Machine) IsVoltageValid() bool {

	/* The machine is activated if all lights are in their expected states */
	for voltageIndex, voltage := range m.voltages {
		if voltage.GetValue() != m.counters[voltageIndex].GetValue() {
			return false
		}
	}

	return true
}

func (m *Machine) PrintVoltages() {

	fmt.Print("-- Voltages --\n")
	/* The machine is activated if all lights are in their expected states */
	for voltageIndex, voltage := range m.voltages {
		fmt.Printf("\tVoltage %d: expected=%d, actual=%d\n", voltageIndex, voltage.GetValue(), m.counters[voltageIndex].GetValue())
	}
}

func (m *Machine) GetCounter(
	number int,
) *Counter {
	return m.counters[number]
}

func (m *Machine) PressGroup(
	groupIndex int,
) {
	buttonGroup := m.buttonGroups[groupIndex]

	for _, button := range buttonGroup.Buttons {
		button.Press(m)
	}
}
