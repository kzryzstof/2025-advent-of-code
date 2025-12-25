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

	/*
		The approach here is that, row after row, we measure how far left the shape can go
		until it overlaps with the other shape
	*/

	stableObjectFreePositions := make([]abstractions.Position, 0)
	packedObjectFreePositions := make([]abstractions.Position, 0)

	oppositeDirection := packDirection.Opposite()

	stableShapeBoundary := abstractions.OriginPosition()
	packedShapeBoundary := otherPosition.Offset(MaximumShapeSize, oppositeDirection)

	for sectionIndex := 0; sectionIndex < MaximumShapeSize; sectionIndex++ {

		/* Find out how much space in the opposite direction the shape can be packed */
		stableObjectFreePositions = append(
			stableObjectFreePositions,
			abstractions.FindEmptyIndex(
				shape,
				position,
				packDirection,
				/* Sets the boundary for testing cells */
				stableShapeBoundary,
			),
		)

		packedObjectFreePositions = append(
			packedObjectFreePositions,
			abstractions.FindEmptyIndex(
				otherShape,
				otherPosition,
				oppositeDirection,
				/* Sets the boundary for testing cells */
				packedShapeBoundary,
			),
		)

		position = position.Add(testDirection)
		otherPosition = otherPosition.Add(testDirection)
	}

	/*
		The overlap has been detected. Let's get the boundaries of the new shape */

	/* Note: the other shape has an offset compared to the shape */
	minColLeft := abstractions.FindMinCol(stableObjectFreePositions)
	maxColRight := abstractions.FindMaxCol(packedObjectFreePositions) + packDirection.Col*MaximumShapeSize

	minRowRight := abstractions.FindMinRow(stableObjectFreePositions)
	maxRowRight := abstractions.FindMaxRow(packedObjectFreePositions) + packDirection.Row*MaximumShapeSize

	wide := maxColRight - minColLeft + 1
	long := maxRowRight - minRowRight + 1

	return abstractions.Dimension{Wide: wide, Long: long}
}
