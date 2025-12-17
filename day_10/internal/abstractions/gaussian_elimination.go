package abstractions

import "fmt"

func Reduce(
	m *Matrix,
	v *Vector,
) {
	Print(m, v)

	for pivot := 0; pivot < m.Rows()-1; pivot++ {

		pivotRow := pivot

		if m.Get(pivotRow, pivot) == 0 {

			/* ***** */
			/* Pivot */

			/* If the pivot is zero, look for a row *below* to swap with */
			pivotRow := findSwappableRow(m, pivot)
			m.Swap(pivot, pivotRow)
			v.Swap(pivot, pivotRow)
		} else {
			pivotRow++
		}

		/* ********* */
		/* Normalize */

		pivotValue := m.Get(pivotRow, pivot)

		if pivotValue == 0 || pivotValue == 1 {
			/* The scaling is necessary only if the pivot is not 0 or 1 */
			continue
		}

		m.Scale(pivotRow, 1/pivotValue)
		v.Scale(pivotRow, 1/pivotValue)

		/* ***************** */
		/* Forward eliminate */

		for row := pivotRow + 1; row < m.Rows(); row++ {

			factor := m.Get(row, pivot)

			/* Skips rows where the factor is already 0
			because this is what we are looking for */
			if factor == 0 {
				continue
			}

			/* We can start the inner column loop at pivot instead of 0
			since entries left of the pivot are already zero by construction. */
			for col := pivot; col < m.Cols(); col++ {
				m.Set(row, col, m.Get(row, col)-factor*m.Get(pivotRow, col))
			}
			v.Set(row, v.Get(row)-factor*v.Get(pivotRow))
		}
	}

	backwardEliminate(m, v)

	Print(m, v)
}

func backwardEliminate(
	m *Matrix,
	v *Vector,
) {
	fmt.Println("Backward eliminate")

	for pivot := m.Rows() - 1; pivot >= 0; pivot-- {

		pivotCol := findPivotOnRow(m, pivot)

		if pivotCol == -1 {
			continue
		}

		for row := pivot - 1; row >= 0; row-- {

			factor := m.Get(row, pivotCol)

			if factor == 0 {
				continue
			}

			for col := pivotCol; col < m.Cols(); col++ {
				m.Set(row, col, m.Get(row, col)-factor*m.Get(pivot, col))
			}

			v.Set(row, v.Get(row)-factor*v.Get(pivot))
		}
	}

	Print(m, v)
}

func findSwappableRow(
	m *Matrix,
	pivotRow int,
) int {

	for candidateRow := pivotRow + 1; candidateRow < m.Rows(); candidateRow++ {
		if m.Get(candidateRow, pivotRow) != 0 {
			return candidateRow
		}
	}

	panic("no swappable row found. OMG! I am panicking too!!!")
}

func findPivotOnRow(
	m *Matrix,
	pivot int,
) int {

	pivotCol := -1

	for col := 0; col < m.Cols(); col++ {
		if m.Get(pivot, col) != 0 {
			pivotCol = col
			break
		}
	}

	return pivotCol
}
