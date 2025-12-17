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
		"simple_2x2_system": {
			rows: 2,
			cols: 2,
			matrixValues: [][]float64{
				{2, 3},
				{1, -1},
			},
			vectorValues: []float64{6, .5},
			expectedMat: [][]float64{
				{1, -1},
				{0, 1},
			},
			expectedVec: []float64{0.5, 1},
		},
		"other_simple_2x2_system": {
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
		"another_simple_2x2_system": {
			rows: 2,
			cols: 2,
			matrixValues: [][]float64{
				{3, 4},
				{6, 8},
			},
			vectorValues: []float64{12, 24},
			expectedMat: [][]float64{
				{6, 8},
				{0, 0},
			},
			expectedVec: []float64{0, 24},
		},
		"simple_3x3_system": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{2, 1, -1},
				{-3, -1, 2},
				{-2, 1, 2},
			},
			vectorValues: []float64{8, -11, -3},
			expectedMat: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expectedVec: []float64{2, 3, -1},
		},
		"other_simple_3x3_system": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{1, -3, 4},
				{2, -5, 6},
				{-3, 3, 4},
			},
			vectorValues: []float64{3, 0, 6},
			expectedMat: [][]float64{
				{1, -3, 4},
				{0, 1, -2},
				{0, 0, 1},
			},
			expectedVec: []float64{3, 0, 15 / 4},
		},
		"system_with_free_variables_4x6": {
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
				{1, 0, 0, 1, -1, -2},
				{0, 1, 0, 0, 0, 1},
				{0, 0, 1, 1, 0, -1},
				{0, 0, 0, 0, 1, 1},
			},
			expectedVec: []float64{-1, 5, 1, 3},
		},
		"identity_matrix_already_reduced": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			vectorValues: []float64{4, 5, 6},
			expectedMat: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expectedVec: []float64{4, 5, 6},
		},
		"upper_triangular_needs_back_elimination": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{1, 2, 3},
				{0, 1, 4},
				{0, 0, 1},
			},
			vectorValues: []float64{14, 11, 3},
			expectedMat: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expectedVec: []float64{-1, -1, 3},
		},
		"system_requiring_row_swap": {
			rows: 3,
			cols: 3,
			matrixValues: [][]float64{
				{0, 1, 2},
				{1, 2, 3},
				{3, 4, 5},
			},
			vectorValues: []float64{8, 13, 18},
			expectedMat: [][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expectedVec: []float64{1, 2, 3},
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
			Reduce(m, v)

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
