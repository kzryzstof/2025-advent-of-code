package abstractions

import "math"

type Rectangle struct {
	A             *Tile
	B             *Tile
	area          uint64
	topLeftCorner *Tile
	width, height uint
}

func NewRectangle(a, b *Tile) *Rectangle {
	width := math.Abs(float64(b.X)-float64(a.X)) + 1
	height := math.Abs(float64(b.Y)-float64(a.Y)) + 1

	topLeftCorner := &Tile{X: uint(math.Min(float64(a.X), float64(b.X))), Y: uint(math.Min(float64(a.Y), float64(b.Y)))}

	return &Rectangle{
		A:             a,
		B:             b,
		area:          uint64(width) * uint64(height),
		topLeftCorner: topLeftCorner,
		width:         uint(width),
		height:        uint(height),
	}
}

func (r *Rectangle) GetArea() uint64 {
	return r.area
}

func (r *Rectangle) IsInside(tiles []*Tile) bool {
	for x := r.topLeftCorner.X; x < r.topLeftCorner.X+r.width; x++ {
		for y := r.topLeftCorner.Y; y < r.topLeftCorner.Y+r.height; y++ {
			tileColor := GetTileColor(tiles, x, y)
			if tileColor != Red && tileColor != Green {
				return false
			}
		}
	}
	return true
}
