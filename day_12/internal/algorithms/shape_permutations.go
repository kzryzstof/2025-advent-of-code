package algorithms

import (
	"day_12/internal/abstractions"
	"fmt"
)

type operation func([][]byte)

var (
	packDown    = abstractions.Direction{Row: 1, Col: 0}
	packUp      = abstractions.Direction{Row: -1, Col: 0}
	packToLeft  = abstractions.Direction{Row: 0, Col: -1}
	packToRight = abstractions.Direction{Row: 0, Col: 1}
)

const (
	// MaximumShapeSize /* All the presents occupies a 3x3 region */
	MaximumShapeSize = 3
	CanvasSize       = MaximumShapeSize * 2
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
	verbose bool,
) *abstractions.CombinationCatalog {
	combinationsCount := 0

	/* List of operations to apply to the presents */
	operations := []operation{
		abstractions.NoOp,
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

	catalog := abstractions.NewCombinationCatalog()

	for leftPresent := range presents.GetAllPresents() {

		if verbose {
			fmt.Println("*************************************************")
			fmt.Printf("Looking for optimal combination for present %d\n", leftPresent.GetIndex())
		}

		leftShape := leftPresent.GetShape()

		for _, leftOperation := range operations {

			/* Apply the operation in-place on the left shape */
			leftOperation(leftShape)

			for rightPresent := range presents.GetAllPresents() {

				rightShape := rightPresent.GetShape()

				for operationIndex, rightOperation := range operations {

					if verbose {
						fmt.Printf("Packing present %d with %d (%d/%d)\r", leftPresent.GetIndex(), rightPresent.GetIndex(), operationIndex+1, len(operations))
					}

					/* Apply the operation in-place on the left shape */
					rightOperation(rightShape)

					if verbose {
						fmt.Println("Left stable shape:")
						abstractions.Print(leftShape)
						fmt.Println("\nRight shape:")
						abstractions.Print(rightShape)
						fmt.Println()
					}

					/* Test packing the shape from the right */
					//for rowIndex := 0; rowIndex < 3; rowIndex++ {

					packedDimension := pack(
						leftShape,
						rightShape,
						packToLeft,
					)

					if verbose {
						fmt.Printf("Dimension of the combination: %dx%d\n", packedDimension.Wide, packedDimension.Long)
						fmt.Println()
					}

					combinationsCount++

					catalog.StoreNewCombination(
						leftPresent.GetIndex(),
						rightPresent.GetIndex(),
						packedDimension,
					)
				}
			}
		}
	}

	if verbose {
		fmt.Printf("All presents packed: %d combinations tested              \n", combinationsCount)
	}

	return catalog
}

func pack(
	stableShape [][]byte,
	packedShape [][]byte,
	packDirection abstractions.Direction,
) abstractions.Dimension {

	packOffset := computePackOffset(
		stableShape,
		packedShape,
		packDirection,
	)

	/*
		Temporary canvas large enough to hold both shapes plus offset.
		For 3x3 inputs, 6x6 is safe.
	*/

	canvas := make([][]byte, CanvasSize)

	for row := range canvas {
		canvas[row] = make([]byte, CanvasSize)
	}

	// Helper to place a shape at an offset.
	placeShape := func(shape [][]byte, rowOffset, colOffset int) {
		for row := 0; row < MaximumShapeSize; row++ {
			for col := 0; col < MaximumShapeSize; col++ {
				if shape[row][col] == 0 {
					continue
				}
				rowWithOffset := row + rowOffset
				colWithOffset := col + colOffset

				if rowWithOffset < 0 || rowWithOffset >= CanvasSize || colWithOffset < 0 || colWithOffset >= CanvasSize {
					continue
				}

				canvas[rowWithOffset][colWithOffset] = 1
			}
		}
	}

	// Stable at origin (0,0).
	placeShape(
		stableShape,
		0,
		0,
	)

	// Packed shape translated by packOffset.
	placeShape(
		packedShape,
		packDirection.Opposite().Row*MaximumShapeSize+packOffset.Row,
		packDirection.Opposite().Col*MaximumShapeSize+packOffset.Col,
	)

	/*
		.###
		####
		####
	*/
	// Compute bounding box of the combined shape.
	minRow, maxRow := CanvasSize, -1
	minCol, maxCol := CanvasSize, -1

	for r := 0; r < CanvasSize; r++ {
		for c := 0; c < CanvasSize; c++ {
			if canvas[r][c] == 0 {
				continue
			}
			if r < minRow {
				minRow = r
			}
			if r > maxRow {
				maxRow = r
			}
			if c < minCol {
				minCol = c
			}
			if c > maxCol {
				maxCol = c
			}
		}
	}

	if maxRow == -1 {
		// No cells set; return empty.
		return abstractions.Dimension{}
	}

	wide := maxCol - minCol + 1
	long := maxRow - minRow + 1

	return abstractions.Dimension{Wide: wide, Long: long}
}

func computePackOffset(
	stableShape [][]byte,
	packedShape [][]byte,
	packDirection abstractions.Direction,
) abstractions.Position {

	oppositeDirection := packDirection.Opposite()

	stableShapeBoundary := abstractions.OriginPosition()
	packedShapeBoundary := abstractions.OriginPosition().Offset(MaximumShapeSize, oppositeDirection)

	/* We compute the delta at each row, and the smallest delta tells us how far we can pack the shape */
	minimumDelta := abstractions.Position{Row: 3, Col: 3}
	minimumDistance := abstractions.OriginPosition().GetDistanceTo(minimumDelta)

	for row := 0; row < MaximumShapeSize; row++ {

		/* At this specific row, we figure out many empty spots they are looking from the right to the left */
		initialStablePosition := abstractions.Position{Row: row, Col: MaximumShapeSize - 1}

		stablePosition := abstractions.FindEmptyIndex(
			stableShape,
			initialStablePosition,
			packDirection,
			/* Sets the boundary for testing cells */
			stableShapeBoundary,
		)

		stableShapeDelta := initialStablePosition.SubPosition(stablePosition)

		/* Gets the empty spot available in the stable shape */
		initialPackedPosition := abstractions.Position{Row: row, Col: 0}

		packedPosition := abstractions.FindEmptyIndex(
			packedShape,
			initialPackedPosition,
			oppositeDirection,
			/* Sets the boundary for testing cells */
			packedShapeBoundary,
		)

		packedShapeDelta := initialPackedPosition.SubPosition(packedPosition)

		/* Compute the number of cells the packed shape can be moved to the left */
		deltaShape := stableShapeDelta.Mul(packDirection).AddPosition(packedShapeDelta)

		distance := abstractions.OriginPosition().GetDistanceTo(deltaShape)

		if distance < minimumDistance {
			minimumDistance = distance
			minimumDelta = deltaShape
		}
	}

	return abstractions.Position{Row: packDirection.Row * int(minimumDistance), Col: packDirection.Col * int(minimumDistance)}
}
