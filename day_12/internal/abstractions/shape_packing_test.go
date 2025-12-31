package abstractions

import "testing"

func TestShapePacking_PackShape(t *testing.T) {
	tests := []struct {
		name            string
		region          [][]int8
		shapeID         uint
		shape           [][]int8
		expectedRegion  [][]int8
		expectedSuccess bool
	}{
		{
			name: "region is empty",
			region: [][]int8{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{5, 5, 5, 0, 0},
				{0, 0, 5, 0, 0},
				{5, 5, 5, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expectedSuccess: true,
		},
		{
			name: "region's first few cols are occupied",
			region: [][]int8{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{5, 5, 5, 0, 0},
				{0, 0, 5, 0, 0},
				{5, 5, 5, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expectedSuccess: true,
		},
		{
			name: "region's columns are occupied",
			region: [][]int8{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{5, 5, 5, 0, 0},
				{0, 0, 5, 0, 0},
				{5, 5, 5, 0, 0},
			},
			expectedSuccess: true,
		},
		{
			name: "region's columns are occupied",
			region: [][]int8{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 5, 5, 5, 0},
				{1, 0, 0, 5, 0},
				{1, 5, 5, 5, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
			},
			expectedSuccess: true,
		},
		{
			name: "first row is occupied",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 5, 5, 5, 0},
				{1, 0, 0, 5, 0},
				{1, 5, 5, 5, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
			},
			expectedSuccess: true,
		},
		{
			name: "bottom corner is empty",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 5, 5, 5},
				{1, 1, 0, 0, 5},
				{1, 1, 5, 5, 5},
			},
			expectedSuccess: true,
		},
		{
			name: "region is full",
			region: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
			},
			shapeID: 5,
			shape: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			expectedRegion: [][]int8{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
			},
			expectedSuccess: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			succeeded := PackShape(
				tt.region,
				tt.shapeID,
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
		wantDim       Dimension
	}{
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			slideOffset: 0,
			wantDim:     Dimension{Wide: 6, Long: 3},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			slideOffset: 1,
			wantDim:     Dimension{Wide: 4, Long: 4},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			slideOffset: 2,
			wantDim:     Dimension{Wide: 6, Long: 5},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{1, 1, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			slideOffset: 3,
			wantDim:     Dimension{Wide: 3, Long: 6},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 0},
			},
			slideOffset: 0,
			wantDim:     Dimension{Wide: 6, Long: 3},
		},
		{
			name:         "documented use case: packing produces expected canvas dimensions",
			fixedShapeID: 4,
			fixed: [][]int8{
				{1, 1, 1},
				{1, 1, 1},
				{1, 0, 0},
			},
			movingShapeID: 5,
			moving: [][]int8{
				{0, 0, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			slideOffset: 2,
			wantDim:     Dimension{Wide: 3, Long: 5},
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

			if !got.Dimension.Equals(tt.wantDim) {
				t.Fatalf(
					"dimension mismatch: got %dx%d, want %dx%d",
					got.Dimension.Wide,
					got.Dimension.Long,
					tt.wantDim.Wide,
					tt.wantDim.Long,
				)
			}
		})
	}
}
