package abstractions

import "math"

type Rectangle struct {
	A    *Tile
	B    *Tile
	area uint64
}

func NewRectangle(a, b *Tile) *Rectangle {
	width := math.Abs(float64(b.X)-float64(a.X)) + 1
	height := math.Abs(float64(b.Y)-float64(a.Y)) + 1

	return &Rectangle{
		A:    a,
		B:    b,
		area: uint64(width) * uint64(height),
	}
}

func (r *Rectangle) GetArea() uint64 {
	return r.area
}
