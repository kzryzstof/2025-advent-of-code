package abstractions

import "math"

type Position struct {
	X uint64
	Y uint64
	Z uint64
}

func (p Position) Distance(other Position) float64 {

	x := math.Pow(float64(p.X)-float64(other.X), 2)
	y := math.Pow(float64(p.Y)-float64(other.Y), 2)
	z := math.Pow(float64(p.Z)-float64(other.Z), 2)

	return math.Sqrt(x + y + z)
}
