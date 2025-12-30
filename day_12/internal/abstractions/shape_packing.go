package abstractions

import "fmt"

func PackShapes(
	fixedShapeId uint,
	fixedShape [][]int8,
	movingShapeId uint,
	movingShape [][]int8,
	slideOffset int,
	verbose bool,
) Shape {

	slidedMovingShape := SlideShape(
		movingShape,
		Vector{
			Row: slideOffset,
			Col: 0,
		},
	)

	if verbose {
		fmt.Println("Slided moving shape:")
		PrintShape(slidedMovingShape)
	}

	/* Finds out the number of empty columns between the 2 shapes */
	colsOffset := computeColOffset(
		fixedShape,
		slidedMovingShape,
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

	newRowsCount := MaximumShapeSize + slideOffset
	newColsCount := MaximumShapeSize + (MaximumShapeSize - colsOffset)

	newShape := make([][]int8, newRowsCount)

	for row := range newShape {
		newShape[row] = make([]int8, newColsCount)
	}

	/* Helper to place a shape at an offset */
	placeShape := func(shapeId uint, shape [][]int8, rowOffset, colOffset int) {
		for row := 0; row < len(shape); row++ {
			for col := 0; col < len(shape[row]); col++ {

				if shape[row][col] == 0 {
					continue
				}

				rowWithOffset := row + rowOffset
				colWithOffset := col + colOffset

				newShape[rowWithOffset][colWithOffset] = int8(shapeId)
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
		slidedMovingShape,
		0,
		MaximumShapeSize-colsOffset,
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
	fixedShape [][]int8,
	movingShape [][]int8,
) int {

	/* We compute the delta at each row, and the smallest delta tells us how far we can pack the shape */
	minimumEmptyCells := MaximumShapeSize

	for row := 0; row < MaximumShapeSize; row++ {

		/* Counts the number of empty cells on the stable shape starting from right to left */
		emptyCellsOnFixedShape, _ := countEmptyCells(
			row,
			MaximumShapeSize-1,
			Vector{Row: 0, Col: -1},
			fixedShape,
		)

		/* Counts the number of empty cells on the packed shape starting from left to right */
		emptyCellsOnMovingShape, found := countEmptyCells(
			row,
			0,
			Vector{Row: 0, Col: 1},
			movingShape,
		)

		if !found {
			/* If there are only cells on the moving shape, then we just ignore it */
			continue
		}

		/* Compute the number of cells the packed shape can be moved to the left on this current row */
		totalEmptyCells := emptyCellsOnFixedShape + emptyCellsOnMovingShape

		/* Only the minimum number of empty cells can be used to move the entire packed shape without overlapping the stable shape */
		if minimumEmptyCells == MaximumShapeSize || totalEmptyCells > minimumEmptyCells {
			minimumEmptyCells = totalEmptyCells
		}
	}

	return minimumEmptyCells
}

func countEmptyCells(
	fromRow int,
	fromCol int,
	direction Vector,
	shape [][]int8,
) (int, bool) {
	initialPosition := Position{
		Row: fromRow,
		Col: fromCol,
	}

	lastEmptyCellPosition := FindLastEmptyCell(
		shape,
		initialPosition,
		direction,
	)

	if lastEmptyCellPosition.Col == MaximumShapeSize {
		return 0, false
	}

	return initialPosition.Col - lastEmptyCellPosition.Col, true
}
