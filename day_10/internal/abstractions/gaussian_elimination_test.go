package abstractions

import (
	"math"
	"testing"
)

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		rows         int
		cols         int
		matrixValues [][]float64
		vectorValues []float64
		expectedMat  [][]float64
		expectedVec  []float64
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
			expectedVec: []float64{0.5, 4},
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
			expectedVec: []float64{0.5, 4},
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
			expectedVec: []float64{0.5, 4},
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
			expectedVec: []float64{3, 0, 3.75},
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
				{0, 1, -12},
				{0, 0, 1},
			},
			expectedVec: []float64{8, -15, 1},
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
			expectedVec: []float64{1, 0, 0},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Create matrix and vector
			m := NewMatrix(tc.cols, tc.rows)
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
			Reduce(&AugmentedMatrix{m, v})

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
		})
	}
}

// floatEquals checks if two floats are equal within a tolerance
func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
