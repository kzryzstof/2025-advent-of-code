package algorithms

import (
	"day_10/internal/abstractions"
	"testing"
)

func TestHermiteNormalForm_Solve_UniqueOnlyFree(t *testing.T) {
	tests := map[string]struct {
		matrixValues     [][]int64
		expectedSolution []int64
		hasFreeVariables bool
	}{
		"unique_solution_simple": {
			matrixValues: [][]int64{
				{1, 0, 0, 0, 3},
				{0, 1, 0, 0, 10},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 1, 2},
			},
			expectedSolution: []int64{3, 10, 5, 2},
			hasFreeVariables: false,
		},
		"unique_solution_with_dependencies": {
			matrixValues: [][]int64{
				{1, 0, -1, -2, -3},
				{0, 1, 1, 5, 7},
				{0, 0, 3, 0, 0},
				{0, 0, 0, 13, 13},
			},
			expectedSolution: []int64{-1, 2, 0, 1},
			hasFreeVariables: false,
		},
		"unique_solution_mixed_order": {
			matrixValues: [][]int64{
				{1, 0, 0, 1, 0, -1, 2},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 0, -1, 1},
				{0, 0, 0, 0, 1, 1, 3},
			},
			expectedSolution: []int64{1, 5, 0, 1, 3, 0},
			hasFreeVariables: false,
		},
		"one_free_variable": {
			matrixValues: [][]int64{
				{1, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 1, 0, 0, 0, 0, 0, 0, 10},
				{0, 0, 1, 0, 0, -1, 0, 0, -7},
				{0, 0, 0, 1, 0, 1, 0, 0, 26},
				{0, 0, 0, 0, 1, 0, 0, 0, 5},
				{0, 0, 0, 0, 0, 0, 1, 0, 19},
				{0, 0, 0, 0, 0, 0, 0, 1, 2},
				{0, 0, 0, 0, 0, 0, 0, 2, 4},
			},
			expectedSolution: []int64{3, 10, 0, 19, 5, 7, 19, 2},
			hasFreeVariables: true,
		},
		"free_variable_eight_vars": {
			matrixValues: [][]int64{
				{1, 0, 0, 0, 0, 1, 0, 0, 33},
				{0, 1, 0, 0, 0, 1, 0, 0, 31},
				{0, 0, 1, 0, 0, 0, 0, 0, 178},
				{0, 0, 0, 1, 0, 1, 0, 0, 17},
				{0, 0, 0, 0, 1, 1, 0, 0, 25},
				{0, 0, 0, 0, 0, 2, 0, 0, 30},
				{0, 0, 0, 0, 0, 0, 1, 0, 3},
				{0, 0, 0, 0, 0, 0, 0, 1, 20},
			},
			expectedSolution: []int64{18, 16, 178, 2, 10, 15, 3, 20},
			hasFreeVariables: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			matrix := abstractions.FromSlice(tc.matrixValues)
			hnf := NewHermiteNormalForm(matrix)

			solution := hnf.Solve(false)

			if solution == nil {
				t.Fatal("Expected solution, got nil")
			}

			if len(solution.Get()) != len(tc.expectedSolution) {
				t.Errorf("Expected %d variables, got %d", len(tc.expectedSolution), len(solution.Get()))
			}

			for i, expected := range tc.expectedSolution {
				variableNum := abstractions.VariableNumber(i + 1)
				actual := solution.GetValue(variableNum)
				if actual != expected {
					t.Errorf("Variable x%d: expected %d, got %d", variableNum, expected, actual)
				}
			}

			// Verify the solution satisfies all equations
			verifySolution(t, matrix, solution)
		})
	}
}

func TestHermiteNormalForm_findFreeVariables(t *testing.T) {
	tests := map[string]struct {
		matrixValues         [][]int64
		expectedFreeVarCount int
		expectedFreeVars     []abstractions.VariableNumber
	}{
		"no_free_variables": {
			matrixValues: [][]int64{
				{1, 0, 0, 0, 3},
				{0, 1, 0, 0, 10},
				{0, 0, 1, 0, 5},
				{0, 0, 0, 1, 2},
			},
			expectedFreeVarCount: 0,
			expectedFreeVars:     []abstractions.VariableNumber{},
		},
		"one_free_variable_column_5": {
			matrixValues: [][]int64{
				{1, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 1, 0, 0, 0, 0, 0, 0, 10},
				{0, 0, 1, 0, 0, -1, 0, 0, -7},
				{0, 0, 0, 1, 0, 1, 0, 0, 26},
				{0, 0, 0, 0, 1, 0, 0, 0, 5},
				{0, 0, 0, 0, 0, 0, 1, 0, 19},
				{0, 0, 0, 0, 0, 0, 0, 1, 2},
				{0, 0, 0, 0, 0, 0, 0, 2, 4},
			},
			expectedFreeVarCount: 1,
			expectedFreeVars:     []abstractions.VariableNumber{6},
		},
		"one_free_variable_column_6": {
			matrixValues: [][]int64{
				{1, 0, 0, 0, 0, 0, 0, 0, -1, 0, -5},
				{0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 31},
				{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 13},
				{0, 0, 0, 1, 0, -1, 0, 0, 1, 0, 166},
				{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 21},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 8},
				{0, 0, 0, 0, 0, 0, 1, 0, -1, 0, -2},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 23},
				{0, 0, 0, 0, 0, 0, 0, 0, 2, -1, 16},
			},
			expectedFreeVarCount: 1,
			expectedFreeVars:     []abstractions.VariableNumber{6},
		},
		"multiple_free_variables": {
			matrixValues: [][]int64{
				{1, 0, 1, 0, 0, 2},
				{0, 1, -1, 0, 0, 5},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 5},
				{0, 0, 0, 0, 1, 0},
			},
			expectedFreeVarCount: 1,
			expectedFreeVars:     []abstractions.VariableNumber{3},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			matrix := abstractions.FromSlice(tc.matrixValues)
			hnf := NewHermiteNormalForm(matrix)

			freeVars := hnf.findFreeVariables()

			if len(freeVars) != tc.expectedFreeVarCount {
				t.Errorf("Expected %d free variables, got %d", tc.expectedFreeVarCount, len(freeVars))
			}

			for i, expected := range tc.expectedFreeVars {
				if i >= len(freeVars) {
					t.Errorf("Missing free variable at index %d, expected %d", i, expected)
					continue
				}
				if freeVars[i] != expected {
					t.Errorf("Free variable %d: expected %d, got %d", i, expected, freeVars[i])
				}
			}
		})
	}
}

func TestHermiteNormalForm_getUniqueSolution(t *testing.T) {
	tests := map[string]struct {
		matrixValues     [][]int64
		expectedSolution []int64
	}{
		"diagonal_matrix": {
			matrixValues: [][]int64{
				{1, 0, 0, 5},
				{0, 1, 0, 10},
				{0, 0, 1, 15},
			},
			expectedSolution: []int64{5, 10, 15},
		},
		"upper_triangular": {
			matrixValues: [][]int64{
				{1, 2, 3, 14},
				{0, 1, 2, 8},
				{0, 0, 1, 3},
			},
			expectedSolution: []int64{1, 2, 3},
		},
		"with_negative_values": {
			matrixValues: [][]int64{
				{1, 0, -1, -2, -3},
				{0, 1, 1, 5, 7},
				{0, 0, 3, 0, 0},
				{0, 0, 0, 13, 13},
			},
			expectedSolution: []int64{-1, 2, 0, 1},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			matrix := abstractions.FromSlice(tc.matrixValues)
			hnf := NewHermiteNormalForm(matrix)

			solution := hnf.getUniqueSolution()

			if solution == nil {
				t.Fatal("Expected solution, got nil")
			}

			if len(solution.Get()) != len(tc.expectedSolution) {
				t.Errorf("Expected %d variables, got %d", len(tc.expectedSolution), len(solution.Get()))
			}

			for i, expected := range tc.expectedSolution {
				variableNum := abstractions.VariableNumber(i + 1)
				actual := solution.GetValue(variableNum)
				if actual != expected {
					t.Errorf("Variable x%d: expected %d, got %d", variableNum, expected, actual)
				}
			}

			// Verify the solution satisfies all equations
			verifySolution(t, matrix, solution)
		})
	}
}

func TestHermiteNormalForm_Get(t *testing.T) {
	matrixValues := [][]int64{
		{1, 0, 0, 3},
		{0, 1, 0, 5},
		{0, 0, 1, 7},
	}

	matrix := abstractions.FromSlice(matrixValues)
	hnf := NewHermiteNormalForm(matrix)

	result := hnf.Get()

	if result != matrix {
		t.Error("Get() should return the same matrix reference")
	}
}

// Helper function to verify that a solution satisfies all equations in the matrix
func verifySolution(t *testing.T, matrix *abstractions.Matrix, solution *abstractions.Variables) {
	t.Helper()

	numVars := matrix.Cols() - 1

	for row := 0; row < matrix.Rows(); row++ {
		sum := int64(0)
		for col := 0; col < numVars; col++ {
			coefficient := matrix.Get(row, col)
			varNum := abstractions.VariableNumber(col + 1)
			varValue := solution.GetValue(varNum)
			sum += coefficient * varValue
		}

		expected := matrix.Get(row, numVars)

		// Skip zero rows (they don't constrain the solution)
		allZeros := true
		for col := 0; col <= numVars; col++ {
			if matrix.Get(row, col) != 0 {
				allZeros = false
				break
			}
		}

		if !allZeros && sum != expected {
			t.Errorf("Row %d: equation not satisfied. Expected %d, got %d", row, expected, sum)
			t.Logf("Solution: %v", solution.Get())
		}
	}
}
