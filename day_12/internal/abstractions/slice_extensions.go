package abstractions

import "fmt"

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
	dst [][]byte,
	src [][]byte,
	x, y int,
) {

	for row := 0; row < len(src); row++ {
		for col := 0; col < len(src[row]); col++ {
			dst[row+y][col+x] = src[row][col]
		}
	}
}

func Print(
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
		println()
	}
}
