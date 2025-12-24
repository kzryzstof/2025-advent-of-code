package abstractions

type Dimension struct {
	Wide int
	Long int
}

func (d Dimension) Equals(
	other Dimension,
) bool {
	return d.Wide == other.Wide && d.Long == other.Long
}
