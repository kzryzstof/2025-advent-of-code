package abstractions

type Machine struct {
	LightIndicators []*LightIndicator
	ButtonGroups    []*ButtonGroup
	Voltages        []*Voltage
}
