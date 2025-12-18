package abstractions

import (
	"fmt"
	"math"
	"testing"
)

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		rows             int
		cols             int
		matrixValues     [][]float64
		vectorValues     []float64
		expectedMat      [][]float64
		expectedVec      []float64
		expectedSolution []float64
	}{
		"1.1-simple_2x2_system": {
			rows: 2,
			cols: 2,
			matrixValues: [][]float64{
				{2, 1},
				{4, 2},
			},
			vectorValues: []float64{1, 6},
			expectedMat: [][]float64{
				{1, .5},
				{0, 0},
			},
			expectedVec:      []float64{0.5, 4},
			expectedSolution: []float64{0, 1},
		},
		"1.2-simple_2x2_system": {
			rows: 2,
			cols: 2,
			matrixValues: [][]float64{
				{2, 1},
				{4, 2},
			},
			vectorValues: []float64{1, 6},
			expectedMat: [][]float64{
				{1, .5},
				{0, 0},
			},
			expectedVec:      []float64{0.5, 4},
			expectedSolution: []float64{0, 1},
		},
		"1.3-simple_2x2_system": {
			rows: 2,
			cols: 2,
			matrixValues: [][]float64{
				{2, 1},
				{4, 2},
			},
			vectorValues: []float64{1, 6},
			expectedMat: [][]float64{
				{1, .5},
				{0, 0},
			},
			expectedVec:      []float64{0.5, 4},
			expectedSolution: []float64{0, 1},
		},
		"2.1-simple_3x3_system": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{1, -3, 4},
				{2, -5, 6},
				{-3, 3, 4},
			},
			vectorValues: []float64{3, 6, 6},
			expectedMat: [][]float64{
				{1, -3, 4},
				{0, 1, -2},
				{0, 0, 1},
			},
			expectedVec:      []float64{3, 0, 3.75},
			expectedSolution: []float64{10.5, 7.5, 3.75},
		},
		"2.2-simple_3x3_system": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{1, -1, 1},
				{2, 3, -1},
				{3, -2, -9},
			},
			vectorValues: []float64{8, -2, 9},
			expectedMat: [][]float64{
				{1, -1, 1},
				{0, 1, -0.6},
				{0, 0, 1},
			},
			expectedVec:      []float64{8, -3.6, 1},
			expectedSolution: []float64{4, -3, 1},
		},
		"2.3-simple_3x3_system": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{-1, -2, 1},
				{2, 3, 0},
				{0, 1, -2},
			},
			vectorValues: []float64{-1, 2, 0},
			expectedMat: [][]float64{
				{1, 2, -1},
				{0, 1, -2},
				{0, 0, 0},
			},
			expectedVec:      []float64{1, 0, 0},
			expectedSolution: []float64{-2, 2, 1},
		},
		"3.1-documented_use-case": {
			rows: 4,
			cols: 6,
			matrixValues: [][]float64{
				{0, 0, 0, 0, 1, 1},
				{0, 1, 0, 0, 0, 1},
				{0, 0, 1, 1, 1, 0},
				{1, 1, 0, 1, 0, 0},
			},
			vectorValues: []float64{3, 5, 4, 7},
			expectedMat: [][]float64{
				{1, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 0, 1},
				{0, 0, 1, 1, 1, 0},
				{0, 0, 0, 0, 1, 1},
			},
			expectedVec:      []float64{7, 5, 4, 3},
			expectedSolution: []float64{2, 4, 2, 1, 1, 1},
		},
		"3.2-documented_use-case": {
			rows: 5,
			cols: 5,
			matrixValues: [][]float64{
				{1, 0, 1, 1, 0},
				{0, 0, 0, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 0, 0, 1},
				{1, 0, 1, 0, 1},
			},
			vectorValues: []float64{7, 5, 12, 7, 2},
			expectedMat: [][]float64{
				{1, 0, 1, 1, 0},
				{0, 1, -1, 0, 1},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1},
			},
			expectedVec:      []float64{7, 5, 5, 5, 0},
			expectedSolution: []float64{2, 5, 0, 5, 0},
		},
		"3.3-documented_use-case": {
			rows: 6,
			cols: 4,
			matrixValues: [][]float64{
				{1, 1, 1, 0},
				{1, 0, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 0, 1, 0},
			},
			vectorValues: []float64{10, 11, 11, 5, 10, 5},
			expectedMat: [][]float64{
				{1, 1, 1, 0},
				{0, 1, 0, -1},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expectedVec:      []float64{10, -1, 5, 0, 0, 0},
			expectedSolution: []float64{5, 0, 5, 1},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Create matrix and vector
			m := NewMatrix(tc.rows, tc.cols)
			v := NewVector(tc.rows)

			// Fill matrix with test values
			for row := 0; row < tc.rows; row++ {
				for col := 0; col < tc.cols; col++ {
					m.Set(row, col, tc.matrixValues[row][col])
				}
			}

			// Fill vector with test values
			for i := 0; i < tc.rows; i++ {
				v.Set(i, tc.vectorValues[i])
			}

			// Run the reduction
			actualSolution := Reduce(&AugmentedMatrix{m, v})

			// Verify matrix results
			for row := 0; row < tc.rows; row++ {
				for col := 0; col < tc.cols; col++ {
					expected := tc.expectedMat[row][col]
					actual := m.Get(row, col)
					if !floatEquals(expected, actual, 0.001) {
						t.Errorf("Matrix[%d][%d]: expected %.4f, got %.4f", row, col, expected, actual)
					}
				}
			}

			// Verify vector results
			for i := 0; i < tc.rows; i++ {
				expected := tc.expectedVec[i]
				actual := v.Get(i)
				if !floatEquals(expected, actual, 0.001) {
					t.Errorf("Vector[%d]: expected %.4f, got %.4f", i, expected, actual)
				}
			}

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
