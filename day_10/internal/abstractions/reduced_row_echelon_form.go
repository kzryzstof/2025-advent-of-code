package abstractions

import "fmt"

type ReducedRowEchelonForm struct {
	matrix *Matrix
}

func NewReducedRowEchelonForm(
	matrix *Matrix,
) *ReducedRowEchelonForm {
	return &ReducedRowEchelonForm{matrix}
}

func (r *ReducedRowEchelonForm) Get() *Matrix {
	return r.matrix
}

func (r *ReducedRowEchelonForm) Solve(
	verbose bool,
) []float64 {

	freeVariablesIndices := r.detectFreeVariables()

	if len(freeVariablesIndices) == 0 {

		if verbose {
			fmt.Print("\nNo free variables detected; unique solution exists\n\n")
		}

		return r.getUniqueSolution()
	}

	fmt.Printf("\n%d free variable(s) have been found\n\n", len(freeVariablesIndices))

	results := make([]float64, r.matrix.Rows())

	return results
}

func (r *ReducedRowEchelonForm) getUniqueSolution() []float64 {
	results := make([]float64, r.matrix.Rows())

	for row := 0; row < r.matrix.Rows(); row++ {
		results[row] = r.matrix.Get(row, r.matrix.Cols()-1)
	}

	return results
}

func (r *ReducedRowEchelonForm) detectFreeVariables() []int {
	freeVariableIndices := make([]int, 0)

	/* Note: The last column is the constants, so we skip it */

	for col := 0; col < r.matrix.Cols()-1; col++ {
		if r.isFreeVariable(col) {
			freeVariableIndices = append(freeVariableIndices, col)
		}
	}

	return freeVariableIndices
}

func (r *ReducedRowEchelonForm) isFreeVariable(
	variableRow int,
) bool {

	// A variable is free if there's no pivot (leading 1) in its column

	if variableRow >= r.matrix.Rows() {
		return true // More variables than equations
	}

	// Check if there's a pivot (non-zero, typically 1) at the diagonal
	return r.matrix.Get(variableRow, variableRow) == 0
}
