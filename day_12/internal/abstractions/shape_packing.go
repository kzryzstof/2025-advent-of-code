package abstractions

import (
	"fmt"
	"math"
)

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
		fmt.Printf("Packing shapes #%d and #%d:\n", fixedShapeId, movingShapeId)
		PrintShapes(fixedShape, slidedMovingShape)
		fmt.Println()
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

	/* Stable shape placed at origin (0,0) */
	PasteShape(
		fixedShapeId,
		fixedShape,
		newShape,
		0,
		0,
	)

	/* Packed shape translated by packOffset */
	PasteShape(
		movingShapeId,
		slidedMovingShape,
		newShape,
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

func PackShape(
	region [][]int8,
	shapeId uint,
	shape [][]int8,
	verbose bool,
) bool {

	insertPosition, isFound := findInsertPosition(
		region,
		shape,
	)

	if !isFound {
		fmt.Printf("Unable to insert Shape in the region\n")
		return false
	}

	if verbose {
		fmt.Printf("Insert position found at %dx%d\n", insertPosition.Row, insertPosition.Col)
	}

	placeShape := func(shapeId uint, shape [][]int8, insertPosition Position) {
		for row := 0; row < len(shape); row++ {
			for col := 0; col < len(shape[row]); col++ {

				if shape[row][col] == 0 {
					continue
				}

				rowWithOffset := row + insertPosition.Row
				colWithOffset := col + insertPosition.Col

				region[rowWithOffset][colWithOffset] = int8(shapeId)
			}
		}
	}

	placeShape(
		shapeId,
		shape,
		insertPosition,
	)

	if verbose {
		fmt.Println()
		PrintShape(
			region,
		)
		fmt.Println()
	}

	return true
}

func computeColOffset(
	fixedShape [][]int8,
	movingShape [][]int8,
) int {

	/*
		We compute the number of empty cells between the fixed shape and the moving on each row,
		and the smallest delta tells us how far we can pack the shape
	*/
	defaultEmptyCells := 2 * MaximumShapeSize
	minimumEmptyCells := defaultEmptyCells

	for row := 0; row < MaximumShapeSize; row++ {

		/* Counts the number of empty cells on the stable shape starting from right to left */
		emptyCellsCountOnFixedShape, isEmptyCellsOnFixedShape := countEmptyCells(
			row,
			MaximumShapeSize-1,
			Vector{Row: 0, Col: -1},
			fixedShape,
		)

		/* Counts the number of empty cells on the packed shape starting from left to right */
		emptyCellsCountOnMovingShape, isEmptyCellsOnMovingShape := countEmptyCells(
			row,
			0,
			Vector{Row: 0, Col: 1},
			movingShape,
		)

		rowEmptyCells := 0

		if isEmptyCellsOnFixedShape {
			/* This row has empty cells on both sides; let's add both values */
			rowEmptyCells = emptyCellsCountOnFixedShape + emptyCellsCountOnMovingShape
		} else if !isEmptyCellsOnMovingShape {
			/* This row can't be packed: no empty cells on both shapes */
			rowEmptyCells = 0
		} else {
			continue
		}

		/* Only the minimum number of empty cells can be used to move the entire packed shape without overlapping the stable shape */
		if minimumEmptyCells == defaultEmptyCells || rowEmptyCells < minimumEmptyCells {
			minimumEmptyCells = rowEmptyCells
		}
	}

	return int(math.Min(float64(minimumEmptyCells), float64(MaximumShapeSize)))
}

func findInsertPosition(
	region [][]int8,
	shape [][]int8,
) (Position, bool) {

	regionRows := len(region)
	shapeRows := len(shape)

	/*
		We compute the number of empty cells between the fixed shape and the moving on each row,
		and the smallest delta tells us how far we can pack the shape
	*/
	currentInsertPositionFound := false
	currentInsertPosition := Position{
		Row: regionRows,
		Col: len(region[0]),
	}

	for regionRow := 0; regionRow < regionRows-shapeRows+1; regionRow++ {

		fmt.Printf("Checking region row %d...\n", regionRow)

		regionCols := len(region[regionRow])

		minimumEmptyCells := regionCols

		for row := 0; row < shapeRows; row++ {

			shapeCols := len(shape[row])

			/* Counts the number of empty cells on the region starting from right to left */
			emptyCellsCountOnRegion, isEmptyCellsOnRegion := countEmptyCells(
				regionRow+row,
				/* Can't start on the last column: we need space for the shape */
				regionCols-shapeCols,
				Vector{Row: 0, Col: -1},
				region,
			)

			/* Counts the number of empty cells on the shape starting from left to right */
			emptyCellsCountOnMovingShape, _ := countEmptyCells(
				row,
				0,
				Vector{Row: 0, Col: 1},
				shape,
			)

			rowEmptyCells := 0

			if isEmptyCellsOnRegion {
				/* This row has empty cells on both sides; let's add both values */
				rowEmptyCells = (emptyCellsCountOnRegion + shapeCols - 1) + emptyCellsCountOnMovingShape
			} else {
				/* This row can't be packed: no empty cells on both shapes */
				rowEmptyCells = 0
			}

			/* Only the minimum number of empty cells can be used to move the entire packed shape without overlapping the stable shape */
			if rowEmptyCells < minimumEmptyCells {
				minimumEmptyCells = rowEmptyCells
			}
		}

		insertCol := regionCols - minimumEmptyCells

		if insertCol < currentInsertPosition.Col {
			currentInsertPositionFound = true
			currentInsertPosition = Position{
				Row: regionRow,
				Col: insertCol,
			}
			fmt.Printf("\tFound insert position in region at %dx%d\n", currentInsertPosition.Row, currentInsertPosition.Col)
		}
	}

	return currentInsertPosition, currentInsertPositionFound
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

	lastEmptyCellPosition, isPositionFound := FindLastEmptyCell(
		shape,
		initialPosition,
		direction,
	)

	if !isPositionFound {
		return 0, false
	}

	return int(math.Abs(float64(initialPosition.Col) - float64(lastEmptyCellPosition.Col))), true
}
