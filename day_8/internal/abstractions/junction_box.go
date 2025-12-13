package abstractions

import "math"

type JunctionBox struct {
	Position Position
}

func (j JunctionBox) Distance(other JunctionBox) float64 {

	x := math.Pow(float64(j.Position.X)-float64(other.Position.X), 2)
	y := math.Pow(float64(j.Position.Y)-float64(other.Position.Y), 2)
	z := math.Pow(float64(j.Position.Z)-float64(other.Position.Z), 2)

	return math.Sqrt(x + y + z)
}
