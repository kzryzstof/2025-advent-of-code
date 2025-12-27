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

	failed := uint(0)

	for christmasTreeIndex, christmasTree := range c.christmasTrees {

		presents := christmasTree.GetPresents()

		presentsCount := uint(0)
		for _, count := range presents {
			presentsCount += count
		}

		fmt.Printf("Placing %d presents under Christmas tree %d. Region available: %d\n", presentsCount, christmasTreeIndex, christmasTree.Region.GetArea())

		catalog := ComputePermutations(
			c.GetPresents(),
			christmasTree.Region,
			false,
		)

		catalog.PrintOptimalCombinations(christmasTree.Region)

		totalRegionArea := christmasTree.Region.GetArea()
		currentRegionArea := uint(0)

		for currentPresentIndex, currentPresentsCount := range presents {

			if currentPresentsCount == 0 {
				/* Nice */
				continue
			}

			if currentRegionArea > totalRegionArea {
				/* It is over */
				break
			}

			fmt.Printf("\tPlacing %d presents #%d\r", currentPresentsCount, currentPresentIndex)

			/* Prioritizes placing combinations of presents first */
			for presents[currentPresentIndex] > 0 {

				for _, combination := range catalog.GetOptimalCombinations(currentPresentIndex) {

					otherPresentCount := presents[combination.OtherPresentIndex]

					if otherPresentCount == 0 {
						/* No more other presents available */
						continue
					}

					if combination.OtherPresentIndex == currentPresentIndex {
						if otherPresentCount < 2 {
							/* Not enough presents to split */
							continue
						}

						/* We split the presents in two groups */
						otherPresentCount = uint(math.Floor(float64(otherPresentCount) / 2.0))
						currentPresentsCount = otherPresentCount
					} else {
						if currentPresentsCount < otherPresentCount {
							otherPresentCount = currentPresentsCount
						} else if otherPresentCount > currentPresentsCount {
							currentPresentsCount = otherPresentCount
						}
					}

					presents[currentPresentIndex] -= currentPresentsCount
					presents[combination.OtherPresentIndex] -= otherPresentCount

					currentRegionArea += currentPresentsCount * combination.Dimension.GetArea()

					fmt.Printf("\tPlaced %d presents #%d and other presents #%d. Area %d\n", currentPresentsCount, currentPresentIndex, combination.OtherPresentIndex, currentRegionArea)

					if presents[currentPresentIndex] == 0 {
						/* Nice: all the presents have been placed */
						break
					}
				}

				remainingPresentsCount := presents[currentPresentIndex]

				if remainingPresentsCount > 0 {
					presents[currentPresentIndex] -= remainingPresentsCount
					currentRegionArea += remainingPresentsCount * c.GetPresents().GetPresent(currentPresentIndex).GetArea()
					fmt.Printf("\tPlaced the last %d presents #%d. Area %d\n", remainingPresentsCount, currentPresentIndex, currentRegionArea)
				}
			}
		}

		if currentRegionArea > totalRegionArea {
			failed++
			fmt.Printf("\tNo more space available for christmas tree %d. Current area: %d. Available area %d\n\n", christmasTreeIndex, currentRegionArea, totalRegionArea)
		} else {
			fmt.Printf("\tAll the presents have been successfully placed under christmas tree %d. Current area: %d. Available area %d\n\n", christmasTreeIndex, currentRegionArea, totalRegionArea)
		}
	}

	return failed
}
