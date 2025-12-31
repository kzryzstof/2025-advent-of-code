package abstractions

import (
	"fmt"
	"math"
)

type Cavern struct {
	presents       *Presents
	christmasTrees []*ChristmasTree
}

func NewCavern(
	presents *Presents,
	christmasTrees []*ChristmasTree,
) *Cavern {
	return &Cavern{
		presents,
		christmasTrees,
	}
}

func (c *Cavern) GetPresentsCount() uint {
	return c.presents.Count()
}

func (c *Cavern) GetPresents() *Presents {
	return c.presents
}

func (c *Cavern) GetChristmasTreesCount() uint {
	return uint(len(c.christmasTrees))
}

func (c *Cavern) PackAll(
	verbose bool,
) uint {

	failedChristmasTrees := make([]uint, 0)

	if verbose {
		fmt.Println()
	}

	/* Combining presents together can yield a better fill ratio than manipulating them alone */
	catalog := ComputePermutations(
		c.GetPresents(),
		false,
	)

	if verbose {
		catalog.PrintOptimalCombinations()
	}

	failed := uint(0)

	for christmasTreeIndex, christmasTree := range c.christmasTrees {

		fmt.Println("------------------------------------------------------------------------")
		fmt.Printf("Placing %d presents under Christmas tree #%d (%dx%d).\n\n", christmasTree.GetPresentsCount(), christmasTreeIndex+1, christmasTree.wide, christmasTree.long)

		region := christmasTree.Region.GetSpace()

		/* It is best to start with the combined presents that have the highest fill ratio */
		for _, index := range catalog.GetCombinationsOrderByFillRatio() {

			currentPresentConfiguration := christmasTree.GetPresentConfiguration(index)

			if currentPresentConfiguration.Count == 0 {
				continue
			}

			fmt.Printf("Present %d | Placing %d combined presents\n", currentPresentConfiguration.Index, currentPresentConfiguration.Count)

			/* Prioritizes placing combinations of presents first */
			for _, combination := range catalog.GetOptimalCombinations(currentPresentConfiguration.Index) {

				otherPresentConfiguration := christmasTree.GetPresentConfiguration(combination.OtherPresentIndex)

				if otherPresentConfiguration.Count == 0 {
					/* No more other presents available */
					continue
				}

				fmt.Printf("\tUsing combination with presents %d (%.2f)\n", combination.OtherPresentIndex, combination.Shape.FillRatio)

				presentsCount := currentPresentConfiguration.Count
				otherPresentCount := otherPresentConfiguration.Count

				if combination.OtherPresentIndex == currentPresentConfiguration.Index {
					if otherPresentConfiguration.Count < 2 {
						/* Not enough presents to split */
						continue
					}

					/* We split the presents in two groups */
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
						shape,
						verbose,
					)

					if !shapesPacked {
						break
					}

					shapesPackedCount++
				}

				currentPresentConfiguration.Count -= shapesPackedCount
				otherPresentConfiguration.Count -= shapesPackedCount

				if !shapesPacked {
					if combination.OtherPresentIndex == currentPresentConfiguration.Index {
						fmt.Printf("\tOnly placed %d presents #%d with presents #%d instead of %d\n", shapesPackedCount, currentPresentConfiguration.Index, otherPresentConfiguration.Index, currentPresentConfiguration.Count)
					} else {
						fmt.Printf("\tOnly placed %d presents #%d instead of %d\n", shapesPackedCount, currentPresentConfiguration.Index, currentPresentConfiguration.Count)
					}
					break
				} else {
					if combination.OtherPresentIndex == currentPresentConfiguration.Index {
						fmt.Printf("\t%d presents #%d have been placed\n", 2*shapesPackedCount, currentPresentConfiguration.Index)
					} else {
						fmt.Printf("\t%d presents #%d combined with presents #%d have been placed\n", shapesPackedCount, currentPresentConfiguration.Index, otherPresentConfiguration.Index)
					}
				}

				if currentPresentConfiguration.Count == 0 {
					break
				}
			}
		}

		allShapesPacked := true

		/* Now place the remaining individual presents */
		for _, index := range catalog.GetCombinationsOrderByFillRatio() {

			currentPresentConfiguration := christmasTree.GetPresentConfiguration(index)

			if currentPresentConfiguration.Count == 0 {
				continue
			}

			fmt.Printf("Placing presents #%d individually \n", currentPresentConfiguration.Index)

			shapePacked := true
			shapesPackedCount := uint(0)

			shape := c.presents.GetPresent(currentPresentConfiguration.Index).GetShape()

			for presentCount := uint(0); presentCount < currentPresentConfiguration.Count; presentCount++ {

				shapePacked = PackShape(
					region,
					shape,
					verbose,
				)

				if !shapePacked {
					allShapesPacked = false
					fmt.Printf("\tOnly placed %d presents #%d instead of %d\n", shapesPackedCount, currentPresentConfiguration.Index, currentPresentConfiguration.Count)
					break
				}

				shapesPackedCount++
			}

			currentPresentConfiguration.Count -= shapesPackedCount

			if !allShapesPacked {
				break
			}

			fmt.Printf("\tPlaced %d presents #%d\n", shapesPackedCount, currentPresentConfiguration.Index)
		}

		if !allShapesPacked {
			failed++
			failedChristmasTrees = append(failedChristmasTrees, uint(christmasTreeIndex))
			fmt.Printf("\nNo more space available under christmas tree #%d\n\n", christmasTreeIndex+1)
			PrintShape(region)

		} else {
			fmt.Printf("\nAll the presents have been successfully placed under christmas tree #%d\n\n", christmasTreeIndex+1)
		}
	}

	for _, failedChristmasTreeIndex := range failedChristmasTrees {
		fmt.Printf("Christmas tree #%d has been failed\n", failedChristmasTreeIndex+1)
	}

	return failed
}
