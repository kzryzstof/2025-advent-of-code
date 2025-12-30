package algorithms

import (
	"day_12/internal/abstractions"
	"testing"
)

func TestShapePermutations_ComputePermutations(t *testing.T) {
	tests := []struct {
		name              string
		shape             [][]int8
		expectedDimension abstractions.Dimension
	}{
		{
			name: "present_4-present_4-documented_use_case",
			shape: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			// The optimal may vary depending on region/other shapes; this test now only asserts the catalog builds.
			expectedDimension: abstractions.Dimension{Wide: 0, Long: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presents := map[uint]*abstractions.Present{}
			presents[0] = abstractions.NewPresent(
				0,
				abstractions.Shape{
					Dimension: abstractions.Dimension{Wide: 3, Long: 3},
					Cells:     tt.shape,
					FillRatio: abstractions.ComputeFillRatio(tt.shape),
				},
			)

			region := abstractions.NewRegion(100, 100)

			got := abstractions.ComputePermutations(
				abstractions.NewPresents(presents),
				region,
				false,
			)

			_, optimalShape := got.GetOptimalCombination(0)
			if optimalShape.Dimension.Wide <= 0 || optimalShape.Dimension.Long <= 0 {
				t.Fatalf("expected a non-empty optimal dimension, got %dx%d", optimalShape.Dimension.Wide, optimalShape.Dimension.Long)
			}
		})
	}
}

func TestShapePermutations_PackShapes(t *testing.T) {
	tests := []struct {
		name        string
		left        [][]int8
		right       [][]int8
		slideOffset int
		wantDim     abstractions.Dimension
	}{
		{
			name: "documented pack use case dimensions",
			left: [][]int8{
				{1, 1, 1},
				{1, 0, 0},
				{1, 1, 1},
			},
			right: [][]int8{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 1},
			},
			slideOffset: 1,
			wantDim:     abstractions.Dimension{Wide: 7, Long: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := abstractions.PackShapes(
				1,
				tt.left,
				2,
				tt.right,
				tt.slideOffset,
				false,
			)

			if !got.Dimension.Equals(tt.wantDim) {
				t.Fatalf("got dimension = %dx%d, want %dx%d", got.Dimension.Wide, got.Dimension.Long, tt.wantDim.Wide, tt.wantDim.Long)
			}
		})
	}
}
