package abstractions

import "testing"

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
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 1, E, E},
				{E, E, 1, E, E},
				{1, 1, 1, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
			},
			expectedSuccess: true,
		},
		{
			name: "region's first few cols are occupied",
			region: [][]int8{
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, E, E, E},
				{P, P, E, E, E},
				{P, P, E, E, E},
				{1, 1, 1, E, E},
				{E, E, 1, E, E},
				{1, 1, 1, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
			},
			expectedSuccess: true,
		},
		{
			name: "region's columns are occupied",
			region: [][]int8{
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
				{E, E, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, E, E, E},
				{P, P, E, E, E},
				{P, P, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{1, 1, 1, E, E},
				{E, E, 1, E, E},
				{1, 1, 1, E, E},
			},
			expectedSuccess: true,
		},
		{
			name: "region's columns are occupied",
			region: [][]int8{
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, E, E, E},
				{P, P, E, E, E},
				{P, P, E, E, E},
				{P, 1, 1, 1, E},
				{P, E, E, 1, E},
				{P, 1, 1, 1, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
			},
			expectedSuccess: true,
		},
		{
			name: "first row is occupied",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
				{1, E, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, P, P, P},
				{P, P, E, E, E},
				{P, P, E, E, E},
				{P, 1, 1, 1, E},
				{P, E, E, 1, E},
				{P, 1, 1, 1, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
				{P, E, E, E, E},
			},
			expectedSuccess: true,
		},
		{
			name: "bottom corner is empty",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, P, P, P},
				{P, P, P, P, E},
				{P, P, P, P, E},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, 1, 1, 1},
				{P, P, E, E, 1},
				{P, P, 1, 1, 1},
			},
			expectedSuccess: true,
		},
		{
			name: "region is full",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, E, E},
				{1, 1, E, E, E},
				{1, 1, E, E, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, P, P, P},
				{P, P, P, P, E},
				{P, P, P, P, E},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, E, E},
				{P, P, E, E, E},
				{P, P, E, E, E},
			},
			expectedSuccess: false,
		},
		{
			name: "region has place in the middle",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, E, E, E, 1},
				{1, E, E, E, 1},
				{1, E, E, E, 1},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, E},
				{1, 1, 1, 1, E},
			},
			shape: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{P, P, P, P, P},
				{P, P, P, P, E},
				{P, P, P, P, E},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, P, P, P, P},
				{P, 1, 1, 1, P},
				{P, E, E, 1, P},
				{P, 1, 1, 1, P},
				{P, P, P, P, E},
				{P, P, P, P, E},
				{P, P, P, P, E},
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
				{1, E, E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			slideOffset: 0,
			expectedShape: [][]int8{
				{4, 4, 4, 5, 5, 5},
				{4, E, E, E, E, 5},
				{4, 4, 4, 5, 5, 5},
			},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, E, E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			slideOffset: 1,
			expectedShape: [][]int8{
				{4, 4, 4, E},
				{4, 5, 5, 5},
				{4, 4, 4, 5},
				{E, 5, 5, 5},
			},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, E, E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			slideOffset: 2,
			expectedShape: [][]int8{
				{4, 4, 4, E, E, E},
				{4, E, E, E, E, E},
				{4, 4, 4, 5, 5, 5},
				{E, E, E, E, E, 5},
				{E, E, E, 5, 5, 5},
			},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, E, E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{E, E, 1},
				{1, 1, 1},
			},
			slideOffset: 3,
			expectedShape: [][]int8{
				{4, 4, 4},
				{4, E, E},
				{4, 4, 4},
				{5, 5, 5},
				{E, E, 5},
				{5, 5, 5},
			},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, E, E},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{E, 1, 1},
				{1, 1, 1},
				{1, 1, E},
			},
			slideOffset: 0,
			expectedShape: [][]int8{
				{4, 4, 4, E, 5, 5},
				{4, E, E, 5, 5, 5},
				{4, 4, 4, 5, 5, E},
			},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 1, 1},
				{1, E, E},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{E, E, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			slideOffset: 2,
			expectedShape: [][]int8{
				{4, 4, 4},
				{4, 4, 4},
				{4, E, 5},
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
