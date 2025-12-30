package abstractions

import "fmt"

func PackShapes(
	fixedShapeId uint,
	fixedShape [][]byte,
	movingShapeId uint,
	movingShape [][]byte,
	slideOffset int,
	verbose bool,
) Shape {

	/* Finds out the number of empty columns between the 2 shapes */
	colsOffset := computeColOffset(
		fixedShape,
		movingShape,
	)

	if verbose {
		fmt.Printf("Columns offset = %d\n", colsOffset)
	}

	/*
		Defines the new dimensions of the canvas that will have the two merged shapes.
		The dimensions of the fixed shape are 3x3 since the shape has not been altered.
		The dimensions of the moving shape may not be 3x3 since the shape has been altered. So we use the `slideOffset`
		to know how many cells the moving shape will slide to the bottom and the `colsOffset` to know how many cells
		the moving shape will slide to the left.
	*/

	newRowsCount := CanvasSize + slideOffset
	newColsCount := CanvasSize + (CanvasSize - colsOffset)

	newShape := make([][]byte, newRowsCount)

	for row := range newShape {
		newShape[row] = make([]byte, newColsCount)
	}

	/* Helper to place a shape at an offset */
	placeShape := func(shapeId uint, shape [][]byte, rowOffset, colOffset int) {
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

				newShape[rowWithOffset][colWithOffset] = byte(shapeId)
			}
		}
	}

	/* Stable shape placed at origin (0,0) */
	placeShape(
		fixedShapeId,
		fixedShape,
		0,
		0,
	)

	/* Packed shape translated by packOffset */
	placeShape(
		movingShapeId,
		movingShape,
		slideOffset,
		CanvasSize+slideOffset,
	)

	if verbose {
		fmt.Println()
		PrintShape(
			newShape,
		)
		fmt.Println()
	}

	return Shape{
		Dimension{
			Wide: newColsCount,
			Long: newRowsCount,
		},
		newShape,
		ComputeFillRatio(newShape),
	}
}

func computeColOffset(
	fixedShape [][]byte,
	movingShape [][]byte,
) int {

	/* We compute the delta at each row, and the smallest delta tells us how far we can pack the shape */
	minimumEmptyCells := -1

	for row := 0; row < MaximumShapeSize; row++ {

		/* Counts the number of empty cells on the stable shape starting from right to left */
		emptyCellsOnStableShape := countEmptyCells(
			row,
			MaximumShapeSize-1,
			Vector{Row: 0, Col: -1},
			fixedShape,
		)

		/* Counts the number of empty cells on the packed shape starting from left to right */
		emptyCellsOnPackedShape := countEmptyCells(
			row,
			0,
			Vector{Row: 0, Col: 1},
			movingShape,
		)

		/* Compute the number of cells the packed shape can be moved to the left on this current row */
		totalEmptyCells := -emptyCellsOnStableShape + emptyCellsOnPackedShape

		/* Only the minimum number of empty cells can be used to move the entire packed shape without overlapping the stable shape */
		if minimumEmptyCells == -1 || totalEmptyCells < minimumEmptyCells {
			minimumEmptyCells = totalEmptyCells
		}
	}

	return minimumEmptyCells
}

func countEmptyCells(
	fromRow int,
	fromCol int,
	direction Vector,
	shape [][]byte,
) int {
	initialPosition := Position{
		Row: fromRow,
		Col: fromCol,
	}

	lastEmptyCellPosition := FindLastEmptyCell(
		shape,
		initialPosition,
		direction,
	)

	return initialPosition.Col - lastEmptyCellPosition.Col
}
