package abstractions

import (
	"fmt"
	"math"
)

const (
	NotFound = -1
)

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

func GetCopy(
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

func SlideShape(
	src [][]int8,
	shift Vector,
) [][]int8 {

	/*
		Defines the size of the destination slice, which can differ from the src slice
		since we must take into account the potential shift
	*/
	dstRows := len(src) + int(math.Abs(float64(shift.Row)))
	dstCols := len(src[0]) + int(math.Abs(float64(shift.Col)))

	dst := make([][]int8, dstRows)

	for row := 0; row < dstRows; row++ {
		dst[row] = make([]int8, dstCols)
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
	slice [][]int8,
	position Position,
) bool {
	return slice[position.Row][position.Col] == 0
}

func FindLastEmptyCell(
	slice [][]int8,
	position Position,
	direction Vector,
) Position {

	for IsEmpty(slice, position) {
		position = position.Add(direction)
		if position.Col > 2 || position.Row > 2 || position.Col < 0 || position.Row < 0 {
			return position
		}
	}

	return position
}

func ComputeFillRatio(
	slice [][]int8,
) float64 {

	empty, occupied := 0, 0

	for row := 0; row < len(slice); row++ {
		for col := 0; col < len(slice[row]); col++ {
			if slice[row][col] == 0 {
				empty++
			} else {
				occupied++
			}
		}
	}

	return float64(occupied) / float64(occupied+empty)
}

func PrintShape(
	slice [][]int8,
) {
	for _, row := range slice {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(fmt.Sprintf("%d", cell))
			}
		}
		fmt.Println()
	}
}

func PrintShapes(
	leftSlice [][]int8,
	rightSlice [][]int8,
) {
	printRowCell := func(row []int8) {
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
