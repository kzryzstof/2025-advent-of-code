package abstractions

import "math"

type Dimension struct {
	Wide int
	Long int
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

func (d Dimension) IsMoreOptimalThan(
	other Dimension,
	region *Region,
) bool {
	ratioDelta := math.Abs(d.GetRatio() - region.GetRatio())

	otherRatioDelta := math.Abs(other.GetRatio() - region.GetRatio())

	return ratioDelta < otherRatioDelta
}

func (d Dimension) IsEmpty() bool {
	return d.Wide == 0 && d.Long == 0
}
