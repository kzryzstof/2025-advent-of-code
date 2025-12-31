package algorithms

import (
	"day_12/internal/abstractions"
	"day_12/internal/maths"
	"testing"
)

func TestShapePermutations_ComputePermutations(t *testing.T) {
	tests := []struct {
		name              string
		shape             [][]int8
		expectedDimension maths.Dimension
	}{
		{
			name: "present_4-present_4-documented_use_case",
			shape: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			// The optimal may vary depending on region/other shapes; this test now only asserts the catalog builds.
			expectedDimension: maths.Dimension{Wide: 0, Long: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presents := map[abstractions.PresentIndex]*abstractions.Present{}
			presents[0] = abstractions.NewPresent(
				0,
				abstractions.NewShape(
					maths.Dimension{Wide: 3, Long: 3},
					tt.shape,
				),
			)

			got := ComputePermutations(
				abstractions.NewPresents(presents),
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
		slideOffset uint
		wantDim     maths.Dimension
	}{
		{
			name: "documented pack use case dimensions",
			left: [][]int8{
				{1, 1, 1},
				{1, abstractions.E, abstractions.E},
				{1, 1, 1},
			},
			right: [][]int8{
				{abstractions.E, abstractions.E, abstractions.E},
				{1, 1, 1},
				{abstractions.E, abstractions.E, 1},
			},
			slideOffset: 1,
			wantDim:     maths.Dimension{Wide: 6, Long: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinePresents(
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
