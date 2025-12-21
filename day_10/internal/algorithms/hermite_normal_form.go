package algorithms

import (
	"day_10/internal/abstractions"
	"fmt"
	"os"
)

const (
	MinNumbers = -50
	MaxNumbers = 50
)

type HermiteNormalForm struct {
	matrix *abstractions.Matrix
}

func NewHermiteNormalForm(
	matrix *abstractions.Matrix,
) *HermiteNormalForm {
	return &HermiteNormalForm{matrix}
}

func (r *HermiteNormalForm) Get() *abstractions.Matrix {
	return r.matrix
}

func (r *HermiteNormalForm) Solve(
	verbose bool,
) *abstractions.Variables {

	variablesNumbers := r.findFreeVariables()

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
	var solution *abstractions.Variables = nil
	maxVariableValues := []int64{25, 100, 250}
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

func (r *HermiteNormalForm) getUniqueSolution() *abstractions.Variables {

	variablesCount := uint64(r.matrix.Cols() - 1)
	variables := abstractions.NewVariables(variablesCount)

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
			left += r.matrix.Get(row, col) * variables.GetValue(abstractions.VariableNumber(col+1))
		}
		variables.SetVariable(&abstractions.Variable{
			abstractions.VariableNumber(row + 1),
			(total - left) / pivot,
		})
	}

	return variables
}

func (r *HermiteNormalForm) findFreeVariables() []abstractions.VariableNumber {

	freeVariables := make([]abstractions.VariableNumber, 0)
	foundVariables := make([]abstractions.VariableNumber, 0)

	/* Note: The last column is the constants, so we skip it */
	expectedVariablesCount := r.matrix.Cols() - 1

	/* Finds out all the variables from the matrix that have a leading 1 */
	for row := 0; row < r.matrix.Rows(); row++ {
		/* The column where the leading 1 is found is the variable number */
		pivotCol := findPivotCol(r.matrix, row)

		if pivotCol != NotFound {

			alreadyExists := false
			for _, foundVariable := range foundVariables {
				if foundVariable == abstractions.VariableNumber(pivotCol+1) {
					alreadyExists = true
					break
				}
			}

			if !alreadyExists {
				foundVariables = append(foundVariables, abstractions.VariableNumber(pivotCol+1))
			}
		}

	}

	if expectedVariablesCount != len(foundVariables) {
		for variable := 1; variable <= expectedVariablesCount; variable++ {
			if !abstractions.ContainsNumber(foundVariables, variable) {
				freeVariables = append(freeVariables, abstractions.VariableNumber(variable))
			}
		}
	}

	return freeVariables
}

func (r *HermiteNormalForm) findMinimalSolution(
	freeVariableNumbers []abstractions.VariableNumber,
	maxVariableValue int64,
	verbose bool,
) *abstractions.Variables {

	variablesCount := uint64(r.matrix.Cols() - 1)
	minVariableValue := int64(MinNumbers)
	lowestTotal := int64(9999)
	var solution *abstractions.Variables

	r.testCombination(
		freeVariableNumbers,
		minVariableValue,
		maxVariableValue,
		func(freeVariables *abstractions.Variables) {

			total := int64(0)
			solvedVariables := abstractions.NewVariables(variablesCount)

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

					currentVariableNumber := abstractions.VariableNumber(pivotCol + 1)

					if solvedVariables.Contains(currentVariableNumber) {
						/* Variable is already assigned. Move on to the next one */
						continue
					}

					/* We start from the constant and then apply the operations */
					rowConstant := r.matrix.Get(row, int(variablesCount))
					pivot := r.matrix.Get(row, pivotCol)
					left := int64(0)

					for columnIndex := pivotCol + 1; columnIndex < r.matrix.Cols()-1; columnIndex++ {
						left += r.matrix.Get(row, columnIndex) * solvedVariables.GetValue(abstractions.VariableNumber(columnIndex+1))
					}

					remainder := (rowConstant - left) % pivot

					if remainder != 0 {
						/* This solution won't provide an integer */
						return
					}

					solvedVariableValue := (rowConstant - left) / (pivot)
					solvedVariable := &abstractions.Variable{
						Number: currentVariableNumber,
						Value:  solvedVariableValue,
					}
					solvedVariables.SetVariable(solvedVariable)
					total += solvedVariable.Value
				}

				allVariablesAssigned = solvedVariables.Count() == variablesCount
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

				fmt.Printf("[ ")
				for _, variable := range solvedVariables.Get() {
					fmt.Printf("%d=%d ", variable.Number, variable.Value)
				}

				fmt.Printf("] = %d\n", total)
			}

			lowestTotal = total
			solution = abstractions.CopyVariables(solvedVariables)
		},
	)

	return solution
}

func (r *HermiteNormalForm) testCombination(
	/* Indicates the number of combinations to test (#,#,#) */
	variableNumbers []abstractions.VariableNumber,
	/* Indicates the biggest number to test (0->5,0,0) -> (0,0->5,0) -> (0,0,0->5) -> ... */
	minVariableValue int64,
	maxVariableValue int64,
	/* Function to call to test the combination */
	testCombinationFunc func(*abstractions.Variables),
) {

	var generateCombinationFunc func(variables *abstractions.Variables, currentVariableIndex uint64)

	generateCombinationFunc = func(variables *abstractions.Variables, currentVariableIndex uint64) {

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

	initialVariables := abstractions.FromVariableNumbers(variableNumbers, minVariableValue)
	generateCombinationFunc(initialVariables, 0)
}
