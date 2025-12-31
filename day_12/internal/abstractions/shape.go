package abstractions

const (
	// E /* Indicates an empty spot; it is not 0 since we have present with an index of 0 */
	E = -99
	// P /* For debugging purpose when packing a new shape: it helps highlight the new shape */
	P = -1
)

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
