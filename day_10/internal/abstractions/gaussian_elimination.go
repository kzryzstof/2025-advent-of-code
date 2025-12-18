package abstractions

import "fmt"

const (
	PivotValue = 1
)

func Reduce(
	augmentedMatrix *AugmentedMatrix,
) []float64 {
	m := augmentedMatrix.Matrix
	v := augmentedMatrix.Vector

	fmt.Println("Augmented matrix reduction started")
	Print(m, v)

	for pivot := 0; pivot < m.Rows(); pivot++ {

		if pivot >= m.Cols() {
			break
		}

		fmt.Println("-----------------------------------------------------------------------------------------------")
		fmt.Printf("Working on row %d\n", pivot+1)
		Print(m, v)

		if m.Get(pivot, pivot) == 0 {

			/* ********** Pivot ********** */

			/* If the pivot is not 1, let's see if there is a row *below* that we can use to swap it with */
			pivotRow := findSwappableRow(m, pivot)

			if pivotRow != -1 {

				m.Swap(pivot, pivotRow)
				v.Swap(pivot, pivotRow)

				fmt.Printf("Pivoting row %d with %d\n", pivot+1, pivotRow+1)
				Print(m, v)
			}
		}

		/* ********** Normalize ********** */

		pivotValue := m.Get(pivot, pivot)

		if pivotValue != 0 && pivotValue != PivotValue {
			/* The scaling is necessary only if the pivot is not 0 or 1 */
			scaling := 1 / pivotValue

			m.Scale(pivot, scaling)
			v.Scale(pivot, scaling)

			fmt.Printf("Normalized on row %d (scaling: %f)\n", pivot+1, scaling)
			Print(m, v)
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

		fmt.Printf("Forward elimination done on row %d\n", pivot+1)
		Print(m, v)
	}

	fmt.Printf("Final form\n")
	Print(m, v)

	return backSubstitution(
		m,
		v,
	)
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
) []float64 {
	solution := make([]float64, m.Cols())
	defaultFreeVariableValue := float64(1)

	variablesCount := m.Cols()
	variableRow := variablesCount - 1

	for ; variableRow >= m.Rows(); variableRow-- {
		solution[variableRow] = 1
		defaultFreeVariableValue = 0
	}

	for ; variableRow >= 0; variableRow-- {
		if isFreeVariable(m, variableRow) {
			//	TODO Check if this one could be zero instead !
			//	If yes, it also means other could be negative...
			solution[variableRow] = defaultFreeVariableValue
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

	return solution
}

func isFreeVariable(
	m *Matrix,
	variableRow int,
) bool {
	for col := 0; col < m.Cols(); col++ {
		if m.Get(variableRow, col) != 0 {
			return false
		}
	}

	return true
}
