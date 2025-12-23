package abstractions

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
