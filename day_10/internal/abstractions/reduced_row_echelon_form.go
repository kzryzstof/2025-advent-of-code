package abstractions

import (
	"fmt"
	"os"
)

const (
	MinNumbers = -250
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
	maxVariableValues := []int64{10, 250, 500}
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

	variablesCount := uint64(r.matrix.Cols() - 1)
	variables := NewVariables(variablesCount)

	for row := r.matrix.Rows() - 1; row >= 0; row-- {

		/* The column where the leading 1 is found is the variable number */
		pivotCol := findPivotCol(r.matrix, row)

		if pivotCol == NotFound {
			continue
		}

		pivot := r.matrix.Get(row, pivotCol)

		if pivot == 0 {
			continue
		}

		total := r.matrix.Get(row, int(variablesCount))
		left := int64(0)
		for col := row + 1; col < int(variablesCount); col++ {
			left += r.matrix.Get(row, col) * variables.GetValue(VariableNumber(col+1))
		}
		variables.SetVariable(&Variable{
			VariableNumber(row + 1),
			(total - left) / pivot,
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

			alreadyExists := false
			for _, foundVariable := range foundVariables {
				if foundVariable == VariableNumber(pivotCol+1) {
					alreadyExists = true
					break
				}
			}

			if !alreadyExists {
				foundVariables = append(foundVariables, VariableNumber(pivotCol+1))
			}
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
	maxVariableValue int64,
	verbose bool,
) *Variables {

	variablesCount := uint64(r.matrix.Cols() - 1)
	minVariableValue := int64(MinNumbers)
	lowestTotal := int64(9999)
	var solution *Variables

	r.testCombination(
		freeVariableNumbers,
		minVariableValue,
		maxVariableValue,
		func(freeVariables *Variables) {

			total := int64(0)
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
					rowConstant := r.matrix.Get(row, int(variablesCount))
					pivot := r.matrix.Get(row, pivotCol)
					left := int64(0)

					for columnIndex := pivotCol + 1; columnIndex < r.matrix.Cols()-1; columnIndex++ {
						left += r.matrix.Get(row, columnIndex) * solvedVariables.GetValue(VariableNumber(columnIndex+1))
					}

					remainder := (rowConstant - left) % pivot

					if remainder != 0 {
						/* This solution won't provide an integer */
						return
					}

					solvedVariableValue := (rowConstant - left) / (pivot)
					solvedVariable := &Variable{
						Number: currentVariableNumber,
						Value:  solvedVariableValue,
					}
					solvedVariables.SetVariable(solvedVariable)
					total += solvedVariable.Value
				}

				allVariablesAssigned = solvedVariables.Count() == variablesCount
			}

			if lowestTotal < 0 || total >= lowestTotal {
				return
			}

			for _, solvedVariable := range solvedVariables.Get() {
				if solvedVariable.Value < 0 {
					return
				}
			}

			if verbose {
				fmt.Printf("Solved the equation with %d free variables:\n", freeVariables.Count())

				fmt.Printf("[ ")
				for _, variable := range solvedVariables.Get() {
					fmt.Printf("%d=%d ", variable.Number, variable.Value)
				}

				fmt.Printf("] = %d\n", total)
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
	minVariableValue int64,
	maxVariableValue int64,
	/* Function to call to test the combination */
	testCombinationFunc func(*Variables),
) {

	var generateCombinationFunc func(variables *Variables, currentVariableIndex uint64)

	generateCombinationFunc = func(variables *Variables, currentVariableIndex uint64) {

		isLastVariable := currentVariableIndex+1 == variables.Count()
		currentVariableNumber := variables.GetNumberByIndex(currentVariableIndex)

		if isLastVariable {
			for currentVariableValue := minVariableValue; currentVariableValue <= maxVariableValue; currentVariableValue++ {
				variables.Set(currentVariableNumber, currentVariableValue)
				testCombinationFunc(variables)
			}
		}

		if !isLastVariable {
			for currentVariableValue := minVariableValue; currentVariableValue < maxVariableValue; currentVariableValue++ {
				variables.Set(currentVariableNumber, currentVariableValue)
				generateCombinationFunc(variables, currentVariableIndex+1)
			}
		}
	}

	initialVariables := FromVariableNumbers(variableNumbers, minVariableValue)
	generateCombinationFunc(initialVariables, 0)
}
