package abstractions

type Dimension struct {
	Wide      int
	Long      int
	FillRatio float64
}

func (d Dimension) GetRatio() float64 {
	return float64(d.Wide) / float64(d.Long)
}

func (d Dimension) GetArea() uint {
	return uint(d.Wide) * uint(d.Long)
}

func (d Dimension) GetPerimeter() uint {
	return 2*uint(d.Wide) + 2*uint(d.Long)
}

func (d Dimension) Equals(
	other Dimension,
) bool {
	return d.Wide == other.Wide && d.Long == other.Long
}

func (d Dimension) IsSquare() bool {
	return d.Wide == d.Long
}

func (d Dimension) IsEmpty() bool {
	return d.Wide == 0 && d.Long == 0
}
