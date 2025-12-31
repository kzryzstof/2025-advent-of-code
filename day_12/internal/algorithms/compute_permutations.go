package algorithms

import (
	"day_12/internal/abstractions"
	"day_12/internal/maths"
	"day_12/internal/services"
	"fmt"
)

type operation func([][]int8)

// ComputePermutations /* Precomputes the combinations of presents together
//
//	The goal is to combine all the presents, in
//	all angles and keep track of the width and length
//	of the new shapes.
//
//	This way, when trying to fit them in the Christmas tree,
//	we can quickly check if there is enough space for the
//	combined shapes instead of trying all the combinations
//	of presents every time.
func ComputePermutations(
	presents *abstractions.Presents,
	verbose bool,
) *services.CombinationCatalog {
	combinationsCount := 0

	/* List of operations to apply to the presents */
	operations := []operation{
		maths.NoOp,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.VerticalFlip,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.HorizontalFlip,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.RotateClockwise,
		maths.RotateClockwise,
	}

	catalog := services.NewCombinationCatalog()

	for fixedPresent := range presents.GetAllPresents() {

		if verbose {
			fmt.Println("*************************************************")
			fmt.Printf("Looking for optimal combination for present %d\n", fixedPresent.GetIndex())
		}

		fixedShape := fixedPresent.GetShape()

		for _, applyOperationOnFixedShape := range operations {

			/* Transform the fixed shape */
			applyOperationOnFixedShape(fixedShape)

			for movingPresent := range presents.GetAllPresents() {

				movingShape := movingPresent.GetShape()

				for operationIndex, applyOperationOnMovingShape := range operations {

					if verbose {
						fmt.Printf("Packing present %d with %d (%d/%d)\r", fixedPresent.GetIndex(), movingPresent.GetIndex(), operationIndex+1, len(operations))
					}

					/* Apply the operation in-place on the right shape */
					applyOperationOnMovingShape(movingShape)

					/*
						Test packing the shape from the right while also altering the row & col
						to test additional combinations
					*/

					for slideOffset := uint(0); slideOffset < abstractions.MaximumShapeSize; slideOffset++ {

						packedShape := CombinePresents(
							fixedPresent.GetIndex(),
							fixedShape,
							movingPresent.GetIndex(),
							movingShape,
							slideOffset,
							verbose,
						)

						if verbose {
							fmt.Printf("Dimensions of the new shape: %dx%d (Ratio: %f)\n", packedShape.Dimension.Wide, packedShape.Dimension.Long, packedShape.FillRatio)
							fmt.Println()
						}

						combinationsCount++

						catalog.StoreNewShape(
							fixedPresent.GetIndex(),
							movingPresent.GetIndex(),
							packedShape,
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
