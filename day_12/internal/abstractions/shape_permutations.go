package abstractions

import (
	"fmt"
)

type operation func([][]byte)

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

					for slideOffset := 0; slideOffset < MaximumShapeSize; slideOffset++ {

						slidedMovingShape := SlideShape(
							movingShape,
							Vector{
								Row: slideOffset,
								Col: 0,
							},
						)

						if verbose {
							PrintShapes(fixedShape, slidedMovingShape)
							fmt.Println()
						}

						packedShape := PackShapes(
							fixedPresent.GetIndex(),
							fixedShape,
							movingPresent.GetIndex(),
							slidedMovingShape,
							slideOffset,
							verbose,
						)

						if verbose {
							fmt.Printf("Dimensions of the new shape: %dx%d\n", packedShape.Dimension.Wide, packedShape.Dimension.Long)
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
