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

	r.findMinimalSolution(freeVariablesIndices, verbose)

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

func (r *ReducedRowEchelonForm) detectFreeVariables() []VariableNumber {
	freeVariableIndices := make([]VariableNumber, 0)

	/* Note: The last column is the constants, so we skip it */
	expectedVariablesCount := r.matrix.Cols() - 1

	for variableNumber := 0; variableNumber < expectedVariablesCount; variableNumber++ {
		if r.isFreeVariable(variableNumber) {
			freeVariableIndices = append(freeVariableIndices, VariableNumber(variableNumber+1))
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

func (r *ReducedRowEchelonForm) findMinimalSolution(
	freeVariableIndices []VariableNumber,
	verbose bool,
) []Variable {

	variablesCount := uint(r.matrix.Cols() - 1)
	maxVariableValue := float64(5) /* Here we choose test free variables values from 0 to 5 to find the minimal solution */
	lowestTotal := float64(9999)

	r.testCombination(
		freeVariableIndices,
		maxVariableValue,
		func(freeVariables *Variables) {

			//total := float64(0)
			solvedVariables := NewVariables(variablesCount)

			for _, freeVariable := range freeVariables.Get() {
				solvedVariables.SetVariable(freeVariable)
			}

			/* Solve the equations row by row with the free variables values
			starting from the bottom of the matrix and upwards */
			for row := r.matrix.Rows() - 1; row >= 0; row-- {

				if freeVariables.Contains(row + 1) {
					continue
				}

				/* We start from the constant and then apply the operations */
				currentVariableConstant := r.matrix.Get(row, r.matrix.Cols()-1)
				currentVariableValue := currentVariableConstant

				for columnIndex := 0; columnIndex < r.matrix.Cols()-1; columnIndex++ {

					operationSign := r.matrix.Get(row, columnIndex)

					if operationSign == 0 {
						continue
					}

					dependantVariableValue := 0.0

					if Contains(freeVariables, columnIndex) {
						/* The current column / variable is a free variable in which case we use the value provided */
						index := IndexOf(freeVariables, columnIndex)
						dependantVariableValue = float64(freeVariables[index])
					} else {
						/* The current column / variable is NOT a free variable in which case from the matrix itself
						on the last column */
						// NOT GOOD
						//dependantVariableValue = r.matrix.Get(row, knownVariablesCols)
					}

					solvedVariablesValues[r.matrix.Rows()-1-row] = dependantVariableValue
					currentVariableValue += operationSign * dependantVariableValue
				}

				total += currentVariableValue
			}

			if total >= lowestTotal {
				return
			}

			if verbose {
				fmt.Printf("Solved the equation with %d free variables:\n", len(freeVariables))

				for index, variableValue := range freeVariables {
					fmt.Printf("\tVariable %d = %d\n", index, variableValue)
				}

				fmt.Printf("\tResult is = %f\n", total)
				fmt.Print("\n\tThis combination has the minimal values so far!\n\n\n")

			}

			lowestTotal = total
		},
	)

	return nil
}

func (r *ReducedRowEchelonForm) testCombination(
	/* Indicates the number of combination to test (#,#,#) */
	variables []VariableNumber,
	/* Indicates the biggest number to test (0->5,0,0) -> (0,0->5,0) -> (0,0,0->5) -> ... */
	maxVariableValue float64,
	/* Function to call to test the combination */
	testCombinationFunc func(*Variables),
) {

	var generateCombinationFunc func(variables *Variables, currentVariableNumber VariableNumber)

	generateCombinationFunc = func(variables *Variables, currentVariableNumber VariableNumber) {

		canTestVariables := uint(currentVariableNumber) == variables.Count()

		if canTestVariables {
			for currentVariableValue := float64(0); currentVariableValue < maxVariableValue; currentVariableValue++ {
				variables.Set(currentVariableNumber, currentVariableValue)
				testCombinationFunc(variables)
			}
		}

		if !variables.IsLast(currentVariableNumber) {
			for currentVariableValue := float64(0); currentVariableValue < maxVariableValue; currentVariableValue++ {
				variables.Set(currentVariableNumber, currentVariableValue)
				generateCombinationFunc(variables, currentVariableNumber+1)
			}
		}
	}

	initialVariables := FromVariableNumbers(variables)
	generateCombinationFunc(initialVariables, 1)
}
