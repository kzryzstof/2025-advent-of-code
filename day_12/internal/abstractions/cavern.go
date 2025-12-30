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

func (c *Cavern) PackAll() uint {

	fmt.Println()

	catalog := ComputePermutations(
		c.GetPresents(),
		false,
	)

	catalog.PrintOptimalCombinations()

	failed := uint(0)

	for christmasTreeIndex, christmasTree := range c.christmasTrees {

		presentConfigurations := christmasTree.GetPresentConfigurations()

		presentsCount := uint(0)
		for _, presentConfiguration := range presentConfigurations {
			presentsCount += presentConfiguration.Count
		}

		fmt.Println("------------------------------------------------------------------------")
		fmt.Printf("Placing %d presents under Christmas tree %d. Region available: %d\n\n", presentsCount, christmasTreeIndex, christmasTree.Region.GetArea())

		totalRegionArea := christmasTree.Region.GetArea()
		currentRegionArea := uint(0)

		/* It is best to start with the presents that have the highest count to try and combine them first */
		for _, currentPresentConfiguration := range presentConfigurations {

			if currentPresentConfiguration.Count == 0 {
				continue
			}

			if currentRegionArea > totalRegionArea {
				/* It is over: we used all the space available */
				break
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

					currentPresentConfiguration.Count -= currentPresentCount
					otherPresentConfiguration.Count -= otherPresentCount

					combinationArea := currentPresentCount * combination.Shape.Dimension.GetArea()
					currentRegionArea += combinationArea

					if combination.OtherPresentIndex == currentPresentConfiguration.Index {
						fmt.Printf("Placed %d present(s) #%d. New=%d. Total area=%d\n", currentPresentCount+otherPresentCount, currentPresentConfiguration.Index, combinationArea, currentRegionArea)
					} else {
						fmt.Printf("Placed %d present(s) (%dx#%d combined with %dx#%d). New=%d. Total area=%d\n", currentPresentCount+otherPresentCount, currentPresentCount, currentPresentConfiguration.Index, otherPresentCount, combination.OtherPresentIndex, combinationArea, currentRegionArea)
					}

					if currentPresentConfiguration.Count == 0 {
						/* Nice: all the presents have been placed */
						break
					}
				}

				if currentPresentConfiguration.Count > 0 {
					lastCurrentPresentCount := currentPresentConfiguration.Count
					currentPresentConfiguration.Count -= lastCurrentPresentCount
					presentsArea := lastCurrentPresentCount * c.GetPresents().GetPresent(currentPresentConfiguration.Index).GetArea()
					currentRegionArea += presentsArea
					fmt.Printf("Placed the last %d present(s) #%d. New=%d. Total area=%d\n", lastCurrentPresentCount, currentPresentConfiguration.Index, presentsArea, currentRegionArea)
				}
			}
		}

		if currentRegionArea > totalRegionArea {
			failed++
			fmt.Printf("\nNo more space available under christmas tree %d. Current area: %d. Available area %d\n\n", christmasTreeIndex, currentRegionArea, totalRegionArea)
		} else {
			fmt.Printf("\nAll the presents have been successfully placed under christmas tree %d. Current area: %d. Available area %d\n\n", christmasTreeIndex, currentRegionArea, totalRegionArea)
		}
	}

	return failed
}
