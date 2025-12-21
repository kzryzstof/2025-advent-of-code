package abstractions

import "fmt"

const (
	PivotValue              = 1
	NotFound                = -1
	AuthorizedScalingFactor = -1
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
		fmt.Printf("Hermite Nominal Form\n")
		Print(rref)
	}

	for row := 0; row < rref.Rows(); row++ {
		pivotCol := findPivotCol(rref, row)

		if pivotCol == NotFound {
			continue
		}

		if pivotCol != row {
			if verbose {
				fmt.Printf("Moving row from %d to row %d\n", row+1, pivotCol+1)
			}
			for startRow := row; startRow < pivotCol-1; startRow++ {
				rref.Swap(startRow, startRow+1)
			}
		}
	}

	if verbose {
		fmt.Printf("After pivoting\n")
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

	for pivotRow := fromRow; pivotRow > 0; pivotRow-- {

		/* When doing a backward elimination, we need to find the actual pivot column,
		since some rows may be all zeroes at this point.
		By definition, the pivot is the first non-zero entry in the row
		*/
		pivotCol := findPivotCol(m, pivotRow)

		if pivotCol == NotFound {
			continue
		}

		if verbose {
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Printf("Working on row %d. Pivot found on column %d\n", pivotRow+1, pivotCol+1)
			Print(m)
		}

		pivotValue := m.Get(pivotRow, pivotCol)

		/* We need the pivot to be positive, so we scale the row if it's negative */
		if pivotValue < 0 {
			scaleRow(m, AuthorizedScalingFactor, pivotRow, verbose)
			pivotValue = -pivotValue
		}

		for row := pivotRow - 1; row >= 0; row-- {

			cellRowValue := m.Get(row, pivotCol)

			q := cellRowValue / pivotValue

			remainder := cellRowValue % pivotValue

			if remainder < 0 {
				q--
			}

			if q != 0 {
				for col := pivotCol; col < m.Cols(); col++ {
					m.Set(row, col, m.Get(row, col)-q*m.Get(pivotRow, col))
				}
			}
		}

		if verbose {
			fmt.Print("\n\n")
			fmt.Printf("Backward elimination done on row %d\n\n", pivotRow+1)
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

		/* Contrary to Gaussian, HNF only allows scaling with a value of -1 */
		if pivotValue < 0 {

			/* The scaling is necessary only if the pivot is negative */
			const AuthorizedScalingFactor = -1
			scaleRow(m, AuthorizedScalingFactor, pivot, verbose)
		}

		/* ********** Forward eliminate ********** */

		/* Now that we have a pivot of one, let's eliminate all the cells for the current column, row after row */
		eliminateRowForward(m, pivot, verbose)

		if verbose {
			fmt.Printf("Forward elimination done on row %d\n\n", pivot+1)
			Print(m)
		}
	}
}

func eliminateRowForward(
	m *Matrix,
	pivotRow int,
	verbose bool,
) {

	for row := pivotRow + 1; row < m.Rows(); row++ {

		// Only work on the pivot column, not all columns
		pivotCol := pivotRow

		for m.Get(row, pivotCol) != 0 {

			pivotCellValue := m.Get(pivotRow, pivotCol)
			rowCellValue := m.Get(row, pivotCol)

			if pivotCellValue == 0 {
				m.Swap(pivotRow, row)
				continue
			}

			// Euclidean Step:
			// We want to reduce m[i][col] using m[pivotRow][col]
			// This loop runs until m[i][col] becomes 0 (GCD logic)

			q := rowCellValue / pivotCellValue

			remainder := rowCellValue % pivotCellValue

			if remainder < 0 {
				q--
			}

			for col := pivotCol; col < m.Cols(); col++ {
				// Row operation: Row[i] = Row[i] - q * Row[pivotRow]
				m.Set(row, col, m.Get(row, col)-q*m.Get(pivotRow, col))
			}

			// Essential HNF step: Swap to continue GCD reduction
			if m.Get(row, pivotCol) != 0 {
				m.Swap(pivotRow, row)
			}
		}
	}
}

func eliminateRowAbove(
	m *Matrix,
	pivotRow int,
	pivotValue int64,
	verbose bool,
) {
	pivotCol := findPivotCol(m, pivotRow)

	for row := pivotRow - 1; row >= 0; row-- {

		cellRowValue := m.Get(row, pivotCol)

		q := cellRowValue / pivotValue

		remainder := cellRowValue % pivotValue

		if remainder < 0 {
			q--
		}

		if q != 0 {
			for col := pivotCol; col < m.Cols(); col++ {
				m.Set(row, col, m.Get(row, col)-q*m.Get(pivotRow, col))
			}
		}
	}
}

func scaleRow(
	m *Matrix,
	pivotValue int64,
	pivot int,
	verbose bool,
) {

	scaling := 1 / pivotValue

	m.Scale(pivot, scaling)

	if verbose {
		fmt.Printf("Scaling done on row %d (scaling: %d)\n\n", pivot+1, scaling)
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
		fmt.Printf("Pivoting row %d with %d\n\n", pivot+1, pivotRow+1)
		Print(m)
	}
}

func findPivotCol(
	m *Matrix,
	row int,
) int {

	for candidateColumn := 0; candidateColumn < m.Cols()-1; candidateColumn++ {
		if m.Get(row, candidateColumn) != 0 {
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
