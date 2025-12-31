package algorithms

import (
	"day_12/internal/abstractions"
	"day_12/internal/io"
	"day_12/internal/maths"
	"fmt"
	"math"
)

func CombinePresents(
	fixedPresentId abstractions.PresentIndex,
	fixedCells [][]int8,
	movingPresentId abstractions.PresentIndex,
	movingCells [][]int8,
	slideOffset uint,
	verbose bool,
) abstractions.Shape {

	slidedMovingCells := movingCells

	if slideOffset > 0 {
		/* Let's slide the cells as requested */
		slidedMovingCells = maths.Slide(
			movingCells,
			maths.Vector{
				Row: int(slideOffset),
				Col: 0,
			},
			abstractions.E,
		)
	}

	if verbose {
		fmt.Printf("Packing shapes #%d and #%d:\n", fixedPresentId, movingPresentId)
		io.PrintBothCells(fixedCells, slidedMovingCells)
		fmt.Println()
	}

	/* Finds out the maximum number of empty columns between the 2 shapes that we can use to slide the shape to the left */
	colsOffset := computeColOffset(
		fixedCells,
		slidedMovingCells,
	)

	if verbose {
		fmt.Printf("Columns offset = %d\n", colsOffset)
	}

	/*
		Defines the new dimensions of the canvas that will contain the two merged shapes.
		The dimensions of the fixed shape are 3x3 since the shape has not been altered.
		The dimensions of the moving shape may not be 3x3 since the shape has been altered. So we use:
			- the `slideOffset` to know how many cells the moving shape will slide to the bottom and
			- the `colsOffset` to know how many cells the moving shape will slide to the left.
	*/

	newRowsCount := abstractions.MaximumShapeSize + slideOffset
	newColsCount := abstractions.MaximumShapeSize + (abstractions.MaximumShapeSize - colsOffset)

	newCells := maths.NewCells(
		newRowsCount,
		newColsCount,
		abstractions.E,
	)

	/* Stable shape placed at origin (0,0) */
	maths.CopyCells(
		fixedPresentId.AsUint(),
		fixedCells,
		newCells,
		maths.OriginPosition(),
		abstractions.E,
	)

	/* Packed shape translated by colsOffset */
	maths.CopyCells(
		movingPresentId.AsUint(),
		slidedMovingCells,
		newCells,
		maths.Position{Row: 0, Col: int(abstractions.MaximumShapeSize - colsOffset)},
		abstractions.E,
	)

	if verbose {
		fmt.Println()
		io.PrintCells(
			newCells,
		)
		fmt.Println()
	}

	return abstractions.NewShape(
		maths.Dimension{
			Wide: newColsCount,
			Long: newRowsCount,
		},
		newCells,
	)
}

func computeColOffset(
	fixedShape [][]int8,
	movingShape [][]int8,
) uint {

	/*
		We compute the number of empty cells between the fixed shape and the moving on each row,
		and the smallest delta tells us how far we can pack the shape
	*/
	defaultEmptyCells := 2 * abstractions.MaximumShapeSize
	minimumEmptyCells := defaultEmptyCells

	for row := uint(0); row < abstractions.MaximumShapeSize; row++ {

		/* Counts the number of empty cells on the stable shape starting from right to left */
		emptyCellsCountOnFixedShape, isEmptyCellsOnFixedShape := maths.CountCells(
			row,
			abstractions.MaximumShapeSize-1,
			maths.Vector{Row: 0, Col: -1},
			fixedShape,
			abstractions.E,
		)

		/* Counts the number of empty cells on the packed shape starting from left to right */
		emptyCellsCountOnMovingShape, isEmptyCellsOnMovingShape := maths.CountCells(
			row,
			0,
			maths.Vector{Row: 0, Col: 1},
			movingShape,
			abstractions.E,
		)

		rowEmptyCells := uint(0)

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

	return uint(math.Min(float64(minimumEmptyCells), float64(abstractions.MaximumShapeSize)))
}
