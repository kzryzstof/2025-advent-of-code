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
			expectedSolution: []float64{1, 5, 0, 1, 3, 0},
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
		"machine.04-from_input": {
			matrixValues: [][]float64{
				{0, 1, 0, 12},
				{1, 1, 0, 29},
				{0, 0, 1, 128},
				{1, 1, 0, 29},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 17},
				{0, 1, 0, 12},
				{0, 0, 1, 128},
				{0, 0, 0, 0},
			},
			expectedSolution: []float64{17, 12, 128},
		},
		"machine.05-from_input": {
			matrixValues: [][]float64{
				{1, 1, 1, 1, 37},
				{1, 0, 1, 0, 4},
				{0, 0, 1, 1, 21},
				{0, 1, 0, 1, 33},
			},
			expectedMat: [][]float64{
				{1, 0, 0, -1, -17},
				{0, 1, 0, 1, 33},
				{0, 0, 1, 1, 21},
				{0, 0, 0, 0, 0},
			},
			expectedSolution: []float64{0, 16, 4, 17},
		},
		"machine-06-from_input": {
			matrixValues: [][]float64{
				{1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 82},
				{1, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 77},
				{1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 28},
				{1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 71},
				{0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 74},
				{1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 79},
				{1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 70},
				{1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 50},
				{1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 82},
				{0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 88},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 12},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0.3333, 20.6667},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0.3333, 3.6667},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, -0.6667, 12.6667},
				{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 20},
				{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, -0.6667, 13.6667},
				{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0.3333, 17.6667},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, -1, 0, 0, -4},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0.3333, 1.6667},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0.6667, 19.3333},
			},
			expectedSolution: []float64{8, 20.6667, 3.6667, 12.6667, 20, 13.6667, 17.6667, 0, 1.6667, 4, 19.3333, 0},
		},
		"machine.09-from_input": {
			matrixValues: [][]float64{
				{1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 72},
				{0, 0, 1, 0, 1, 1, 0, 1, 1, 0, 59},
				{0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 15},
				{0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 37},
				{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 47},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 16},
				{0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 48},
				{0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 32},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 10},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 7},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 16},
				{0, 0, 0, 0, 1, 0, -1, 0, 0, 0, -2},
				{0, 0, 0, 0, 0, 1, 1, 0, 0, -1, 25},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 15},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, -1, 14},
			},
			expectedSolution: []float64{10, 0, 7, 16, 0, 23, 2, 15, 14, 0},
		},
		"machine-33-documented_use-case": {
			matrixValues: [][]float64{
				{1, 0, 1, 1, 0, 1, 0, 0, 33},
				{1, 1, 1, 0, 1, 0, 1, 0, 60},
				{1, 1, 0, 1, 1, 0, 0, 1, 46},
				{1, 1, 1, 1, 0, 0, 0, 0, 48},
				{0, 1, 0, 0, 1, 1, 1, 0, 30},
				{0, 1, 1, 1, 0, 0, 0, 0, 37},
				{1, 1, 1, 0, 1, 0, 0, 0, 58},
			},
			expectedMat: [][]float64{
				{1, 0, 0, 0, 0, 0, 0, 0, 11},
				{0, 1, 0, 0, 0, 0, 0, -0.3333, 13.6667},
				{0, 0, 1, 0, 0, 0, 0, -0.3333, 17.6667},
				{0, 0, 0, 1, 0, 0, 0, 0.6667, 5.6667},
				{0, 0, 0, 0, 1, 0, 0, 0.6667, 15.6667},
				{0, 0, 0, 0, 0, 1, 0, -0.3333, -1.3333},
				{0, 0, 0, 0, 0, 0, 1, 0, 2},
			},
			expectedSolution: []float64{11, 15, 19, 3, 13, 0, 2, 4},
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
			actualSolution.Print()

			fmt.Println("Expected solution")
			PrintSlice(tc.expectedSolution)

			if actualSolution.Count() != uint(len(tc.expectedSolution)) {
				t.Errorf("Solution length: expected %d, got %d", len(tc.expectedSolution), actualSolution.Count())
			} else {
				for i := 0; i < len(tc.expectedSolution); i++ {
					expected := tc.expectedSolution[i]
					actual := actualSolution.GetValue(VariableNumber(i + 1))
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
