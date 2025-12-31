package maths

import (
	"math"
)

func NewSlice(
	rows, cols uint,
	defaultVal int8,
) [][]int8 {
	slice := make([][]int8, rows)

	for row := range slice {
		slice[row] = make([]int8, cols)
		for col := range slice[row] {
			slice[row][col] = defaultVal
		}
	}

	return slice
}

func Transpose(
	slice [][]int8,
) {
	/* We assume square data for simplicity */

	for row := 0; row < len(slice); row++ {
		for col := row + 1; col < len(slice); col++ {
			slice[row][col], slice[col][row] = slice[col][row], slice[row][col]
		}
	}
}

func HorizontalFlip(
	slice [][]int8,
) {
	for r := 0; r < 3; r++ {
		slice[r][0], slice[r][2] = slice[r][2], slice[r][0]
	}
}

func VerticalFlip(
	slice [][]int8,
) {
	for c := 0; c < 3; c++ {
		slice[0][c], slice[2][c] = slice[2][c], slice[0][c]
	}
}

func RotateClockwise(
	slice [][]int8,
) {
	Transpose(slice)
	HorizontalFlip(slice)
}

func NoOp(
	slice [][]int8,
) {
}

func CopySlice(
	slice [][]int8,
) [][]int8 {

	sliceCopy := make([][]int8, len(slice))

	for row := 0; row < len(sliceCopy); row++ {
		sliceCopy[row] = make([]int8, len(slice[row]))
		for col := 0; col < len(sliceCopy[row]); col++ {
			sliceCopy[row][col] = slice[row][col]
		}
	}

	return sliceCopy
}

func Slide(
	src [][]int8,
	direction Vector,
	defaultValue int8,
) [][]int8 {

	/*
		Defines the size of the destination slice, which can differ from the src slice
		since we must take into account the potential shift
	*/
	dstRows := len(src) + int(math.Abs(float64(direction.Row)))
	dstCols := len(src[0]) + int(math.Abs(float64(direction.Col)))

	dst := NewSlice(uint(dstRows), uint(dstCols), defaultValue)

	/* Copies the src slice into the dst slice at the specified shift */

	for row := 0; row < len(src); row++ {

		if row+direction.Row >= dstRows || row+direction.Row < 0 {
			continue
		}

		for col := 0; col < len(src[row]); col++ {

			if col+direction.Col >= dstCols || col+direction.Col < 0 {
				continue
			}

			dst[row+direction.Row][col+direction.Col] = src[row][col]
		}
	}

	return dst
}

func FindLastCellWithValueOnRow(
	slice [][]int8,
	position Position,
	direction Vector,
	value int8,
) (Position, bool) {

	maxRow := len(slice)
	emptyCellFound := false

	for slice[position.Row][position.Col] == value {

		emptyCellFound = true
		position = position.Add(direction)

		maxCol := len(slice[position.Row])

		if position.Col >= maxCol || position.Row >= maxRow {
			return position, true
		}

		if position.Col < 0 || position.Row < 0 {
			return Position{
				Row: position.Row,
				Col: -1,
			}, true
		}
	}

	return position, emptyCellFound
}

func PasteShape(
	id uint,
	src [][]int8,
	dst [][]int8,
	rowOffset, colOffset uint,
	ignoredValue int8,
) {
	for row := uint(0); row < uint(len(src)); row++ {
		for col := uint(0); col < uint(len(src[row])); col++ {

			if src[row][col] == ignoredValue {
				continue
			}

			rowWithOffset := row + rowOffset
			colWithOffset := col + colOffset

			dst[rowWithOffset][colWithOffset] = int8(id)
		}
	}
}
