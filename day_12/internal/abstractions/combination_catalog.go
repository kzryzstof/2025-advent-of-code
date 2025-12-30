package abstractions

import (
	"fmt"
	"sort"
)

/*
	Stores combinations of presents (stored by index)
*/

type CombinationCatalog struct {
	combinations map[uint][]Combination
}

func NewCombinationCatalog() *CombinationCatalog {
	return &CombinationCatalog{
		combinations: make(map[uint][]Combination),
	}
}

func (c *CombinationCatalog) StoreNewShape(
	leftIndex uint,
	rightIndex uint,
	shape Shape,
) {
	if _, ok := c.combinations[leftIndex]; !ok {
		c.combinations[leftIndex] = make([]Combination, 0)
	}

	var combination *Combination = nil

	removedIndex := -1

	for existingIndex, existingCombination := range c.combinations[leftIndex] {
		if existingCombination.OtherPresentIndex == rightIndex {
			combination = &existingCombination
			removedIndex = existingIndex
		}
	}

	if combination == nil || shape.IsMoreOptimalThan(combination.Shape) {
		if removedIndex != -1 {
			c.combinations[leftIndex] = append(c.combinations[leftIndex][:removedIndex], c.combinations[leftIndex][removedIndex+1:]...)
		}
		c.combinations[leftIndex] = append(
			c.combinations[leftIndex],
			Combination{
				PresentIndex:      leftIndex,
				OtherPresentIndex: rightIndex,
				Shape:             shape,
			})
	}

	c.sort()
}

func (c *CombinationCatalog) sort() {

	for _, combinations := range c.combinations {
		sort.Slice(combinations, func(i, j int) bool {
			return combinations[i].Shape.IsMoreOptimalThan(combinations[j].Shape)
		})
	}
}

func (c *CombinationCatalog) GetOptimalCombination(
	leftIndex uint,
) (int, Shape) {

	optimalRightIndex := -1
	var optimalShape *Shape = nil

	for rightIndex, combination := range c.combinations[leftIndex] {
		if optimalShape == nil {
			optimalRightIndex = rightIndex
			optimalShape = &combination.Shape
			continue
		}

		if combination.Shape.IsMoreOptimalThan(*optimalShape) {
			optimalRightIndex = rightIndex
			optimalShape = &combination.Shape
		}
	}

	if optimalShape == nil {
		panic("I guess I was not expecting that! :sweat_grin:")
	}

	return optimalRightIndex, *optimalShape
}

func (c *CombinationCatalog) GetOptimalCombinations(
	leftIndex uint,
) []Combination {

	return c.combinations[leftIndex]
}

func (c *CombinationCatalog) PrintOptimalCombinations() {

	for leftIndex, combinations := range c.combinations {
		fmt.Printf("Present %d\n", leftIndex)
		for rightIndex, combination := range combinations {
			_, optimalShape := c.GetOptimalCombination(leftIndex)
			isOptimalText := ""
			if optimalShape.Dimension.Equals(combination.Shape.Dimension) {
				isOptimalText = " (optimal)"
			}
			fmt.Printf("\tWith %d = Dimensions: %dx%d with fill ratio %.2f%s\n", rightIndex, combination.Shape.Dimension.Wide, combination.Shape.Dimension.Long, combination.Shape.Dimension.FillRatio, isOptimalText)
		}
	}

	fmt.Print("\n")
}
