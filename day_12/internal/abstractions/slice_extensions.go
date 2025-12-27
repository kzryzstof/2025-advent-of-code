package abstractions

import (
	"fmt"
	"math"
)

const (
	NotFound = -1
)

func Transpose(
	slice [][]byte,
) {
	/* We assume square data for simplicity */

	for row := 0; row < len(slice); row++ {
		for col := row + 1; col < len(slice); col++ {
			slice[row][col], slice[col][row] = slice[col][row], slice[row][col]
		}
	}
}

func HorizontalFlip(
	slice [][]byte,
) {
	for r := 0; r < 3; r++ {
		slice[r][0], slice[r][2] = slice[r][2], slice[r][0]
	}
}

func VerticalFlip(
	slice [][]byte,
) {
	for c := 0; c < 3; c++ {
		slice[0][c], slice[2][c] = slice[2][c], slice[0][c]
	}
}

func RotateClockwise(
	slice [][]byte,
) {
	Transpose(slice)
	HorizontalFlip(slice)
}

func NoOp(
	slice [][]byte,
) {
}

func GetCopy(
	slice [][]byte,
) [][]byte {

	sliceCopy := make([][]byte, len(slice))

	for row := 0; row < len(sliceCopy); row++ {
		sliceCopy[row] = make([]byte, len(slice[row]))
		for col := 0; col < len(sliceCopy[row]); col++ {
			sliceCopy[row][col] = slice[row][col]
		}
	}

	return sliceCopy
}

func Clear(
	slice [][]byte,
) {

	for row := 0; row < len(slice); row++ {
		for col := 0; col < len(slice[row]); col++ {
			slice[row][col] = 0
		}
	}
}

func CopyTo(
	src [][]byte,
	shift Direction,
) [][]byte {

	/*
		Defines the size of the destination slice, which can differ from the src slice
		since we must take into account the potential shift
	*/
	dstRows := len(src) + int(math.Abs(float64(shift.Row)))
	dstCols := len(src[0]) + int(math.Abs(float64(shift.Col)))

	dst := make([][]byte, dstRows)

	for row := 0; row < dstRows; row++ {
		dst[row] = make([]byte, dstCols)
	}

	/* Copies the src slice into the dst slice at the specified shift */

	for row := 0; row < len(src); row++ {

		if row+shift.Row >= dstRows || row+shift.Row < 0 {
			continue
		}

		for col := 0; col < len(src[row]); col++ {

			if col+shift.Col >= dstCols || col+shift.Col < 0 {
				continue
			}

			dst[row+shift.Row][col+shift.Col] = src[row][col]
		}
	}

	return dst
}

func IsEmpty(
	slice [][]byte,
	position Position,
) bool {
	return slice[position.Row][position.Col] == 0
}

func FindEmptyIndex(
	slice [][]byte,
	position Position,
	direction Direction,
) Position {

	for IsEmpty(slice, position) {
		position = position.Add(direction)
		if position.Col > 2 || position.Row > 2 || position.Col < 0 || position.Row < 0 {
			return position
		}
	}

	return position
}

func PrintShape(
	slice [][]byte,
) {
	for _, row := range slice {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func PrintShapes(
	leftSlice [][]byte,
	rightSlice [][]byte,
) {
	printRowCell := func(row []byte) {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
	}

	for rowIndex := range leftSlice {
		printRowCell(leftSlice[rowIndex])
		fmt.Print(" | ")
		printRowCell(rightSlice[rowIndex])
		fmt.Println()
	}
}
