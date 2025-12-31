package algorithms

import (
	"day_12/internal/abstractions"
	"day_12/internal/io"
	"day_12/internal/maths"
	"fmt"
)

func PackShape(
	region [][]int8,
	presentId abstractions.PresentIndex,
	cells [][]int8,
	verbose bool,
) bool {

	markExistingShapeCellAsPrevious(region)

	insertPosition, isFound := findInsertPosition(
		region,
		cells,
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

	maths.CopyCells(
		presentId.AsUint(),
		cells,
		region,
		insertPosition,
		abstractions.E,
	)

	if verbose {
		fmt.Println()
		io.PrintCells(
			region,
		)
		fmt.Println()
	}

	return true
}

func markExistingShapeCellAsPrevious(
	region [][]int8,
) {
	for row := 0; row < len(region); row++ {
		for col := 0; col < len(region[row]); col++ {
			if region[row][col] != abstractions.E {
				region[row][col] = abstractions.P
			}
		}
	}
}

func findInsertPosition(
	region [][]int8,
	shape [][]int8,
	verbose bool,
) (maths.Position, bool) {

	regionRows := uint(len(region))
	shapeRows := uint(len(shape))

	/*
		We compute the number of empty cells between the fixed shape and the moving on each row,
		and the smallest delta tells us how far we can pack the shape
	*/
	currentInsertPositionFound := false
	currentInsertPosition := maths.Position{
		Row: int(regionRows),
		Col: len(region[0]),
	}

	for regionRow := uint(0); regionRow < regionRows-shapeRows+1; regionRow++ {

		regionCols := uint(len(region[regionRow]))

		minimumEmptyCells := regionCols

		for row := uint(0); row < shapeRows; row++ {

			shapeCols := uint(len(shape[row]))

			/* Counts the number of empty cells on the region starting from right to left */
			emptyCellsCountOnRegion, isEmptyCellsOnRegion := maths.CountCells(
				regionRow+row,
				/* Can't start on the last column: we need space for the shape */
				regionCols-shapeCols,
				maths.Vector{Row: 0, Col: -1},
				region,
				abstractions.E,
			)

			/* Counts the number of empty cells on the shape starting from left to right */
			emptyCellsCountOnMovingShape, _ := maths.CountCells(
				row,
				0,
				maths.Vector{Row: 0, Col: 1},
				shape,
				abstractions.E,
			)

			rowEmptyCells := uint(0)

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

		if insertCol < uint(currentInsertPosition.Col) {
			currentInsertPositionFound = true
			currentInsertPosition = maths.Position{
				Row: int(regionRow),
				Col: int(insertCol),
			}

			if verbose {
				fmt.Printf("\tFound insert position in region at %dx%d\n", currentInsertPosition.Row, currentInsertPosition.Col)
			}
		}
	}

	return currentInsertPosition, currentInsertPositionFound
}
