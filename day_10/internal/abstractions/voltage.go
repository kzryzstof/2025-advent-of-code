package abstractions

type Voltage struct {
	value uint32
}

func NewVoltage(
	value uint32,
) *Voltage {
	return &Voltage{value}
}

func (v *Voltage) GetValue() uint32 {
	return v.value
}
