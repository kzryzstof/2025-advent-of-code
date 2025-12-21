package algorithms

import (
	"day_11/internal/abstractions"
	"fmt"
	"math"
)

const (
	NotFound                = -1
	AuthorizedScalingFactor = -1
)

func ToHermiteNormalForm(
	augmentedMatrix *abstractions.AugmentedMatrix,
	verbose bool,
) *HermiteNormalForm {

	matrixCopy := abstractions.CopyMatrix(augmentedMatrix.Matrix)

	if verbose {
		fmt.Print("\n** Row reduction **\n\n")
		abstractions.Print(matrixCopy)
	}

	/* Step 1: Forward elimination */
	doForwardElimination(matrixCopy, verbose)

	if verbose {
		fmt.Print("***********************************************************************************************\n\n")
	}

	/* Step 2: Backward elimination */
	doBackwardElimination(matrixCopy, verbose)

	if verbose {
		fmt.Printf("Hermite Nominal Form\n")
		abstractions.Print(matrixCopy)
	}

	/* Step 3: Optimizes row order */

	/*
		In certain cases, a pivot may not be on the expected row,
		which makes solving the equation a bit harder later on.
		So we try to force the pivots to be on the diagonal
		by moving rows around.
	*/
	pivotRows(matrixCopy)

	if verbose {
		fmt.Printf("After final pivoting\n")
		abstractions.Print(matrixCopy)
	}

	return NewHermiteNormalForm(matrixCopy)
}

func doBackwardElimination(
	m *abstractions.Matrix,
	verbose bool,
) {

	if verbose {
		fmt.Print("**** Step 2: backward elimination **\n\n")
	}

	fromRow := m.Rows() - 1

	for pivotRow := fromRow; pivotRow > 0; pivotRow-- {

		/*
			When doing a backward elimination, we need to find the actual pivot column,
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
			abstractions.Print(m)
		}

		pivotValue := m.Get(pivotRow, pivotCol)

		/* We need the pivot to be positive, so we scale the row if it's negative */
		if pivotValue < 0 {
			scaleRow(m, AuthorizedScalingFactor, pivotRow, verbose)
			pivotValue = -pivotValue
		}

		for row := pivotRow - 1; row >= 0; row-- {

			cellRowValue := m.Get(row, pivotCol)

			q := computeQuotient(cellRowValue, pivotValue)

			if q != 0 {
				eliminateRow(m, row, pivotRow, pivotCol, q)
			}
		}

		if verbose {
			fmt.Print("\n\n")
			fmt.Printf("Backward elimination done on row %d\n\n", pivotRow+1)
			abstractions.Print(m)
		}
	}
}

func doForwardElimination(
	m *abstractions.Matrix,
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
			abstractions.Print(m)
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
			scaleRow(m, AuthorizedScalingFactor, pivot, verbose)
		}

		/* ********** Forward eliminate ********** */

		/* Now that we have a pivot of one, let's eliminate all the cells for the current column, row after row */
		eliminateRowForward(m, pivot)

		if verbose {
			fmt.Printf("Forward elimination done on row %d\n\n", pivot+1)
			abstractions.Print(m)
		}
	}
}

func eliminateRowForward(
	m *abstractions.Matrix,
	pivotRow int,
) {

	for row := pivotRow + 1; row < m.Rows(); row++ {

		/* Only work on the pivot column, not all columns! */
		pivotCol := pivotRow

		for m.Get(row, pivotCol) != 0 {

			pivotCellValue := m.Get(pivotRow, pivotCol)
			rowCellValue := m.Get(row, pivotCol)

			/*
				Euclidean Step
				--------------
				We want to reduce m[row][col] using m[pivotRow][col]
				This loop runs until m[row][col] becomes 0 (GCD logic)
			*/

			q := computeQuotient(rowCellValue, pivotCellValue)

			eliminateRow(m, row, pivotRow, pivotCol, q)

			/* Essential HNF step: Swap to continue GCD reduction */
			if m.Get(row, pivotCol) != 0 {
				m.Swap(pivotRow, row)
			}
		}
	}
}

func eliminateRow(
	m *abstractions.Matrix,
	row int,
	pivotRow int,
	pivotCol int,
	quotient int64,
) {
	for col := pivotCol; col < m.Cols(); col++ {
		/*
			Row operation: Row[i] = Row[i] - q * Row[pivotRow]

			Compared to the Gaussian elimination, this approach
			preserves integers
		*/
		m.Set(row, col, m.Get(row, col)-quotient*m.Get(pivotRow, col))
	}
}

func computeQuotient(
	rowCellValue int64,
	pivotCellValue int64,
) int64 {

	quotient := rowCellValue / pivotCellValue

	remainder := rowCellValue % pivotCellValue

	if remainder < 0 {
		quotient--
	}

	return quotient
}

func scaleRow(
	m *abstractions.Matrix,
	pivotValue int64,
	pivot int,
	verbose bool,
) {

	scaling := 1 / pivotValue

	m.Scale(pivot, scaling)

	if verbose {
		fmt.Printf("Scaling done on row %d (scaling: %d)\n\n", pivot+1, scaling)
		abstractions.Print(m)
	}
}

func pivotRow(
	m *abstractions.Matrix,
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
		abstractions.Print(m)
	}
}

func findPivotCol(
	m *abstractions.Matrix,
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
	m *abstractions.Matrix,
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

func pivotRows(
	m *abstractions.Matrix,
) {
	for row := 0; row < m.Rows(); row++ {

		pivotCol := findPivotCol(m, row)

		if pivotCol == NotFound {
			continue
		}

		if pivotCol != row {
			endRow := int(math.Min(float64(m.Rows()-1), float64(pivotCol-1)))
			m.MoveRowToEnd(row, endRow)
		}
	}
}
