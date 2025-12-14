package app

import (
	"day_10/internal/abstractions"
	"math"
	"testing"
)

type expectedPair struct {
	indexA   int
	indexB   int
	distance float64
}

func TestPair(t *testing.T) {
	tests := []struct {
		name          string
		junctionBoxes []*abstractions.JunctionBox
		expectedPairs []expectedPair
	}{
		{
			name:          "Empty_List_Returns_No_Pairs",
			junctionBoxes: []*abstractions.JunctionBox{},
			expectedPairs: []expectedPair{},
		},
		{
			name: "Single_Junction_Box_Returns_No_Pairs",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
			},
			expectedPairs: []expectedPair{},
		},
		{
			name: "Two_Junction_Boxes_Return_One_Pair",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 1.0},
			},
		},
		{
			name: "Three_Junction_Boxes_Return_Three_Pairs",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 2, Y: 0, Z: 0}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 1.0},
				{indexA: 0, indexB: 2, distance: 2.0},
				{indexA: 1, indexB: 2, distance: 1.0},
			},
		},
		{
			name: "Four_Junction_Boxes_Return_Six_Pairs",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 0, Y: 1, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 1, Z: 0}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 1.0},
				{indexA: 0, indexB: 2, distance: 1.0},
				{indexA: 0, indexB: 3, distance: math.Sqrt(2)},
				{indexA: 1, indexB: 2, distance: math.Sqrt(2)},
				{indexA: 1, indexB: 3, distance: 1.0},
				{indexA: 2, indexB: 3, distance: 1.0},
			},
		},
		{
			name: "Five_Junction_Boxes_Return_Ten_Pairs",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 2, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 3, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 4, Y: 0, Z: 0}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 1.0},
				{indexA: 0, indexB: 2, distance: 2.0},
				{indexA: 0, indexB: 3, distance: 3.0},
				{indexA: 0, indexB: 4, distance: 4.0},
				{indexA: 1, indexB: 2, distance: 1.0},
				{indexA: 1, indexB: 3, distance: 2.0},
				{indexA: 1, indexB: 4, distance: 3.0},
				{indexA: 2, indexB: 3, distance: 1.0},
				{indexA: 2, indexB: 4, distance: 2.0},
				{indexA: 3, indexB: 4, distance: 1.0},
			},
		},
		{
			name: "Pairs_Have_Correct_3D_Distances",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 3, Y: 4, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 1, Z: 1}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 5.0},
				{indexA: 0, indexB: 2, distance: math.Sqrt(3)},
				{indexA: 1, indexB: 2, distance: math.Sqrt(14)},
			},
		},
		{
			name: "No_Duplicate_Pairs",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 2, Y: 0, Z: 0}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 1.0},
				{indexA: 0, indexB: 2, distance: 2.0},
				{indexA: 1, indexB: 2, distance: 1.0},
			},
		},
		{
			name: "Large_Set_Returns_Correct_Count",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 2, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 3, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 4, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 5, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 6, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 7, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 8, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 9, Y: 0, Z: 0}},
			},
			expectedPairs: []expectedPair{
				// Only verify count with C(10,2) = 45 pairs
				// Using empty slice to skip individual pair validation
			},
		},
		{
			name: "Same_Position_Junction_Boxes_Have_Zero_Distance",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 5, Y: 5, Z: 5}},
				{Position: abstractions.Position{X: 5, Y: 5, Z: 5}},
			},
			expectedPairs: []expectedPair{
				{indexA: 0, indexB: 1, distance: 0.0},
			},
		},
		{
			name: "Real_World_Example_From_Test_Data",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 162, Y: 817, Z: 812}},
				{Position: abstractions.Position{X: 57, Y: 618, Z: 57}},
				{Position: abstractions.Position{X: 906, Y: 360, Z: 560}},
				{Position: abstractions.Position{X: 592, Y: 479, Z: 940}},
				{Position: abstractions.Position{X: 352, Y: 342, Z: 300}},
			},
			expectedPairs: []expectedPair{
				// Only verify count - C(5,2) = 10 pairs
			},
		},
		{
			name: "Junction_Boxes_In_3D_Cube",
			junctionBoxes: []*abstractions.JunctionBox{
				{Position: abstractions.Position{X: 0, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 0}},
				{Position: abstractions.Position{X: 0, Y: 1, Z: 0}},
				{Position: abstractions.Position{X: 1, Y: 1, Z: 0}},
				{Position: abstractions.Position{X: 0, Y: 0, Z: 1}},
				{Position: abstractions.Position{X: 1, Y: 0, Z: 1}},
				{Position: abstractions.Position{X: 0, Y: 1, Z: 1}},
				{Position: abstractions.Position{X: 1, Y: 1, Z: 1}},
			},
			expectedPairs: []expectedPair{
				// Only verify count - C(8,2) = 28 pairs
				// In a unit cube: 12 edges, 12 face diagonals, 4 space diagonals
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPairs := Pair(tt.junctionBoxes)

			if actualPairs == nil {
				t.Fatal("Expected non-nil pairs slice")
			}

			// Verify the combinatorics formula: C(n,2) = n*(n-1)/2
			n := len(tt.junctionBoxes)
			expectedCount := n * (n - 1) / 2
			if len(actualPairs) != expectedCount {
				t.Errorf("Combinatorics check failed: C(%d,2)=%d, got %d", n, expectedCount, len(actualPairs))
			}

			// If expectedPairs is empty but expectedCount > 0, only verify count
			if len(tt.expectedPairs) == 0 && expectedCount > 0 {
				if len(actualPairs) != expectedCount {
					t.Errorf("Expected %d pairs, got %d", expectedCount, len(actualPairs))
				}
				return
			}

			// Compare expected pairs with actual pairs
			if len(actualPairs) != len(tt.expectedPairs) {
				t.Errorf("Expected %d pairs, got %d", len(tt.expectedPairs), len(actualPairs))
			}

			tolerance := 0.000001
			for i, expected := range tt.expectedPairs {
				if i >= len(actualPairs) {
					t.Errorf("Expected pair %d not found in actual pairs", i)
					continue
				}

				actual := actualPairs[i]

				// Verify junction boxes match by index
				if actual.A != tt.junctionBoxes[expected.indexA] {
					t.Errorf("Pair %d: Expected A to be junction box at index %d", i, expected.indexA)
				}
				if actual.B != tt.junctionBoxes[expected.indexB] {
					t.Errorf("Pair %d: Expected B to be junction box at index %d", i, expected.indexB)
				}

				// Verify distance with tolerance
				if math.Abs(actual.Distance-expected.distance) > tolerance {
					t.Errorf("Pair %d: Expected distance %f, got %f", i, expected.distance, actual.Distance)
				}

				// Verify junction boxes are non-nil
				if actual.A == nil || actual.B == nil {
					t.Errorf("Pair %d: Junction boxes should be non-nil", i)
				}

				// Verify no duplicate pairs
				for j := 0; j < i; j++ {
					other := actualPairs[j]
					if (actual.A == other.A && actual.B == other.B) ||
						(actual.A == other.B && actual.B == other.A) {
						t.Errorf("Pair %d is duplicate of pair %d", i, j)
					}
				}
			}
		})
	}
}
