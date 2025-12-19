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

	freeVariables := make([]VariableNumber, 0)
	foundVariables := make([]VariableNumber, 0)

	/* Note: The last column is the constants, so we skip it */
	expectedVariablesCount := r.matrix.Cols() - 1

	/* Finds out all the variables from the matrix that have a leading 1 */
	for row := 0; row < r.matrix.Rows(); row++ {
		/* The column where the leading 1 is found is the variable number */
		pivotCol := findPivotCol(r.matrix, row)

		if pivotCol != NotFound {
			foundVariables = append(foundVariables, VariableNumber(pivotCol+1))
		}
	}

	if expectedVariablesCount != len(foundVariables) {
		for variable := 1; variable <= expectedVariablesCount; variable++ {
			if !ContainsNumber(foundVariables, variable) {
				freeVariables = append(freeVariables, VariableNumber(variable))
			}
		}
	}

	return freeVariables
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

			total := float64(0)
			solvedVariables := NewVariables(variablesCount)

			for _, freeVariable := range freeVariables.Get() {
				solvedVariables.SetVariable(freeVariable)
			}

			/* Solve the equations row by row with the free variables values
			starting from the bottom of the matrix and upwards */
			for row := r.matrix.Rows() - 1; row >= 0; row-- {

				pivotCol := findPivotCol(r.matrix, row)

				variableNumber := VariableNumber(pivotCol + 1)

				if freeVariables.Contains(variableNumber) {
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
					dependantVariableNumber := VariableNumber(columnIndex + 1)

					if solvedVariables.Contains(dependantVariableNumber) {
						/* The current column / variable is a free variable in which case we use the value provided */
						dependantVariableValue = solvedVariables.GetValue(dependantVariableNumber)
					} else {
						/* The current column / variable is NOT a free variable in which case from the matrix itself
						on the last column */
						// NOT GOOD
						fmt.Printf("Not good.")
						//dependantVariableValue = r.matrix.Get(row, knownVariablesCols)
					}

					currentVariableValue -= operationSign * dependantVariableValue
				}

				solvedVariables.Set(variableNumber, currentVariableValue)
				total += currentVariableValue
			}

			if total >= lowestTotal {
				return
			}

			for _, solvedVariable := range solvedVariables.Get() {
				if solvedVariable.Value < 0 {
					return
				}
			}

			if verbose {
				fmt.Printf("Solved the equation with %d free variables:\n", freeVariables.Count())

				for _, variable := range solvedVariables.Get() {
					fmt.Printf("\tVariable %d = %.2f\n", variable.Number, variable.Value)
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
