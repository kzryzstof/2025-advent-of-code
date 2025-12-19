package abstractions

import (
	"fmt"
	"math"
	"testing"
)

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		matrixValues     [][]float64
		expectedMat      [][]float64
		expectedSolution []float64
	}{
		/* https://ksuweb.kennesaw.edu/~plaval/previous%20semesters/Fall2001/m3261_01/dm1.pdf*/
		"1.0-paper": {
			matrixValues: [][]float64{
				{1, 1, 0, 3, 4},
				{2, 1, -1, 1, 1},
				{3, -1, -1, 2, -3},
				{-1, 2, 3, -1, 4},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 0, -1},
				{0, 1, 0, 0, 2},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			},
			expectedSolution: []float64{-1, 2, 0, 1},
		},
		"3.1-documented_use-case": {
			matrixValues: [][]float64{
				{0, 0, 0, 0, 1, 1, 3},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 1, 0, 4},
				{1, 1, 0, 1, 0, 0, 7},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 1, 0, -1, 2},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 0, -1, 1},
				{0, 0, 0, 0, 1, 1, 3},
			},
			expectedSolution: []float64{1, 3, 0, 3, 1, 2},
		},
		"3.2-documented_use-case": {
			matrixValues: [][]float64{
				{1, 0, 1, 1, 0, 7},
				{0, 0, 0, 1, 1, 5},
				{1, 1, 0, 1, 1, 12},
				{1, 1, 0, 0, 1, 7},
				{1, 0, 1, 0, 1, 2},
			},
			expectedMat: [][]float64{
				{1, 0, 1, 0, 0, 2},
				{0, 1, -1, 0, 0, 5},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 5},
				{0, 0, 0, 0, 1, 0},
			},
			expectedSolution: []float64{2, 5, 0, 5, 0},
		},
		"3.3-documented_use-case": {
			matrixValues: [][]float64{
				{1, 1, 1, 0, 10},
				{1, 0, 1, 1, 11},
				{1, 0, 1, 1, 11},
				{1, 1, 0, 0, 5},
				{1, 1, 1, 0, 10},
				{0, 0, 1, 0, 5},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 1, 6},
				{0, 1, 0, -1, -1},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expectedSolution: []float64{5, 0, 5, 1},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			// Create matrix and vector
			m := FromSlice(tc.matrixValues)

			// Run the reduction
			rref := ToReducedRowEchelonForm(&AugmentedMatrix{m}, true)

			actualRref := rref.Get()

			// Verify matrix results
			for row := 0; row < m.Rows(); row++ {
				for col := 0; col < m.Cols(); col++ {
					expected := tc.expectedMat[row][col]
					actual := actualRref.Get(row, col)
					if !floatEquals(expected, actual, 0.001) {
						t.Errorf("Matrix[%d][%d]: expected %.4f, got %.4f", row, col, expected, actual)
					}
				}
			}

			actualSolution := rref.Solve(true)

			fmt.Println("Actual Solution")
			PrintSlice(actualSolution)

			fmt.Println("Expected solution")
			PrintSlice(tc.expectedSolution)

			if len(actualSolution) != len(tc.expectedSolution) {
				t.Errorf("Solution length: expected %d, got %d", len(tc.expectedSolution), len(actualSolution))
			} else {
				for i := 0; i < len(tc.expectedSolution); i++ {
					expected := tc.expectedSolution[i]
					actual := actualSolution[i]
					if !floatEquals(expected, actual, 0.001) {
						t.Errorf("Solution[%d]: expected %.4f, got %.4f", i, expected, actual)
					}
				}
			}
		})
	}
}

// floatEquals checks if two floats are equal within a tolerance
func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
