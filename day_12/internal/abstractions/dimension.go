package abstractions

type Dimension struct {
	Wide int
	Long int
}

func (d Dimension) GetArea() uint {
	return uint(d.Wide) * uint(d.Long)
}

func (d Dimension) Equals(
	other Dimension,
) bool {
	return d.Wide == other.Wide && d.Long == other.Long
}

func (d Dimension) IsLessThan(
	other Dimension,
) bool {
	return d.GetArea() < other.GetArea()
}

func (d Dimension) IsEmpty() bool {
	return d.Wide == 0 && d.Long == 0
}
