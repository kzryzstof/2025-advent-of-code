package algorithms

import (
	"day_12/internal/abstractions"
	"testing"
)

func TestPack_Empty3x3ShapesPackRight(t *testing.T) {
	tests := []struct {
		name              string
		shape             [][]byte
		otherShape        [][]byte
		packDir           abstractions.Direction
		expectedDimension abstractions.Dimension
	}{
		{
			name: "simple 3x3 columns touching when packed right",
			// shape has a vertical bar in the first column: (0,0),(1,0),(2,0) = 1
			shape: [][]byte{
				{1, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
			// otherShape has a vertical bar in the last column: (0,2),(1,2),(2,2) = 1
			otherShape: [][]byte{
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
			packDir:           packToLeft,
			expectedDimension: abstractions.Dimension{Wide: 2, Long: 3},
		},
		{
			name: "simple 3x3 columns touching when packed from the right",
			// shape has a vertical bar in the first column: (0,0),(1,0),(2,0) = 1
			shape: [][]byte{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
			// otherShape has a vertical bar in the last column: (0,2),(1,2),(2,2) = 1
			otherShape: [][]byte{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			packDir:           packToLeft,
			expectedDimension: abstractions.Dimension{Wide: 6, Long: 3},
		},
		{
			name: "3x3 shapes that can be packed from the right",
			// shape has a vertical bar in the first column: (0,0),(1,0),(2,0) = 1
			shape: [][]byte{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
			// otherShape has a vertical bar in the last column: (0,2),(1,2),(2,2) = 1
			otherShape: [][]byte{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
			packDir:           packToLeft,
			expectedDimension: abstractions.Dimension{Wide: 4, Long: 3},
		},
		{
			name: "simple 3x3 columns touching when packed from the bottom",
			// shape has a vertical bar in the first column: (0,0),(1,0),(2,0) = 1
			shape: [][]byte{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
			// otherShape has a vertical bar in the last column: (0,2),(1,2),(2,2) = 1
			otherShape: [][]byte{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			packDir:           packUp,
			expectedDimension: abstractions.Dimension{Wide: 3, Long: 6},
		},
		{
			name: "present_0-present_2-documented_use_case",
			// shape has a vertical bar in the first column: (0,0),(1,0),(2,0) = 1
			shape: [][]byte{
				{0, 1, 1},
				{1, 1, 0},
				{1, 0, 0},
			},
			// otherShape has a vertical bar in the last column: (0,2),(1,2),(2,2) = 1
			otherShape: [][]byte{
				{0, 0, 1},
				{0, 1, 1},
				{1, 1, 1},
			},
			packDir:           packToLeft,
			expectedDimension: abstractions.Dimension{Wide: 4, Long: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pack(
				tt.shape,
				tt.otherShape,
				tt.packDir,
			)

			if !got.Equals(tt.expectedDimension) {
				t.Fatalf("got dimension = %dx%d, want %dx%d", got.Wide, got.Long, tt.expectedDimension.Wide, tt.expectedDimension.Long)
			}
		})
	}
}
