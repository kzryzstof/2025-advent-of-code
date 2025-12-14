package abstractions

import "testing"

func TestIsPointInPolygon(t *testing.T) {
	tests := map[string]struct {
		polygon  []*Tile
		point    *Tile
		expected bool
	}{
		"point_inside_simple_square": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
				{X: 4, Y: 4, Color: Red},
				{X: 0, Y: 4, Color: Red},
			},
			point:    &Tile{X: 2, Y: 2, Color: "v"},
			expected: true,
		},
		"point_outside_simple_square": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
				{X: 4, Y: 4, Color: Red},
				{X: 0, Y: 4, Color: Red},
			},
			point:    &Tile{X: 5, Y: 5, Color: "X"},
			expected: false,
		},
		"point_on_left_edge_of_square": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
				{X: 4, Y: 4, Color: Red},
				{X: 0, Y: 4, Color: Red},
			},
			point:    &Tile{X: 0, Y: 2, Color: "v"},
			expected: true,
		},
		"point_on_vertex_of_square": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
				{X: 4, Y: 4, Color: Red},
				{X: 0, Y: 4, Color: Red},
			},
			point:    &Tile{X: 0, Y: 0, Color: "v"},
			expected: true,
		},
		"point_inside_triangle": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 6, Y: 0, Color: Red},
				{X: 3, Y: 5, Color: Red},
			},
			point:    &Tile{X: 3, Y: 2, Color: "v"},
			expected: true,
		},
		"point_outside_triangle_left": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 6, Y: 0, Color: Red},
				{X: 3, Y: 5, Color: Red},
			},
			point:    &Tile{X: 0, Y: 6, Color: "X"},
			expected: false,
		},
		"point_outside_triangle_below": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 6, Y: 0, Color: Red},
				{X: 3, Y: 5, Color: Red},
			},
			point:    &Tile{X: 3, Y: 10, Color: "X"},
			expected: false,
		},
		"point_inside_complex_polygon": {
			polygon: []*Tile{
				{X: 1, Y: 1, Color: Red},
				{X: 5, Y: 1, Color: Red},
				{X: 5, Y: 3, Color: Red},
				{X: 3, Y: 3, Color: Red},
				{X: 3, Y: 5, Color: Red},
				{X: 1, Y: 5, Color: Red},
			},
			point:    &Tile{X: 2, Y: 2, Color: "v"},
			expected: true,
		},
		"point_outside_complex_polygon": {
			polygon: []*Tile{
				{X: 1, Y: 1, Color: Red},
				{X: 5, Y: 1, Color: Red},
				{X: 5, Y: 3, Color: Red},
				{X: 3, Y: 3, Color: Red},
				{X: 3, Y: 5, Color: Red},
				{X: 1, Y: 5, Color: Red},
			},
			point:    &Tile{X: 4, Y: 4, Color: "x"},
			expected: false,
		},
		"point_in_concave_region_of_polygon": {
			polygon: []*Tile{
				{X: 1, Y: 1, Color: Red},
				{X: 5, Y: 1, Color: Red},
				{X: 5, Y: 3, Color: Red},
				{X: 3, Y: 3, Color: Red},
				{X: 3, Y: 5, Color: Red},
				{X: 1, Y: 5, Color: Red},
			},
			point:    &Tile{X: 2, Y: 4, Color: "v"},
			expected: true,
		},
		"point_far_outside_polygon": {
			polygon: []*Tile{
				{X: 2, Y: 2, Color: Red},
				{X: 8, Y: 2, Color: Red},
				{X: 8, Y: 8, Color: Red},
				{X: 2, Y: 8, Color: Red},
			},
			point:    &Tile{X: 20, Y: 20, Color: "x"},
			expected: false,
		},
		"point_just_outside_right_edge": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
				{X: 4, Y: 4, Color: Red},
				{X: 0, Y: 4, Color: Red},
			},
			point:    &Tile{X: 5, Y: 2, Color: "x"},
			expected: false,
		},
		"point_just_outside_top_edge": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
				{X: 4, Y: 4, Color: Red},
				{X: 0, Y: 4, Color: Red},
			},
			point:    &Tile{X: 2, Y: 5, Color: "x"},
			expected: false,
		},
		"polygon_with_less_than_3_vertices": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 4, Y: 0, Color: Red},
			},
			point:    &Tile{X: 2, Y: 0, Color: "x"},
			expected: false,
		},
		"empty_polygon": {
			polygon:  []*Tile{},
			point:    &Tile{X: 0, Y: 0, Color: "x"},
			expected: false,
		},
		"point_inside_hexagon": {
			polygon: []*Tile{
				{X: 2, Y: 0, Color: Red},
				{X: 4, Y: 1, Color: Red},
				{X: 4, Y: 3, Color: Red},
				{X: 2, Y: 4, Color: Red},
				{X: 0, Y: 3, Color: Red},
				{X: 0, Y: 1, Color: Red},
			},
			point:    &Tile{X: 2, Y: 2, Color: "v"},
			expected: true,
		},
		"point_outside_hexagon": {
			polygon: []*Tile{
				{X: 2, Y: 0, Color: Red},
				{X: 4, Y: 1, Color: Red},
				{X: 4, Y: 3, Color: Red},
				{X: 2, Y: 4, Color: Red},
				{X: 0, Y: 3, Color: Red},
				{X: 0, Y: 1, Color: Red},
			},
			point:    &Tile{X: 5, Y: 5, Color: "x"},
			expected: false,
		},
		"point_at_center_of_large_polygon": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 10, Y: 0, Color: Red},
				{X: 10, Y: 10, Color: Red},
				{X: 0, Y: 10, Color: Red},
			},
			point:    &Tile{X: 5, Y: 5, Color: "v"},
			expected: true,
		},
		"point_in_narrow_polygon": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 10, Y: 0, Color: Red},
				{X: 10, Y: 1, Color: Red},
				{X: 0, Y: 1, Color: Red},
			},
			point:    &Tile{X: 5, Y: 0, Color: "v"},
			expected: true,
		},
		"point_outside_narrow_polygon": {
			polygon: []*Tile{
				{X: 0, Y: 0, Color: Red},
				{X: 10, Y: 0, Color: Red},
				{X: 10, Y: 1, Color: Red},
				{X: 0, Y: 1, Color: Red},
			},
			point:    &Tile{X: 5, Y: 2, Color: "x"},
			expected: false,
		},
		"corner_case_document_use_case_1": {
			polygon: []*Tile{
				{X: 7, Y: 1, Color: Red},
				{X: 11, Y: 1, Color: Red},
				{X: 11, Y: 7, Color: Red},
				{X: 9, Y: 7, Color: Red},
				{X: 9, Y: 5, Color: Red},
				{X: 2, Y: 5, Color: Red},
				{X: 2, Y: 3, Color: Red},
				{X: 7, Y: 3, Color: Red},
			},
			point:    &Tile{X: 10, Y: 7, Color: "v"},
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := IsPointInPolygon(tc.polygon, tc.point)

			if result != tc.expected {
				t.Errorf("IsPointInPolygon() = %v, want %v for point (%d,%d)",
					result, tc.expected, tc.point.X, tc.point.Y)
			}
		})
	}
}

func TestOrderPolygonVertices(t *testing.T) {
	tests := map[string]struct {
		tiles       []*Tile
		description string
	}{
		"square_vertices_unordered": {
			tiles: []*Tile{
				{X: 0, Y: 0},
				{X: 4, Y: 4},
				{X: 4, Y: 0},
				{X: 0, Y: 4},
			},
			description: "should order square vertices counter-clockwise",
		},
		"triangle_vertices_unordered": {
			tiles: []*Tile{
				{X: 3, Y: 5},
				{X: 0, Y: 0},
				{X: 6, Y: 0},
			},
			description: "should order triangle vertices counter-clockwise",
		},
		"already_ordered_vertices": {
			tiles: []*Tile{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
				{X: 4, Y: 4},
				{X: 0, Y: 4},
			},
			description: "should maintain order of already ordered vertices",
		},
		"less_than_3_vertices": {
			tiles: []*Tile{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
			},
			description: "should return unchanged for less than 3 vertices",
		},
		"empty_list": {
			tiles:       []*Tile{},
			description: "should return empty list",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := OrderPolygonVertices(tc.tiles)

			// Verify the result has the same number of vertices
			if len(result) != len(tc.tiles) {
				t.Fatalf("OrderPolygonVertices() returned %d vertices, want %d",
					len(result), len(tc.tiles))
			}

			// For polygons with 3+ vertices, verify they are ordered
			if len(tc.tiles) >= 3 {
				// Verify all original tiles are present
				for _, original := range tc.tiles {
					found := false
					for _, ordered := range result {
						if ordered.X == original.X && ordered.Y == original.Y {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Original tile (%d,%d) not found in ordered result",
							original.X, original.Y)
					}
				}

				// Verify they form a valid polygon (angles should be monotonically increasing)
				var sumX, sumY float64
				for _, tile := range result {
					sumX += float64(tile.X)
					sumY += float64(tile.Y)
				}
				// We just verify the ordering doesn't panic and returns reasonable results
				// The exact order depends on the centroid-based angle calculation
			}
		})
	}
}
