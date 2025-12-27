package abstractions

import (
	"fmt"
)

type operation func([][]byte)

var (
	packDown    = Direction{Row: 1, Col: 0}
	packUp      = Direction{Row: -1, Col: 0}
	packToLeft  = Direction{Row: 0, Col: -1}
	packToRight = Direction{Row: 0, Col: 1}
)

const (
	// MaximumShapeSize /* All the presents occupies a 3x3 region */
	MaximumShapeSize = 3
	CanvasSize       = MaximumShapeSize * 2
)

// ComputePermutations /* Precomputes the combinations of presents together
//
//	The goal is to combine all the presents together, in
//	all angles and keep track of the width and length
//	of the new shapes.
//
//	This way, when trying to fit them in the christmas tree,
//	we can quickly check if there is enough space for the
//	combined shapes instead of trying all the combinations
//	of presents every time.
func ComputePermutations(
	presents *Presents,
	region *Region,
	verbose bool,
) *CombinationCatalog {
	combinationsCount := 0

	/* List of operations to apply to the presents */
	operations := []operation{
		NoOp,
		RotateClockwise,
		RotateClockwise,
		RotateClockwise,
		RotateClockwise,
		VerticalFlip,
		RotateClockwise,
		RotateClockwise,
		RotateClockwise,
		RotateClockwise,
		HorizontalFlip,
		RotateClockwise,
		RotateClockwise,
		RotateClockwise,
		RotateClockwise,
	}

	catalog := NewCombinationCatalog()

	for leftPresent := range presents.GetAllPresents() {

		if verbose {
			fmt.Println("*************************************************")
			fmt.Printf("Looking for optimal combination for present %d\n", leftPresent.GetIndex())
		}

		leftShape := leftPresent.GetShape()

		for _, applyLeftOperation := range operations {

			/* Apply the operation in-place on the left shape */
			applyLeftOperation(leftShape)

			for rightPresent := range presents.GetAllPresents() {

				rightShape := rightPresent.GetShape()

				for operationIndex, applyRightOperation := range operations {

					if verbose {
						fmt.Printf("Packing present %d with %d (%d/%d)\r", leftPresent.GetIndex(), rightPresent.GetIndex(), operationIndex+1, len(operations))
					}

					/* Apply the operation in-place on the right shape */
					applyRightOperation(rightShape)

					/*
						Test packing the shape from the right while also altering the row & col
						to test additional combinations
					*/

					for directionShift := 0; directionShift < MaximumShapeSize; directionShift++ {

						/*
							NOTE
							The computed shift is based on the direction of the packing.
							If we are packing horizontally, we must shift the vertical position
						*/
						additionalShift := Direction{
							Row: directionShift * -packToLeft.Col,
							Col: directionShift * -packToLeft.Row,
						}

						shiftRightShape := CopyTo(
							rightShape,
							additionalShift,
						)

						if verbose {
							PrintShapes(leftShape, shiftRightShape)
							fmt.Println()
						}

						packedDimension := pack(
							leftShape,
							shiftRightShape,
							packToLeft,
							verbose,
						)

						if verbose {
							fmt.Printf("Dimension of the combination: %dx%d\n", packedDimension.Wide, packedDimension.Long)
							fmt.Println()
						}

						combinationsCount++

						catalog.StoreNewCombination(
							leftPresent.GetIndex(),
							rightPresent.GetIndex(),
							Dimension{
								Wide: packedDimension.Wide + additionalShift.Col,
								Long: packedDimension.Long + additionalShift.Row,
							},
							region,
						)
					}
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
	packDirection Direction,
	verbose bool,
) Dimension {

	packOffset := computePackOffset(
		stableShape,
		packedShape,
		packDirection,
	)

	if verbose {
		fmt.Printf("Offset detected is %dx%d\n", packOffset.Row, packOffset.Col)
	}

	/*
		Temporary canvas large enough to hold both shapes plus offset.
		For 3x3 inputs, 8x8 is safe, taking into account offsets
	*/

	canvas := make([][]byte, CanvasSize)

	for row := range canvas {
		canvas[row] = make([]byte, CanvasSize)
	}

	/* Helper to place a shape at an offset */
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

	/* Stable shape placed at origin (0,0) */
	placeShape(
		stableShape,
		0,
		0,
	)

	/* Packed shape translated by packOffset */
	placeShape(
		packedShape,
		packDirection.Opposite().Row*MaximumShapeSize+packOffset.Row,
		packDirection.Opposite().Col*MaximumShapeSize+packOffset.Col,
	)

	if verbose {
		fmt.Println()
		PrintShape(
			canvas,
		)
		fmt.Println()
	}

	/* Compute bounding box of the combined shape */
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
		return Dimension{}
	}

	wide := maxCol - minCol + 1
	long := maxRow - minRow + 1

	return Dimension{
		Wide: wide,
		Long: long,
	}
}

func computePackOffset(
	stableShape [][]byte,
	packedShape [][]byte,
	packDirection Direction,
) Position {

	oppositeDirection := packDirection.Opposite()

	/* We compute the delta at each row, and the smallest delta tells us how far we can pack the shape */
	minimumDelta := Position{Row: 3, Col: 3}
	minimumDistance := OriginPosition().GetDistanceTo(minimumDelta)

	for row := 0; row < MaximumShapeSize; row++ {

		/* At this specific row, we figure out many empty spots they are looking from the right to the left */
		initialStablePosition := Position{Row: row, Col: MaximumShapeSize - 1}

		stablePosition := FindEmptyIndex(
			stableShape,
			initialStablePosition,
			packDirection,
		)

		stableShapeDelta := initialStablePosition.SubPosition(stablePosition)

		/* Gets the empty spot available in the stable shape */
		initialPackedPosition := Position{Row: row, Col: 0}

		packedPosition := FindEmptyIndex(
			packedShape,
			initialPackedPosition,
			oppositeDirection,
		)

		packedShapeDelta := initialPackedPosition.SubPosition(packedPosition)

		/* Compute the number of cells the packed shape can be moved to the left */
		deltaShape := stableShapeDelta.Mul(packDirection).AddPosition(packedShapeDelta)

		distance := OriginPosition().GetDistanceTo(deltaShape)

		if distance < minimumDistance {
			minimumDistance = distance
			minimumDelta = deltaShape
		}
	}

	return Position{
		Row: packDirection.Row * int(minimumDistance),
		Col: packDirection.Col * int(minimumDistance),
	}
}
