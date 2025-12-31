package algorithms

import (
	"day_12/internal/abstractions"
	"testing"
)

func TestCombinePresents(t *testing.T) {
	tests := []struct {
		name          string
		fixedShapeID  abstractions.PresentIndex
		fixed         [][]int8
		movingShapeID abstractions.PresentIndex
		moving        [][]int8
		slideOffset   uint
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
			got := CombinePresents(
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
