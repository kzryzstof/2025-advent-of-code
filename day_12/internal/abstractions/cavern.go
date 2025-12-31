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

	if verbose {
		fmt.Println()
	}

	/*
		Combining presents together can yield a better fill ratio than manipulating them alone
	*/
	catalog := ComputePermutations(
		c.GetPresents(),
		false,
	)

	if verbose {
		catalog.PrintOptimalCombinations()
	}

	failed := uint(0)

	for christmasTreeIndex, christmasTree := range c.christmasTrees {

		presentConfigurations := christmasTree.GetPresentConfigurations()
		presentsCount := christmasTree.GetPresentsCount()

		fmt.Println("------------------------------------------------------------------------")
		fmt.Printf("Placing %d presents under Christmas tree #%d.\n\n", presentsCount, christmasTreeIndex+1)

		region := christmasTree.Region.GetSpace()

		allShapesPacked := true

		/*
			It is best to start with the presents that have the highest count to try and combine them first
		*/
		for _, currentPresentConfiguration := range presentConfigurations {

			if !allShapesPacked {
				break
			}

			if currentPresentConfiguration.Count == 0 {
				continue
			}

			fmt.Printf("Placing %d presents #%d\r", currentPresentConfiguration.Count, currentPresentConfiguration.Index)

			/* Prioritizes placing combinations of presents first */
			for currentPresentConfiguration.Count > 0 {

				for _, combination := range catalog.GetOptimalCombinations(currentPresentConfiguration.Index) {

					otherPresentConfiguration := christmasTree.GetPresentConfiguration(combination.OtherPresentIndex)

					if otherPresentConfiguration.Count == 0 {
						/* No more other presents available */
						continue
					}

					currentPresentCount := currentPresentConfiguration.Count
					otherPresentCount := otherPresentConfiguration.Count

					if combination.OtherPresentIndex == currentPresentConfiguration.Index {
						if otherPresentConfiguration.Count < 2 {
							/* Not enough presents to split */
							continue
						}

						/* We split the presents in two groups */
						otherPresentCount = uint(math.Floor(float64(otherPresentCount) / 2.0))
						currentPresentCount = otherPresentCount
					} else {
						if currentPresentCount < otherPresentCount {
							otherPresentCount = currentPresentCount
						} else if currentPresentCount > otherPresentCount {
							currentPresentCount = otherPresentCount
						}
					}

					shapesPacked := true

					for shapeNumber := uint(0); shapeNumber < currentPresentCount; shapeNumber++ {
						shapesPacked = PackShape(
							region,
							combination.PresentIndex,
							combination.Shape.GetCopy(),
							verbose,
						)
					}

					currentPresentConfiguration.Count -= currentPresentCount
					otherPresentConfiguration.Count -= otherPresentCount

					if !shapesPacked {
						allShapesPacked = false
						break
					}

					if currentPresentConfiguration.Count == 0 {
						break
					}
				}

				if !allShapesPacked {
					break
				}

				if currentPresentConfiguration.Count > 0 {
					lastCurrentPresentCount := currentPresentConfiguration.Count
					currentPresentConfiguration.Count -= lastCurrentPresentCount
					fmt.Printf("Placed the last %d present(s) #%d\n", lastCurrentPresentCount, currentPresentConfiguration.Index)
				}
			}
		}

		if !allShapesPacked {
			failed++
			fmt.Printf("\nNo more space available under christmas tree #%d\n\n", christmasTreeIndex+1)
		} else {
			fmt.Printf("\nAll the presents have been successfully placed under christmas tree #%d\n\n", christmasTreeIndex+1)
		}

		PrintShape(region)
	}

	return failed
}
