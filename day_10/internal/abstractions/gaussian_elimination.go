package abstractions

import "fmt"

const (
	PivotValue = 1
	NotFound   = -1
)

func ToReducedRowEchelonForm(
	augmentedMatrix *AugmentedMatrix,
	verbose bool,
) *ReducedRowEchelonForm {

	rref := CopyMatrix(augmentedMatrix.Matrix)

	if verbose {
		fmt.Print("\n** Gaussian elimination **\n\n")
		Print(rref)
	}

	doForwardElimination(rref, verbose)

	if verbose {
		fmt.Print("***********************************************************************************************\n\n")
	}

	doBackwardElimination(rref, verbose)

	if verbose {
		fmt.Printf("Reduced Row Echelon Form\n")
		Print(rref)
	}

	return NewReducedRowEchelonForm(rref)
}

func doBackwardElimination(
	m *Matrix,
	verbose bool,
) {

	if verbose {
		fmt.Print("**** Step 2: backward elimination **\n\n")
	}

	fromRow := m.Rows() - 1

	for currentRow := fromRow; currentRow > 0; currentRow-- {

		/* When doing a backward elimination, we need to find the actual pivot column,
		since some rows may be all zeroes at this point.
		By definition, the pivot is the first non-zero entry in the row
		*/
		pivotCol := findPivotCol(m, currentRow)

		if pivotCol == NotFound {
			continue
		}

		if verbose {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Printf("Working on row %d. Pivot found on column %d\n", currentRow+1, pivotCol+1)
			Print(m)
		}

		for row := currentRow - 1; row >= 0; row-- {

			factor := m.Get(row, pivotCol)

			if verbose {
				fmt.Printf("E%d - %.0f E%d | ", row+1, factor, row)
			}

			/* Skips rows where the factor is already 0
			because this is what we are looking for */
			if factor == 0 {
				continue
			}

			/* We can start the inner column loop at pivot instead of 0
			since entries left of the pivot are already zero by construction. */
			for col := pivotCol; col < m.Cols(); col++ {
				m.Set(row, col, m.Get(row, col)-factor*m.Get(currentRow, col))
			}
		}

		if verbose {
			fmt.Print("\n\n")
			fmt.Printf("Backward elimination done on row %d\n\n", currentRow+1)
			Print(m)
		}
	}
}

func doForwardElimination(
	m *Matrix,
	verbose bool,
) {

	if verbose {
		fmt.Print("**** Step 1: forward elimination **\n\n")
	}

	for pivot := 0; pivot < m.Rows(); pivot++ {

		if pivot >= m.Cols() {
			break
		}

		if verbose {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Printf("Working on row %d\n\n", pivot+1)
			Print(m)
		}

		/* ********** Pivot row ********** */

		if m.Get(pivot, pivot) == 0 {

			/* If the pivot is 0, we try to find is a row *below* with a cell that is not 0 so that a swap can be done */
			pivotRow(m, pivot, verbose)
		}

		/* ********** Normalize ********** */

		pivotValue := m.Get(pivot, pivot)

		if pivotValue != 0 && pivotValue != PivotValue {

			/* The scaling is necessary only if the pivot is not 0 or 1 */
			scaleRow(m, pivotValue, pivot, verbose)
		}

		/* ********** Forward eliminate ********** */

		/* Now that we have a pivot of one, let's eliminate all the cells for the current column, row after row */
		eliminateRow(m, pivot, verbose)

		if verbose {
			fmt.Printf("Forward elimination done on row %d\n\n", pivot+1)
			Print(m)
		}
	}
}

func eliminateRow(
	m *Matrix,
	pivot int,
	verbose bool,
) {

	for row := pivot + 1; row < m.Rows(); row++ {

		factor := m.Get(row, pivot)

		if verbose {
			fmt.Printf("E%d - %.0f E%d | ", row+1, factor, row)
		}

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
	}

	if verbose {
		fmt.Print("\n\n")
	}
}

func scaleRow(
	m *Matrix,
	pivotValue float64,
	pivot int,
	verbose bool,
) {

	scaling := 1 / pivotValue

	m.Scale(pivot, scaling)

	if verbose {
		fmt.Printf("Scaling done on row %d (scaling: %.2f)\n\n", pivot+1, scaling)
		Print(m)
	}
}

func pivotRow(
	m *Matrix,
	pivot int,
	verbose bool,
) {
	pivotRow := findSwappableRow(m, pivot)

	if pivotRow == NotFound {
		return
	}

	m.Swap(pivot, pivotRow)

	if verbose {
		fmt.Printf("Pivoting row %d with %d\n", pivot+1, pivotRow+1)
		Print(m)
	}
}

func findPivotCol(
	m *Matrix,
	row int,
) int {

	for candidateColumn := 0; candidateColumn < m.Cols()-1; candidateColumn++ {
		if m.Get(row, candidateColumn) == PivotValue {
			return candidateColumn
		}
	}

	return NotFound
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

	return NotFound
}
