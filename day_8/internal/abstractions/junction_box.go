package abstractions

type JunctionBox struct {
	Position Position
}

func (j JunctionBox) MeasureDistance(
	other JunctionBox,
) float64 {
	return j.Position.Distance(other.Position)
}
