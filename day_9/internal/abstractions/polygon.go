package abstractions

import (
	"math"
	"sort"
)

func IsPointInPolygon(
	polygon []*Tile,
	point *Tile,
) bool {

	if len(polygon) < 3 {
		return false
	}

	inside := false
	n := len(polygon)

	for i := 0; i < n; i++ {

		j := (i + 1) % n

		xi, yi := polygon[i].X, polygon[i].Y
		xj, yj := polygon[j].X, polygon[j].Y
		px, py := point.X, point.Y

		intersectY := (yi > py) != (yj > py)

		intersectX := false

		if intersectY {
			if !inside {
				intersectX = px <= (xj-xi)*(py-yi)/(yj-yi)+xi
			} else {
				intersectX = px < (xj-xi)*(py-yi)/(yj-yi)+xi
			}
		}

		// Check if ray from point crosses this edge
		if intersectY && intersectX {
			inside = !inside
		}
	}

	return inside
}

func OrderPolygonVertices(
	tiles []*Tile,
) []*Tile {

	if len(tiles) < 3 {
		return tiles
	}

	// Find centroid
	var sumX, sumY float64
	for _, t := range tiles {
		sumX += float64(t.X)
		sumY += float64(t.Y)
	}
	cx := sumX / float64(len(tiles))
	cy := sumY / float64(len(tiles))

	// Sort by angle from centroid
	ordered := make([]*Tile, len(tiles))
	copy(ordered, tiles)

	sort.Slice(ordered, func(i, j int) bool {
		angle1 := math.Atan2(float64(ordered[i].Y)-cy, float64(ordered[i].X)-cx)
		angle2 := math.Atan2(float64(ordered[j].Y)-cy, float64(ordered[j].X)-cx)
		return angle1 < angle2
	})

	return ordered
}
