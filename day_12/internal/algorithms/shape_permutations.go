package algorithms

import (
	"day_12/internal/abstractions"
)

type operation func([][]byte)

var (
	packDown  = abstractions.Direction{Row: 1, Col: 0}
	packUp    = abstractions.Direction{Row: -1, Col: 0}
	packRight = abstractions.Direction{Row: 0, Col: -1}
	packLeft  = abstractions.Direction{Row: 0, Col: 1}
)

const (
	// MaximumShapeSize /* All the presents occupies a 3x3 region */
	MaximumShapeSize = 3
)

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

					/* Test packing the shape from the right */

					for rowIndex := 0; rowIndex < 3; rowIndex++ {

						pack(
							leftShape,
							rightShape,
							abstractions.Position{Row: rowIndex, Col: 2},
							abstractions.Position{Row: rowIndex, Col: 0},
							packRight,
							abstractions.Direction{Row: 0, Col: 1},
						)
					}

				}
			}
		}
	}
}

func pack(
	shape [][]byte,
	otherShape [][]byte,
	position abstractions.Position,
	otherPosition abstractions.Position,
	packDirection abstractions.Direction,
	testDirection abstractions.Direction,
) abstractions.Dimension {

	leftObjectFreePositions := make([]abstractions.Position, 0, MaximumShapeSize)
	rightObjectFreePositions := make([]abstractions.Position, 0, MaximumShapeSize)

	oppositeDirection := packDirection.Opposite()

	/* The approach here, row after row, we measure how far left the shape can go */
	for sectionIndex := 0; sectionIndex < MaximumShapeSize; sectionIndex++ {

		/* Find out how much space in the opposite direction the shape can be packed */
		leftObjectFreePositions = append(
			leftObjectFreePositions,
			abstractions.FindEmptyIndex(
				shape,
				position,
				packDirection,
				/* Sets the boundary for testing cells */
				position.Offset(MaximumShapeSize, packDirection),
			),
		)

		rightObjectFreePositions = append(
			rightObjectFreePositions,
			abstractions.FindEmptyIndex(
				otherShape,
				otherPosition,
				oppositeDirection,
				/* Sets the boundary for testing cells */
				otherPosition.Offset(MaximumShapeSize, oppositeDirection),
			),
		)

		position = position.Add(testDirection)
		otherPosition = otherPosition.Add(testDirection)
	}

	return abstractions.Dimension{Wide: 0, Long: 0}
}
