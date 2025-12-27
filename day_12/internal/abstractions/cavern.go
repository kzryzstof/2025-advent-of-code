package abstractions

import "fmt"

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
	catalog *CombinationCatalog,
) uint {
	failed := uint(0)

	for christmasTreeIndex, christmasTree := range c.christmasTrees {

		fmt.Printf("Placing presents under Christmas tree %d\n", christmasTreeIndex)

		totalRegionArea := christmasTree.Region.GetArea()
		currentRegionArea := uint(0)

		presents := christmasTree.GetPresents()

		for currentPresentIndex, currentPresentsCount := range presents {

			fmt.Printf("\tPlacing %d presents #%d\n", currentPresentsCount, currentPresentIndex)

			if currentPresentsCount == 0 {
				/* Nice */
				continue
			}

			for presents[currentPresentIndex] > 0 {
				for _, combination := range catalog.GetOptimalCombinations(currentPresentIndex) {

					if combination.OtherPresentIndex == currentPresentIndex {
						continue
					}

					otherPresentCount := presents[combination.OtherPresentIndex]

					presents[currentPresentIndex] -= currentPresentsCount
					presents[combination.OtherPresentIndex] -= otherPresentCount

					currentRegionArea += currentPresentsCount * combination.Dimension.GetArea()

					if presents[currentPresentIndex] == 0 {
						/* Nice: all the presents have been placed */
						break
					}
				}
			}
		}

		if currentRegionArea > totalRegionArea {
			failed++
			fmt.Printf("\tNo more space available for christmas tree %d. Current area: %d. Available area %d", christmasTreeIndex, currentRegionArea, totalRegionArea)
		} else {
			fmt.Printf("\tAll the presents have been successfully placed under christmas tree %d. Current area: %d. Available area %d", christmasTreeIndex, currentRegionArea, totalRegionArea)
		}
	}

	return failed
}
