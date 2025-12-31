package algorithms

import (
	"day_12/internal/abstractions"
	"day_12/internal/io"
	"day_12/internal/maths"
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
) abstractions.Shape {

	slidedMovingShape := maths.Slide(
		movingShape,
		maths.Vector{
			Row: slideOffset,
			Col: 0,
		},
		abstractions.E,
	)

	if verbose {
		fmt.Printf("Packing shapes #%d and #%d:\n", fixedShapeId, movingShapeId)
		io.PrintShapes(fixedShape, slidedMovingShape)
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

	newRowsCount := abstractions.MaximumShapeSize + slideOffset
	newColsCount := abstractions.MaximumShapeSize + (abstractions.MaximumShapeSize - colsOffset)

	newShape := maths.NewSlice(newRowsCount, newColsCount, abstractions.E)

	/* Stable shape placed at origin (0,0) */
	maths.PasteShape(
		fixedShapeId,
		fixedShape,
		newShape,
		0,
		0,
	)

	/* Packed shape translated by packOffset */
	maths.PasteShape(
		movingShapeId,
		slidedMovingShape,
		newShape,
		0,
		abstractions.MaximumShapeSize-colsOffset,
	)

	if verbose {
		fmt.Println()
		io.PrintShape(
			newShape,
		)
		fmt.Println()
	}

	return abstractions.NewShape(
		maths.Dimension{
			Wide: newColsCount,
			Long: newRowsCount,
		},
		newShape,
	)
}

func PackShape(
	region [][]int8,
	shape [][]int8,
	verbose bool,
) bool {

	for row := 0; row < len(region); row++ {
		for col := 0; col < len(region[row]); col++ {
			if region[row][col] != abstractions.E {
				region[row][col] = abstractions.P
			}
		}
	}

	insertPosition, isFound := findInsertPosition(
		region,
		shape,
		verbose,
	)

	if !isFound {
		if verbose {
			fmt.Printf("\tUnable to insert Shape in the region\n")
		}
		return false
	}

	if verbose {
		fmt.Printf("\tInsert position found at %dx%d\n", insertPosition.Row, insertPosition.Col)
	}

	placeShape := func(shape [][]int8, insertPosition maths.Position) {
		for row := 0; row < len(shape); row++ {
			for col := 0; col < len(shape[row]); col++ {

				if shape[row][col] == abstractions.E {
					continue
				}

				rowWithOffset := row + insertPosition.Row
				colWithOffset := col + insertPosition.Col

				region[rowWithOffset][colWithOffset] = shape[row][col]
			}
		}
	}

	placeShape(
		shape,
		insertPosition,
	)

	if verbose {
		fmt.Println()
		io.PrintShape(
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
	defaultEmptyCells := 2 * abstractions.MaximumShapeSize
	minimumEmptyCells := defaultEmptyCells

	for row := 0; row < abstractions.MaximumShapeSize; row++ {

		/* Counts the number of empty cells on the stable shape starting from right to left */
		emptyCellsCountOnFixedShape, isEmptyCellsOnFixedShape := countEmptyCells(
			row,
			abstractions.MaximumShapeSize-1,
			maths.Vector{Row: 0, Col: -1},
			fixedShape,
		)

		/* Counts the number of empty cells on the packed shape starting from left to right */
		emptyCellsCountOnMovingShape, isEmptyCellsOnMovingShape := countEmptyCells(
			row,
			0,
			maths.Vector{Row: 0, Col: 1},
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

	return int(math.Min(float64(minimumEmptyCells), float64(abstractions.MaximumShapeSize)))
}

func findInsertPosition(
	region [][]int8,
	shape [][]int8,
	verbose bool,
) (maths.Position, bool) {

	regionRows := len(region)
	shapeRows := len(shape)

	/*
		We compute the number of empty cells between the fixed shape and the moving on each row,
		and the smallest delta tells us how far we can pack the shape
	*/
	currentInsertPositionFound := false
	currentInsertPosition := maths.Position{
		Row: regionRows,
		Col: len(region[0]),
	}

	for regionRow := 0; regionRow < regionRows-shapeRows+1; regionRow++ {

		if verbose {
			fmt.Printf("Checking region row %d...\n", regionRow)
		}

		regionCols := len(region[regionRow])

		minimumEmptyCells := regionCols

		for row := 0; row < shapeRows; row++ {

			shapeCols := len(shape[row])

			/* Counts the number of empty cells on the region starting from right to left */
			emptyCellsCountOnRegion, isEmptyCellsOnRegion := countEmptyCells(
				regionRow+row,
				/* Can't start on the last column: we need space for the shape */
				regionCols-shapeCols,
				maths.Vector{Row: 0, Col: -1},
				region,
			)

			/* Counts the number of empty cells on the shape starting from left to right */
			emptyCellsCountOnMovingShape, _ := countEmptyCells(
				row,
				0,
				maths.Vector{Row: 0, Col: 1},
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
			currentInsertPosition = maths.Position{
				Row: regionRow,
				Col: insertCol,
			}

			if verbose {
				fmt.Printf("\tFound insert position in region at %dx%d\n", currentInsertPosition.Row, currentInsertPosition.Col)
			}
		}
	}

	return currentInsertPosition, currentInsertPositionFound
}

func countEmptyCells(
	fromRow int,
	fromCol int,
	direction maths.Vector,
	shape [][]int8,
) (int, bool) {

	initialPosition := maths.Position{
		Row: fromRow,
		Col: fromCol,
	}

	lastEmptyCellPosition, isPositionFound := maths.FindLastCellWithValueOnRow(
		shape,
		initialPosition,
		direction,
		abstractions.E,
	)

	if !isPositionFound {
		return 0, false
	}

	return int(math.Abs(float64(initialPosition.Col) - float64(lastEmptyCellPosition.Col))), true
}
