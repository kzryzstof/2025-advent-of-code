package algorithms

import (
	"day_12/internal/abstractions"
	"day_12/internal/services"
	"fmt"
	"math"
)

func PackAll(
	cavern *abstractions.Cavern,
	verbose bool,
) uint {

	if verbose {
		fmt.Println()
	}

	/* Combining presents together yields a better fill ratio than manipulating them alone */
	catalog := ComputePermutations(
		cavern.GetPresents(),
		false,
	)

	if verbose {
		catalog.PrintOptimalCombinations()
	}

	failed := uint(0)

	for _, christmasTree := range cavern.GetChristmasTrees() {

		allPresentsPacked := packPresentsUnderChristmasTree(
			cavern,
			christmasTree,
			catalog,
			verbose,
		)

		if !allPresentsPacked {
			failed++
			fmt.Printf("\nNo more space available under christmas tree #%d\n", christmasTree.Index)
		} else {
			fmt.Printf("\nAll the presents have been successfully placed under christmas tree #%d\n", christmasTree.Index)
		}
	}

	return failed
}

func packPresentsUnderChristmasTree(
	cavern *abstractions.Cavern,
	christmasTree *abstractions.ChristmasTree,
	catalog *services.CombinationCatalog,
	verbose bool,
) bool {

	if verbose {
		fmt.Println("------------------------------------------------------------------------")
		fmt.Printf("Placing %d presents under Christmas tree #%d (%dx%d).\n", christmasTree.GetPresentsCount(), christmasTree.Index, christmasTree.Dimension.Wide, christmasTree.Dimension.Long)
	}

	region := christmasTree.Region.GetSpace()

	/* Start packing the region with combined presents that have the highest fill ratio */
	processCombinations(
		catalog,
		christmasTree,
		region,
		verbose,
	)

	/* Now process the remaining individual presents */
	return processRemainingIndividualPresents(
		cavern,
		catalog,
		christmasTree,
		region,
		verbose,
	)
}

func processRemainingIndividualPresents(
	cavern *abstractions.Cavern,
	catalog *services.CombinationCatalog,
	christmasTree *abstractions.ChristmasTree,
	region [][]int8,
	verbose bool,
) bool {

	allPresentsPacked := true

	for _, index := range catalog.GetCombinationsOrderByFillRatio() {

		currentPresentConfiguration := christmasTree.GetPresentConfiguration(index)

		if currentPresentConfiguration.Count == 0 {
			continue
		}

		if verbose {
			fmt.Printf("Placing presents #%d individually \n", currentPresentConfiguration.Index)
		}

		shapePacked := true
		shapesPackedCount := uint(0)

		present := cavern.GetPresents().GetPresent(currentPresentConfiguration.Index)
		shape := present.GetShape()

		for presentCount := uint(0); presentCount < currentPresentConfiguration.Count; presentCount++ {

			shapePacked = PackShape(
				region,
				present.GetIndex(),
				shape,
				verbose,
			)

			if !shapePacked {
				allPresentsPacked = false
				if verbose {
					fmt.Printf("\tOnly placed %d presents #%d instead of %d\n", shapesPackedCount, currentPresentConfiguration.Index, currentPresentConfiguration.Count)
				}
				break
			}

			shapesPackedCount++
		}

		currentPresentConfiguration.Count -= shapesPackedCount

		if !allPresentsPacked {
			break
		}

		if verbose {
			fmt.Printf("\tPlaced %d presents #%d\n", shapesPackedCount, currentPresentConfiguration.Index)
		}
	}

	return allPresentsPacked
}

func processCombinations(
	catalog *services.CombinationCatalog,
	christmasTree *abstractions.ChristmasTree,
	region [][]int8,
	verbose bool,
) {

	for _, index := range catalog.GetCombinationsOrderByFillRatio() {

		currentPresentConfiguration := christmasTree.GetPresentConfiguration(index)

		if currentPresentConfiguration.Count == 0 {
			/* The number of presents could be 0 since it has been combined and placed with another present */
			continue
		}

		if verbose {
			fmt.Printf("Present %d | Placing %d combined presents\n", currentPresentConfiguration.Index, currentPresentConfiguration.Count)
		}

		for _, combination := range catalog.GetOptimalCombinations(currentPresentConfiguration.Index) {

			otherPresentConfiguration := christmasTree.GetPresentConfiguration(combination.OtherIndex)

			if otherPresentConfiguration.Count == 0 {
				/* This combination cannot be used: there are no more other presents available */
				continue
			}

			if verbose {
				fmt.Printf("\tUsing combination with presents %d (%.2f)\n", combination.OtherIndex, combination.Shape.FillRatio)
			}

			presentsCount := currentPresentConfiguration.Count
			otherPresentCount := otherPresentConfiguration.Count

			if combination.OtherIndex == currentPresentConfiguration.Index {

				if otherPresentConfiguration.Count < 2 {
					/* Not enough presents to split */
					continue
				}

				/* We split the presents count in two */
				presentsCount = uint(math.Floor(float64(presentsCount) / 2.0))

			} else if presentsCount > otherPresentCount {
				presentsCount = otherPresentCount
			}

			shapesPacked := true
			shapesPackedCount := uint(0)

			shape := combination.Shape.GetCopy()

			for shapeNumber := uint(0); shapeNumber < presentsCount; shapeNumber++ {

				shapesPacked = PackShape(
					region,
					combination.Index,
					shape,
					verbose,
				)

				if !shapesPacked {
					/* No more space available */
					break
				}

				shapesPackedCount++
			}

			currentPresentConfiguration.Count -= shapesPackedCount
			otherPresentConfiguration.Count -= shapesPackedCount

			if !shapesPacked {
				if verbose {
					if combination.OtherIndex == currentPresentConfiguration.Index {
						fmt.Printf("\tOnly placed %d presents #%d with presents #%d instead of %d\n", shapesPackedCount, currentPresentConfiguration.Index, otherPresentConfiguration.Index, currentPresentConfiguration.Count)
					} else {
						fmt.Printf("\tOnly placed %d presents #%d instead of %d\n", shapesPackedCount, currentPresentConfiguration.Index, currentPresentConfiguration.Count)
					}
				}
				break
			} else {
				if verbose {
					if combination.OtherIndex == currentPresentConfiguration.Index {
						fmt.Printf("\t%d presents #%d have been placed\n", 2*shapesPackedCount, currentPresentConfiguration.Index)
					} else {
						fmt.Printf("\t%d presents #%d combined with presents #%d have been placed\n", shapesPackedCount, currentPresentConfiguration.Index, otherPresentConfiguration.Index)
					}
				}
			}

			if currentPresentConfiguration.Count == 0 {
				break
			}
		}
	}
}
