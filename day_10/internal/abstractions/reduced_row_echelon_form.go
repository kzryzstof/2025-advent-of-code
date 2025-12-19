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

func (r *ReducedRowEchelonForm) findMinimalSolution(
	freeVariableIndices []int,
	verbose bool,
) []float64 {

	maxValue := 5
	lowestTotal := 9999

	r.testCombination(
		len(freeVariableIndices),
		maxValue, /* Here we choose test free variables values from 0 to 5 to find the minimal solution */
		func(freeVariables []int) {

			total := 0
			knownVariablesCols := r.matrix.Cols() - 1
			solvedVariablesValues := make([]float64, r.matrix.Rows())

			/* Solve the equations row by row with the free variables values
			starting from the bottom of the matrix and upwards */
			for row := r.matrix.Rows() - 1; row >= 0; row-- {

				/* We start from the constant and then apply the operations */
				currentVariableConstant := r.matrix.Get(row, r.matrix.Cols()-1)
				currentVariableValue := currentVariableConstant

				for columnIndex := 0; columnIndex < r.matrix.Cols()-1; columnIndex++ {

					operationSign := r.matrix.Get(row, columnIndex)
					dependantVariableValue := 0.0

					if Contains(freeVariables, columnIndex) {
						/* The current column / variable is a free variable in which case we use the value provided */
						index := IndexOf(freeVariables, columnIndex)
						dependantVariableValue = float64(freeVariables[index])
					} else {
						/* The current column / variable is NOT a free variable in which case from the matrix itself
						on the last column */
						dependantVariableValue = r.matrix.Get(row, knownVariablesCols)
					}

					if dependantVariableValue == 0 {
						continue
					}

					solvedVariablesValues[r.matrix.Rows()-1-row] = dependantVariableValue
					currentVariableValue += operationSign * dependantVariableValue
				}

				total += int(currentVariableValue)
			}

			if total >= lowestTotal {
				return
			}

			if verbose {
				fmt.Printf("Solved the equation with %d free variables:\n", len(freeVariables))

				for index, variableValue := range freeVariables {
					fmt.Printf("\tVariable %d = %d\n", index, variableValue)
				}

				fmt.Printf("\tResult is = %d\n", total)
				fmt.Print("\n\tThis combination has the minimal values so far!\n\n\n")

			}

			lowestTotal = total
		},
	)

	return make([]float64, 1)
}

func (r *ReducedRowEchelonForm) testCombination(
	/* Indicates the maximum number of combination to test (#) -> (#,#) -> (#,#,#) -> ... */
	maximumCombinationLength int,
	/* Indicates the biggest number to test (0->5) -> (0-5,0-5) -> (0-5,0-5,0-5) -> ... */
	maxNumber int,
	/* Function to call to test the combination */
	testCombination func([]int),
) {

	var testGroups func(currentButtonGroups []int, currentNumberToTest int)

	testGroups = func(currentButtons []int, currentNumberToTest int) {

		currentCombinationLength := len(currentButtons)

		canTest := currentCombinationLength == currentNumberToTest

		for currentNumber := 0; currentNumber < maxNumber; currentNumber++ {

			/* Test all the combinations with the current list of buttons */
			currentButtons[len(currentButtons)-1] = currentNumber

			if canTest {
				testCombination(currentButtons)
			}
		}

		if currentCombinationLength < maximumCombinationLength {

			/* Creates a new list of buttons for the next iteration that will have one more button added */
			numbersPrefix := make([]int, currentCombinationLength+1)
			Clear(numbersPrefix)

			/* Makes sure to include the current list */
			copy(numbersPrefix, currentButtons)

			/* Loops to test all the combinations with one more button group */
			for number := 0; number < maxNumber; number++ {

				/* Exclude the group index if it is already in the combination */
				if Contains(numbersPrefix, number) {
					continue
				}

				numbersPrefix[currentCombinationLength-1] = number

				testGroups(numbersPrefix, currentNumberToTest)
			}
		}
	}

	count := 1

	for count <= maximumCombinationLength {

		initialButtonGroups := make([]int, 1)
		Clear(initialButtonGroups)

		testGroups(initialButtonGroups, count)
		count++
	}
}
