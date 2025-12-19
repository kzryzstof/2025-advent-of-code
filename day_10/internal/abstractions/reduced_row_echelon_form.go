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

	r.findMinimalSolution(freeVariablesIndices)

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
) []float64 {

	maxValue := 5

	presses, _ := r.testCombination(
		len(freeVariableIndices),
		maxValue, /* Here we choose test free variables values from 0 to 5 to find the minimal solution */
		func(freeVariables []int) bool {

			//fmt.Printf("Processing machine %d with %d button groups [ ", machineIndex+1, machine.GetButtonGroupsCount())

			for index, variableValue := range freeVariables {
				fmt.Printf("%d=%d ", index, variableValue)
			}

			fmt.Println()

			/* Tests if the machine is activated */
			return false
		},
	)

	return make([]float64, presses)
}

func (r *ReducedRowEchelonForm) testCombination(
	/* Indicates the maximum number of combination to test (#) -> (#,#) -> (#,#,#) -> ... */
	maximumCombinationLength int,
	/* Indicates the biggest number to test (0->5) -> (0-5,0-5) -> (0-5,0-5,0-5) -> ... */
	maxNumber int,
	/* Function to call to test the combination */
	testCombination func([]int) bool,
) (int, bool) {

	var testGroups func(currentButtonGroups []int, currentNumberToTest int) (int, bool)

	testGroups = func(currentButtons []int, currentNumberToTest int) (int, bool) {

		currentCombinationLength := len(currentButtons)

		canTest := currentCombinationLength == currentNumberToTest

		for currentNumber := 0; currentNumber < maxNumber; currentNumber++ {

			/* Test all the combinations with the current list of buttons */
			currentButtons[len(currentButtons)-1] = currentNumber

			if canTest && testCombination(currentButtons) {
				return currentCombinationLength, true
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

				pressedCount, succeeded := testGroups(numbersPrefix, currentNumberToTest)

				if succeeded {
					return pressedCount, true
				}
			}
		}

		return -1, false
	}

	count := 1

	for count <= maximumCombinationLength {

		initialButtonGroups := make([]int, 1)
		Clear(initialButtonGroups)

		pressesCount, succeeded := testGroups(initialButtonGroups, count)

		if succeeded {
			return pressesCount, true
		}

		count++
	}

	return NotFound, false
}
