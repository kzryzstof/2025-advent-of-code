package maths

type Dimension struct {
	Wide uint
	Long uint
}

func (d Dimension) Equals(
	other Dimension,
) bool {
	return d.Wide == other.Wide && d.Long == other.Long
}
