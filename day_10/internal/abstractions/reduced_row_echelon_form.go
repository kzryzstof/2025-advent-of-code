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

func (r *ReducedRowEchelonForm) Solve() []float64 {
	//	--- WE NEED TO DETERMINE IF THERE ARE FREE VARIABLES ---
	//	THEN RUN ASSIGN MULTIPLE SOLUTIONS IF NEEDED THAT GET THE LOWEST NUMBERS
	freeVariablesIndices := detectFreeVariables(r.matrix)
	fmt.Printf("Free variables: %v\n", freeVariablesIndices)

	results := make([]float64, r.matrix.Rows())

	for row := 0; row < r.matrix.Rows(); row++ {
		results[row] = r.matrix.Get(row, r.matrix.Cols()-1)
	}

	return results
}
