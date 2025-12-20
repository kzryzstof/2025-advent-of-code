package abstractions

import (
	"fmt"
	"os"
)

const (
	MinNumbers = 0
	MaxNumbers = 250
)

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
) *Variables {

	variablesNumbers := r.detectFreeVariables()

	if len(variablesNumbers) == 0 {

		if verbose {
			fmt.Print("\nNo free variables detected; unique solution exists\n\n")
		}

		return r.getUniqueSolution()
	}

	if verbose {
		fmt.Printf("\n%d free variable(s) have been found: ", len(variablesNumbers))

		for _, variableNumber := range variablesNumbers {
			fmt.Printf("%d, ", variableNumber)
		}

		fmt.Print("\n\n")
	}

	/* To increase the speed of the process, we use different max numbers */
	var solution *Variables = nil
	maxVariableValues := []float64{10, 250, 500}
	maxVariableValueIndex := 0

	for solution == nil {

		solution = r.findMinimalSolution(
			variablesNumbers,
			maxVariableValues[maxVariableValueIndex],
			verbose,
		)

		if solution == nil {
			maxVariableValueIndex++
			if maxVariableValueIndex >= len(maxVariableValues) {
				fmt.Print("\nNo solution found!\n")
				os.Exit(1)
			}
		}
	}

	return solution
}

func (r *ReducedRowEchelonForm) getUniqueSolution() *Variables {

	variablesCount := uint(r.matrix.Cols() - 1)
	variables := NewVariables(variablesCount)

	for row := uint(0); row < variablesCount; row++ {
		variables.SetVariable(&Variable{
			VariableNumber(row + 1),
			r.matrix.Get(int(row), r.matrix.Cols()-1),
		})
	}

	return variables
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
	freeVariableNumbers []VariableNumber,
	maxVariableValue float64,
	verbose bool,
) *Variables {

	variablesCount := uint(r.matrix.Cols() - 1)
	minVariableValue := float64(MinNumbers)
	lowestTotal := float64(9999)
	var solution *Variables

	r.testCombination(
		freeVariableNumbers,
		minVariableValue,
		maxVariableValue,
		func(freeVariables *Variables) {

			total := float64(0)
			solvedVariables := NewVariables(variablesCount)

			for _, freeVariable := range freeVariables.Get() {
				solvedVariables.SetVariable(freeVariable)
				total += freeVariable.Value
			}

			allVariablesAssigned := false

			for !allVariablesAssigned {

				/* Solve the equations row by row with the free variables values
				starting from the bottom of the matrix and upwards */
				for row := r.matrix.Rows() - 1; row >= 0; row-- {

					pivotCol := findPivotCol(r.matrix, row)

					if pivotCol == NotFound {
						continue
					}

					currentVariableNumber := VariableNumber(pivotCol + 1)

					if solvedVariables.Contains(currentVariableNumber) {
						/* Variable is already assigned. Move on to the next one */
						continue
					}

					/* We start from the constant and then apply the operations */
					currentVariableConstant := r.matrix.Get(row, r.matrix.Cols()-1)
					currentVariableValue := currentVariableConstant
					currentVariableSign := r.matrix.Get(row, pivotCol)
					isVariableComputed := true

					for columnIndex := pivotCol + 1; columnIndex < r.matrix.Cols()-1; columnIndex++ {

						operationSign := r.matrix.Get(row, columnIndex)

						if operationSign == 0 {
							continue
						}

						dependantVariableValue := 0.0
						dependantVariableNumber := VariableNumber(columnIndex + 1)

						if solvedVariables.Contains(dependantVariableNumber) {
							/* The current column / variable has been assigned a value */
							dependantVariableValue = solvedVariables.GetValue(dependantVariableNumber)
						} else {
							/* The current column / variable is NOT available yet
							on the last column */
							isVariableComputed = false
							break
						}

						currentVariableValue -= operationSign * dependantVariableValue
					}

					if isVariableComputed {
						solvedVariables.SetVariable(&Variable{
							Number: currentVariableNumber,
							Value:  currentVariableValue,
						})
						total += currentVariableSign * currentVariableValue
					}
				}

				allVariablesAssigned = solvedVariables.Count() == variablesCount
			}

			if lowestTotal < 0 || total >= lowestTotal {
				return
			}

			for _, solvedVariable := range solvedVariables.Get() {
				if solvedVariable.Value < -0.01 {
					return
				}
			}

			if verbose {
				fmt.Printf("Solved the equation with %d free variables:\n", freeVariables.Count())

				fmt.Printf("[ ")
				for _, variable := range solvedVariables.Get() {
					fmt.Printf("%d=%.2f ", variable.Number, variable.Value)
				}

				fmt.Printf("] = %.2f\n", total)
			}

			lowestTotal = total
			solution = CopyVariables(solvedVariables)
		},
	)

	return solution
}

func (r *ReducedRowEchelonForm) testCombination(
	/* Indicates the number of combinations to test (#,#,#) */
	variableNumbers []VariableNumber,
	/* Indicates the biggest number to test (0->5,0,0) -> (0,0->5,0) -> (0,0,0->5) -> ... */
	minVariableValue float64,
	maxVariableValue float64,
	/* Function to call to test the combination */
	testCombinationFunc func(*Variables),
) {

	var generateCombinationFunc func(variables *Variables, currentVariableIndex uint)

	generateCombinationFunc = func(variables *Variables, currentVariableIndex uint) {

		isLastVariable := currentVariableIndex+1 == variables.Count()
		currentVariableNumber := variables.GetNumberByIndex(currentVariableIndex)

		if isLastVariable {
			for currentVariableValue := minVariableValue; currentVariableValue <= maxVariableValue; currentVariableValue++ {
				variables.Set(currentVariableNumber, currentVariableValue)
				testCombinationFunc(variables)
			}
		}

		if !isLastVariable {
			for currentVariableValue := float64(0); currentVariableValue < maxVariableValue; currentVariableValue++ {
				variables.Set(currentVariableNumber, currentVariableValue)
				generateCombinationFunc(variables, currentVariableIndex+1)
			}
		}
	}

	initialVariables := FromVariableNumbers(variableNumbers, minVariableValue)
	generateCombinationFunc(initialVariables, 0)
}
