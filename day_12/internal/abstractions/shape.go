package abstractions

type Shape struct {
	Dimension Dimension
	Cells     [][]byte
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

func (s Shape) GetCopy() [][]byte {
	return GetCopy(s.Cells)
}
