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

	px, py := point.X, point.Y
	inside := false
	n := len(polygon)

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi, yi := polygon[i].X, polygon[i].Y
		xj, yj := polygon[j].X, polygon[j].Y

		/* Fast check if the point is on either vertex */
		if (px == xi && py == yi) || (px == xj && py == yj) {
			return true
		}

		/* Checks if point lies exactly on the edge segment (xi,yi)-(xj,yj) */
		dx := xj - xi
		dy := yj - yi
		pxRel := px - xi
		pyRel := py - yi

		/* Cross-product zero => collinear */
		if dx*pyRel-dy*pxRel == 0 {
			/* Checks if the point is within the bounding box of the segment (inclusive) */
			if (px >= xi && px <= xj || px >= xj && px <= xi) && (py >= yi && py <= yj || py >= yj && py <= yi) {
				return true
			}
		}

		/* Even-odd rule: checks if the edge straddles the horizontal ray from the point */
		condY := (yi > py) != (yj > py)
		if !condY {
			continue
		}

		/* Computes intersection x-coordinate of edge with horizontal line y = py */
		ix := float64(xj-xi)*float64(py-yi)/float64(yj-yi) + float64(xi)
		if float64(px) < ix {
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
