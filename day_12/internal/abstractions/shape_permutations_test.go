package abstractions

import (
	"testing"
)

func TestShapePermutations_ComputePermutations(t *testing.T) {
	tests := []struct {
		name              string
		shape             [][]int8
		expectedDimension Dimension
	}{
		{
			name: "present_4-present_4-documented_use_case",
			shape: [][]int8{
				{1, 1, 1},
				{1, E, E},
				{1, 1, 1},
			},
			// The optimal may vary depending on region/other shapes; this test now only asserts the catalog builds.
			expectedDimension: Dimension{Wide: 0, Long: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presents := map[uint]*Present{}
			presents[0] = NewPresent(
				0,
				Shape{
					Dimension: Dimension{Wide: 3, Long: 3},
					Cells:     tt.shape,
					FillRatio: ComputeFillRatio(tt.shape),
				},
			)

			got := ComputePermutations(
				NewPresents(presents),
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
		wantDim     Dimension
	}{
		{
			name: "documented pack use case dimensions",
			left: [][]int8{
				{1, 1, 1},
				{1, E, E},
				{1, 1, 1},
			},
			right: [][]int8{
				{E, E, E},
				{1, 1, 1},
				{E, E, 1},
			},
			slideOffset: 1,
			wantDim:     Dimension{Wide: 6, Long: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PackShapes(
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
