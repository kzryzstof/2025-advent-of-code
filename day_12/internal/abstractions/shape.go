package abstractions

import "day_12/internal/maths"

const (
	// E /* Indicates an empty spot; it is not 0 since we have present with an index of 0 */
	E = -99
	// P /* For debugging purpose when packing a new shape: it helps highlight the new shape */
	P = -1
)

type Shape struct {
	Dimension maths.Dimension
	Cells     [][]int8
	FillRatio float64
}

func NewShape(
	dimension maths.Dimension,
	cells [][]int8,
) Shape {
	return Shape{
		dimension,
		cells,
		computeFillRatio(cells),
	}
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
	return maths.CopySlice(s.Cells)
}

func computeFillRatio(
	slice [][]int8,
) float64 {

	empty, occupied := 0, 0

	for row := 0; row < len(slice); row++ {
		for col := 0; col < len(slice[row]); col++ {
			if slice[row][col] == E {
				empty++
			} else {
				occupied++
			}
		}
	}

	return float64(occupied) / float64(occupied+empty)
}
