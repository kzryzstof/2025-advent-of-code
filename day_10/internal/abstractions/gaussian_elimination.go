package abstractions

func Reduce(
	m *Matrix,
	v *Vector,
) {
	Print(m, v)

	forwardEliminate(m, v)
	backwardEliminate(m, v)
}

func validate(
	m *Matrix,
	v *Vector,
) {
	/* Makes sure all the pivots are at 1 */
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			cell := m.Get(row, col)
			if col == row {
				if cell != 1 {
					panic("pivot is not 1")
				}
			}
		}
	}

}

func forwardEliminate(
	m *Matrix,
	v *Vector,
) {
	pivot(m, v)
	normalize(m, v)
	eliminate(m, v)
}

/* Swaps rows to make the pivots [x,x] non-zero. */
func pivot(
	m *Matrix,
	v *Vector,
) {
	for pivot := 0; pivot < m.Rows()-1; pivot++ {
		if m.Get(pivot, pivot) == 0 {
			/* If the pivot is zero, look for a row *below* to swap with */
			row := findSwappableRow(m, pivot)
			m.Swap(pivot, row)
			v.Swap(pivot, row)
		}
	}

	Print(m, v)
}

/* Normalizes each row so that the pivots become 1 */
func normalize(
	m *Matrix,
	v *Vector,
) {
	for pivot := 0; pivot < m.Rows(); pivot++ {
		pivotValue := m.Get(pivot, pivot)

		if pivotValue == 0 || pivotValue == 1 {
			/* The scaling is necessary only if the pivot is not 0 or 1 */
			continue
		}

		m.Scale(pivot, 1/pivotValue)
		v.Scale(pivot, 1/pivotValue)
	}

	Print(m, v)
}

/* Eliminates (zeroes) entries below each pivot */
func eliminate(
	m *Matrix,
	v *Vector,
) {
	for pivot := 0; pivot < m.Rows(); pivot++ {

		/* Finds the pivot on this row */
		pivotCol := findPivotOnRow(m, pivot)

		if pivotCol == -1 {
			continue
		}

		for row := pivot + 1; row < m.Rows(); row++ {

			factor := m.Get(row, pivotCol)

			/* Skips rows where the factor is already 0
			because this is what we are looking for */
			if factor == 0 {
				continue
			}

			/* We can start the inner column loop at pivot instead of 0
			since entries left of the pivot are already zero by construction. */
			for col := pivotCol; col < m.Cols(); col++ {
				m.Set(row, col, m.Get(row, col)-factor*m.Get(pivot, col))
			}
			v.Set(row, v.Get(row)-factor*v.Get(pivot))
		}
	}

	Print(m, v)
}

func backwardEliminate(
	m *Matrix,
	v *Vector,
) {
	for pivot := m.Rows() - 1; pivot >= 0; pivot-- {

		pivotCol := findPivotOnRow(m, pivot)

		if pivotCol == -1 {
			continue
		}

		for row := pivot - 1; row >= 0; row-- {

			factor := m.Get(row, pivot)

			if factor == 0 {
				continue
			}

			for col := pivot; col < m.Cols(); col++ {
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
