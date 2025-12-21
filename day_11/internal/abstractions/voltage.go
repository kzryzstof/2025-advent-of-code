package abstractions

type Voltage struct {
	value int64
}

func NewVoltage(
	value int64,
) *Voltage {
	return &Voltage{value}
}

func (v *Voltage) GetValue() int64 {
	return v.value
}
