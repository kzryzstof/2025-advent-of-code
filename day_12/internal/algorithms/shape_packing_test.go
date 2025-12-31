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
			name: "region is empty",
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
			name: "region's first few cols are occupied",
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
			name: "region's columns are occupied",
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
			name: "region's columns are occupied",
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
			name: "first row is occupied",
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
			name: "bottom corner is empty",
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
			name: "region is full",
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
			name: "region has place in the middle",
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
				true,
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

func TestShapePacking_PackShapes(t *testing.T) {
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
			name:         "documented use case: packing produces expected canvas dimensions",
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
			name:         "documented use case: packing produces expected canvas dimensions",
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
			name:         "documented use case: packing produces expected canvas dimensions",
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
			name:         "documented use case: packing produces expected canvas dimensions",
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
			name:         "documented use case: packing produces expected canvas dimensions",
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
			name:         "documented use case: packing produces expected canvas dimensions",
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
			got := PackShapes(
				tt.fixedShapeID,
				tt.fixed,
				tt.movingShapeID,
				tt.moving,
				tt.slideOffset,
				true,
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
