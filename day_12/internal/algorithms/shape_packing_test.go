package algorithms

import (
	"day_12/internal/abstractions"
	"testing"
)

func TestShapePacking_PackShape(t *testing.T) {
	tests := []struct {
		name            string
		region          [][]int8
		shape           [][]int8
		expectedRegion  [][]int8
		expectedSuccess bool
	}{
		{
			name: "given empty region when packing shape then inserted at origin",
			region: [][]int8{
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 1, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, 1, abstractions.E, abstractions.E},
				{1, 1, 1, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			expectedSuccess: true,
		},
		{
			name: "given top-left 2x3 occupied block when packing shape then inserted below occupied area",
			region: [][]int8{
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, 1, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, 1, abstractions.E, abstractions.E},
				{1, 1, 1, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			expectedSuccess: true,
		},
		{
			name: "given left column mostly occupied when packing shape then inserted at first available lower row",
			region: [][]int8{
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, 1, abstractions.E, abstractions.E},
				{abstractions.E, abstractions.E, 1, abstractions.E, abstractions.E},
				{1, 1, 1, abstractions.E, abstractions.E},
			},
			expectedSuccess: true,
		},
		{
			name: "given left column fully occupied when packing shape then inserted shifted right",
			region: [][]int8{
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, 1, 1, 1, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, 1, abstractions.E},
				{abstractions.P, 1, 1, 1, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			expectedSuccess: true,
		},
		{
			name: "given first row fully occupied when packing shape then inserted on next available row",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{1, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, 1, 1, 1, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, 1, abstractions.E},
				{abstractions.P, 1, 1, 1, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
			},
			expectedSuccess: true,
		},
		{
			name: "given only bottom-right pocket available when packing shape then inserted in bottom pocket",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, 1, 1, 1},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, 1},
				{abstractions.P, abstractions.P, 1, 1, 1},
			},
			expectedSuccess: true,
		},
		{
			name: "given no 3x3 placement possible when packing shape then fails",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
				{1, 1, abstractions.E, abstractions.E, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.E, abstractions.E, abstractions.E},
			},
			expectedSuccess: false,
		},
		{
			name: "given central 3x3 cavity when packing shape then inserted in cavity",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, abstractions.E, abstractions.E, abstractions.E, 1},
				{1, abstractions.E, abstractions.E, abstractions.E, 1},
				{1, abstractions.E, abstractions.E, abstractions.E, 1},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, abstractions.E},
				{1, 1, 1, 1, abstractions.E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.P},
				{abstractions.P, 1, 1, 1, abstractions.P},
				{abstractions.P, abstractions.E, abstractions.E, 1, abstractions.P},
				{abstractions.P, 1, 1, 1, abstractions.P},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
				{abstractions.P, abstractions.P, abstractions.P, abstractions.P, abstractions.E},
			},
			expectedSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			succeeded := PackShape(
				tt.region,
				tt.shape,
				false,
			)

			if succeeded != tt.expectedSuccess {
				t.Fatalf("expected success: %t, got %t", tt.expectedSuccess, succeeded)
			}

			for i := 0; i < len(tt.region); i++ {
				for j := 0; j < len(tt.region[i]); j++ {
					if tt.region[i][j] != tt.expectedRegion[i][j] {
						t.Fatalf("mismatch at (%d,%d): got %d, want %d", i, j, tt.region[i][j], tt.expectedRegion[i][j])
					}
				}
			}
		})
	}
}

func TestShapePacking_CombineShapes(t *testing.T) {
	tests := []struct {
		name          string
		fixedShapeID  uint
		fixed         [][]int8
		movingShapeID uint
		moving        [][]int8
		slideOffset   int
		expectedShape [][]int8
	}{
		{
			name:         "given shapes #4 and #5 aligned when slideOffset=0 then merged side-by-side",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			slideOffset: 0,
			expectedShape: [][]int8{
				{4, 4, 4, 5, 5, 5},
				{4, abstractions.E, abstractions.E, abstractions.E, abstractions.E, 5},
				{4, 4, 4, 5, 5, 5},
			},
		},
		{
			name:         "given shapes #4 and #5 aligned when slideOffset=1 then merged with one-row vertical shift",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			slideOffset: 1,
			expectedShape: [][]int8{
				{4, 4, 4, abstractions.E},
				{4, 5, 5, 5},
				{4, 4, 4, 5},
				{abstractions.E, 5, 5, 5},
			},
		},
		{
			name:         "given shapes #4 and #5 aligned when slideOffset=2 then merged with two-row vertical shift",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			slideOffset: 2,
			expectedShape: [][]int8{
				{4, 4, 4, abstractions.E, abstractions.E, abstractions.E},
				{4, abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E},
				{4, 4, 4, 5, 5, 5},
				{abstractions.E, abstractions.E, abstractions.E, abstractions.E, abstractions.E, 5},
				{abstractions.E, abstractions.E, abstractions.E, 5, 5, 5},
			},
		},
		{
			name:         "given shapes #4 and #5 aligned when slideOffset=3 then merged stacked vertically",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
			},
			slideOffset: 3,
			expectedShape: [][]int8{
				{4, 4, 4},
				{4, abstractions.E, abstractions.E},
				{4, 4, 4},
				{5, 5, 5},
				{abstractions.E, abstractions.E, 5},
				{5, 5, 5},
			},
		},
		{
			name:         "given rotated/offset moving shape when slideOffset=0 then merged with diagonal contact",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{abstractions.E, 1, 1},
				{1, 1, 1},
				{1, 1, abstractions.E},
			},
			slideOffset: 0,
			expectedShape: [][]int8{
				{4, 4, 4, abstractions.E, 5, 5},
				{4, abstractions.E, abstractions.E, 5, 5, 5},
				{4, 4, 4, 5, 5, abstractions.E},
			},
		},
		{
			name:         "given fixed shape variant and moving shape variant when slideOffset=2 then merged tightly",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{abstractions.E, abstractions.E, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			slideOffset: 2,
			expectedShape: [][]int8{
				{4, 4, 4},
				{4, 4, 4},
				{4, abstractions.E, 5},
				{5, 5, 5},
				{5, 5, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombineShapes(
				tt.fixedShapeID,
				tt.fixed,
				tt.movingShapeID,
				tt.moving,
				tt.slideOffset,
				false,
			)

			for row := 0; row < len(tt.expectedShape); row++ {
				for col := 0; col < len(tt.expectedShape[row]); col++ {
					if got.Cells[row][col] != tt.expectedShape[row][col] {
						t.Fatalf("cell mismatch at (%d,%d): got %d, want %d", row, col, got.Cells[row][col], tt.expectedShape[row][col])
					}
				}
			}
		})
	}
}
