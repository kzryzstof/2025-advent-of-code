package abstractions

import "fmt"

type Matrix struct {
	values   [][]float64
	colCount int
	rowCount int
}

func CopyMatrix(
	m *Matrix,
) *Matrix {

	values := make([][]float64, m.Rows())

	for row := 0; row < m.Rows(); row++ {
		values[row] = make([]float64, m.Cols())

		for col := 0; col < m.Cols(); col++ {
			values[row][col] = m.Get(row, col)
		}
	}

	return &Matrix{
		values,
		m.Cols(),
		m.Rows(),
	}
}

func NewMatrix(
	rowCount int,
	colCount int,
) *Matrix {

	values := make([][]float64, rowCount)

	for i := 0; i < rowCount; i++ {
		values[i] = make([]float64, colCount)
	}

	return &Matrix{
		values,
		colCount,
		rowCount,
	}
}

func FromSlice(
	slice [][]float64,
) *Matrix {

	rowsCount := len(slice)
	colsCount := len(slice[0])

	values := make([][]float64, rowsCount)

	for i := 0; i < rowsCount; i++ {
		values[i] = make([]float64, colsCount)
		for j := 0; j < colsCount; j++ {
			values[i][j] = slice[i][j]
		}
	}

	return &Matrix{
		values,
		colsCount,
		rowsCount,
	}
}

func (m *Matrix) Rows() int {
	return int(m.rowCount)
}

func (m *Matrix) Cols() int {
	return int(m.colCount)
}

func (m *Matrix) Set(row, col int, value float64) {
	m.values[row][col] = value
}

func (m *Matrix) Get(row, col int) float64 {
	return m.values[row][col]
}

func (m *Matrix) Swap(fromRow, toRow int) {
	for col := 0; col < m.colCount; col++ {
		value := m.values[fromRow][col]
		m.values[fromRow][col] = m.values[toRow][col]
		m.values[toRow][col] = value
	}
}

func (m *Matrix) Scale(row int, factor float64) {
	for col := 0; col < m.Cols(); col++ {
		m.Set(row, col, m.Get(row, col)*factor)
	}
}

func (m *Matrix) Print() {
	for row := 0; row < m.Rows(); row++ {
		fmt.Printf("[ ")
		for col := 0; col < m.Cols(); col++ {
			fmt.Printf("%.2f ", m.Get(row, col))
		}
		fmt.Println("]")
	}
	fmt.Println()
}
