package abstractions

import "fmt"

const (
	PivotValue = 1
)

func Reduce(
	augmentedMatrix *AugmentedMatrix,
	verbose bool,
) []float64 {
	m := augmentedMatrix.Matrix
	v := augmentedMatrix.Vector

	if verbose {
		fmt.Println("Forward elimination")
		Print(m, v)
	}

	/* Forward elimination */
	for pivot := 0; pivot < m.Rows()-1; pivot++ {

		if pivot >= m.Cols() {
			break
		}

		if verbose {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Printf("Working on row %d\n", pivot+1)
			Print(m, v)
		}

		if m.Get(pivot, pivot) == 0 {

			/* ********** Pivot ********** */

			/* If the pivot is not 1, let's see if there is a row *below* that we can use to swap it with */
			pivotRow := findSwappableRow(m, pivot)

			if pivotRow != -1 {

				m.Swap(pivot, pivotRow)
				v.Swap(pivot, pivotRow)

				if verbose {
					fmt.Printf("Pivoting row %d with %d\n", pivot+1, pivotRow+1)
					Print(m, v)
				}
			}
		}

		/* ********** Normalize ********** */

		pivotValue := m.Get(pivot, pivot)

		if pivotValue != 0 && pivotValue != PivotValue {
			/* The scaling is necessary only if the pivot is not 0 or 1 */
			scaling := 1 / pivotValue

			m.Scale(pivot, scaling)
			v.Scale(pivot, scaling)

			if verbose {
				fmt.Printf("Normalized on row %d (scaling: %f)\n", pivot+1, scaling)
				Print(m, v)
			}
		}

		/* ********** Forward eliminate ********** */

		/* Now that we have a pivot of one, let's eliminate all the cells for the current column */

		for row := pivot + 1; row < m.Rows(); row++ {

			factor := m.Get(row, pivot)

			/* Skips rows where the factor is already 0
			because this is what we are looking for */
			if factor == 0 {
				continue
			}

			/* We can start the inner column loop at pivot instead of 0
			since entries left of the pivot are already zero by construction. */
			for col := pivot; col < m.Cols(); col++ {
				m.Set(row, col, m.Get(row, col)-factor*m.Get(pivot, col))
			}
			v.Set(row, v.Get(row)-factor*v.Get(pivot))
		}

		if verbose {
			fmt.Printf("Forward elimination done on row %d\n", pivot+1)
			Print(m, v)
		}
	}

	if verbose {
		fmt.Printf("Final form\n")
		Print(m, v)
	}

	return nil
	/*
		return backSubstitution(
			m,
			v,
			verbose,
		)
	*/
}

func findSwappableRow(
	m *Matrix,
	pivotRow int,
) int {

	pivotCol := pivotRow

	for candidateRow := pivotRow + 1; candidateRow < m.Rows(); candidateRow++ {
		if m.Get(candidateRow, pivotCol) != 0 {
			return candidateRow
		}
	}

	return -1
}

func backSubstitution(
	m *Matrix,
	v *Vector,
	verbose bool,
) []float64 {

	solution := make([]float64, m.Cols())

	/* Keep track of the free variables */
	freeVariableIndices := detectFreeVariables(m)
	freeVariablesValues := make([]float64, len(freeVariableIndices))

	continueSubstitution := true
	currentFreeVariableIncrementIndex := 0
	currentFreeVariableBase := float64(0)

	/* All the free variables are set to zero. If one of the none free variables is negative? */

	for continueSubstitution {

		variablesCount := m.Rows()
		variableRow := variablesCount - 1

		for ; variableRow >= 0; variableRow-- {

			if isFreeVariable(m, variableRow) {
				/* The free variable is already set in the solution */
				continue
			}

			/* Because of the way the matrix has been constructed, if a variable has dependency on other variables,
			we know their value (assuming the equation is solvable).
			*/

			total := v.Get(variableRow)

			for otherVariableColumn := variableRow + 1; otherVariableColumn < variablesCount; otherVariableColumn++ {

				otherVariableDependencySign := m.Get(variableRow, otherVariableColumn)

				if otherVariableDependencySign == 0 {
					/* Not a dependency */
					continue
				}

				total -= solution[otherVariableColumn] * otherVariableDependencySign
			}

			solution[variableRow] = total
		}

		continueSubstitution = false

		if verbose {
			fmt.Println("Current solution attempt:")
			PrintSlice(solution)
		}

		/* Validate if there is a negative variable in the solution */
		for variableIndex := 0; variableIndex < len(solution); variableIndex++ {
			if solution[variableIndex] < 0 {

				/* We have found a negative value: need to increment one of the free variables and try again */
				continueSubstitution = true

				/* Clears the solution for the next attempt */
				for variableIndex := 0; variableIndex < len(solution); variableIndex++ {
					solution[variableIndex] = 0
				}

				/* Increments the free variable in the solution */
				for index, freeVariableIndex := range freeVariableIndices {
					if index == currentFreeVariableIncrementIndex {
						solution[freeVariableIndex] = currentFreeVariableBase + 1
					}
				}

				currentFreeVariableIncrementIndex++
				if currentFreeVariableIncrementIndex >= len(freeVariablesValues) {
					currentFreeVariableBase++
					currentFreeVariableIncrementIndex = 0

					if currentFreeVariableBase == 2 {
						panic("The system of equations has no solution with non-negative variables")
					}
				}
				break
			}
		}
	}

	return solution
}

func detectFreeVariables(
	m *Matrix,
) []int {
	freeVariableIndices := make([]int, 0)

	for col := 0; col < m.Cols(); col++ {
		if isFreeVariable(m, col) {
			freeVariableIndices = append(freeVariableIndices, col)
		}
	}

	return freeVariableIndices
}

func isFreeVariable(
	m *Matrix,
	variableRow int,
) bool {

	// A variable is free if there's no pivot (leading 1) in its column
	// Check if row variableCol has a pivot at column variableCol
	if variableRow >= m.Rows() {
		return true // More variables than equations
	}

	// Check if there's a pivot (non-zero, typically 1) at the diagonal
	return m.Get(variableRow, variableRow) == 0
}
