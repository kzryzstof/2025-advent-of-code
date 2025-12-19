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
	//lowestTotal := float64(9999)

	r.testCombination(
		len(freeVariableIndices),
		maxValue, /* Here we choose test free variables values from 0 to 5 to find the minimal solution */
		func(freeVariables []int) {

			if verbose {
				//fmt.Printf("Received %d free variables:\n", len(freeVariables))

				for _, freeVariable := range freeVariables {
					fmt.Printf("%2d ", freeVariable)
				}

				fmt.Println()
			}

			//total := float64(0)
			//knownVariablesCols := r.matrix.Cols() - 1
			//solvedVariablesValues := make([]float64, r.matrix.Rows())
			//
			//missingVariablesCount := r.matrix.Cols() - 1
			//
			//for freeVariableIndex := missingVariablesCount - 1; freeVariableIndex >= r.matrix.Rows()-1; freeVariableIndex-- {
			//	index := IndexOf(freeVariableIndices, freeVariableIndex)
			//	solvedVariablesValues[index] = float64(freeVariables[index])
			//}
			//
			///* Solve the equations row by row with the free variables values
			//starting from the bottom of the matrix and upwards */
			//for row := r.matrix.Rows() - 1; row >= 0; row-- {
			//
			//	if Contains(freeVariableIndices, row) {
			//		/* The current column / variable is a free variable in which case we use the value provided */
			//		index := IndexOf(freeVariableIndices, row)
			//		solvedVariablesValues[row] = float64(freeVariableIndices[index])
			//		continue
			//	}
			//
			//	/* We start from the constant and then apply the operations */
			//	currentVariableConstant := r.matrix.Get(row, r.matrix.Cols()-1)
			//	currentVariableValue := currentVariableConstant
			//
			//	for columnIndex := 0; columnIndex < r.matrix.Cols()-1; columnIndex++ {
			//
			//		operationSign := r.matrix.Get(row, columnIndex)
			//
			//		if operationSign == 0 {
			//			continue
			//		}
			//
			//		dependantVariableValue := 0.0
			//
			//		if Contains(freeVariables, columnIndex) {
			//			/* The current column / variable is a free variable in which case we use the value provided */
			//			index := IndexOf(freeVariables, columnIndex)
			//			dependantVariableValue = float64(freeVariables[index])
			//		} else {
			//			/* The current column / variable is NOT a free variable in which case from the matrix itself
			//			on the last column */
			//			// NOT GOOD
			//			dependantVariableValue = r.matrix.Get(row, knownVariablesCols)
			//		}
			//
			//		solvedVariablesValues[r.matrix.Rows()-1-row] = dependantVariableValue
			//		currentVariableValue += operationSign * dependantVariableValue
			//	}
			//
			//	total += currentVariableValue
			//}
			//
			//if total >= lowestTotal {
			//	return
			//}
			//
			//if verbose {
			//	fmt.Printf("Solved the equation with %d free variables:\n", len(freeVariables))
			//
			//	for index, variableValue := range freeVariables {
			//		fmt.Printf("\tVariable %d = %d\n", index, variableValue)
			//	}
			//
			//	fmt.Printf("\tResult is = %f\n", total)
			//	fmt.Print("\n\tThis combination has the minimal values so far!\n\n\n")
			//
			//}
			//
			//lowestTotal = total
		},
	)

	return make([]float64, 1)
}

func (r *ReducedRowEchelonForm) testCombination(
	/* Indicates the maximum number of combination to test (#,#,#) */
	combinationLength int,
	/* Indicates the biggest number to test (0->5,0,0) -> (0,0->5,0) -> (0,0,0->5) -> ... */
	maxNumber int,
	/* Function to call to test the combination */
	testCombinationFunc func([]int),
) {

	var generateCombinationFunc func(currentButtonGroups []int, currentIndex int)

	generateCombinationFunc = func(numbersCombination []int, currentIndex int) {

		canTestCombination := currentIndex == combinationLength-1

		if canTestCombination {
			for currentNumber := 0; currentNumber < maxNumber; currentNumber++ {

				numbersCombination[currentIndex] = currentNumber

				testCombinationFunc(numbersCombination)
			}
		}

		if currentIndex < combinationLength {

			/* Creates a new list of buttons for the next iteration that will have one more button added */
			numbersPrefix := make([]int, len(numbersCombination))
			Clear(numbersPrefix)

			/* Makes sure to include the current list */
			copy(numbersPrefix, numbersCombination)

			/* Loops to test all the combinations with one more button group */
			for number := 0; number < maxNumber; number++ {

				numbersPrefix[currentIndex] = number

				generateCombinationFunc(numbersPrefix, currentIndex+1)
			}
		}
	}

	initialButtonGroups := make([]int, combinationLength)
	Clear(initialButtonGroups)
	generateCombinationFunc(initialButtonGroups, 0)
}
