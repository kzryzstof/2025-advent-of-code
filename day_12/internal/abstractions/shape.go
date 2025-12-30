package abstractions

type Shape struct {
	Dimension Dimension
	Cells     [][]int8
	FillRatio float64
}

func (s Shape) IsMoreOptimalThan(
	other Shape,
) bool {

	currentFillRatio := s.FillRatio
	otherFillRatio := other.FillRatio

	if currentFillRatio > otherFillRatio {
		return true
	}

	return false
}

func (s Shape) GetCopy() [][]int8 {
	return GetCopy(s.Cells)
}
