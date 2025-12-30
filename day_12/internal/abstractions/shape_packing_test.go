package abstractions

import "testing"

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
