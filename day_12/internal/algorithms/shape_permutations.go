package algorithms

import (
	"day_12/internal/abstractions"
	"fmt"
)

type operation func([][]byte)

/*
Precomputes the combinations of presents together
The goal is to combine all the presents together, in
all angles and keep track of the width and length
of the new shapes.

This way, when trying to fit them in the christmas tree,
we can quickly check if there is enough space for the
combined shapes instead of trying all the combinations
of presents every time.
*/
func ComputePermutations(
	presents *abstractions.Presents,
) {
	/* List of operations to apply to the presents */
	operations := []operation{
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.VerticalFlip,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.HorizontalFlip,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
		abstractions.RotateClockwise,
	}

	canvas := createCanvas()

	for leftPresent := range presents.GetAllPresents() {

		leftShape := leftPresent.GetShape()

		for _, leftOperation := range operations {

			/* Apply the operation in-place on the left shape */
			leftOperation(leftShape)

			for rightPresent := range presents.GetAllPresents() {

				rightShape := rightPresent.GetShape()

				for _, rightOperation := range operations {

					/* Apply the operation in-place on the left shape */
					rightOperation(rightShape)

					abstractions.Clear(canvas)
					abstractions.CopyTo(canvas, leftShape, 0, 0)
				}
			}
		}
	}
}

func createCanvas() [][]byte {
	canvas := make([][]byte, 6)

	for i := 0; i < 6; i++ {
		canvas[i] = make([]byte, 6)
		for j := 0; j < 6; j++ {
			canvas[i][j] = byte(0)
		}
	}
}

func packFromRight(
	leftShape [][]byte,
	rightShape [][]byte,
) [][]byte {

}
