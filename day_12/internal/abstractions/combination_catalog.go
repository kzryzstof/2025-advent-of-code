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

func (c *CombinationCatalog) StoreNewCombination(
	leftIndex uint,
	rightIndex uint,
	dimension Dimension,
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

	if combination == nil || dimension.IsLessThan(combination.Dimension) {
		if removedIndex != -1 {
			c.combinations[leftIndex] = append(c.combinations[leftIndex][:removedIndex], c.combinations[leftIndex][removedIndex+1:]...)
		}
		c.combinations[leftIndex] = append(c.combinations[leftIndex], Combination{PresentIndex: leftIndex, OtherPresentIndex: rightIndex, Dimension: dimension})
	}
}

func (c *CombinationCatalog) Sort() {

	type kv struct {
		Index uint
		Value Dimension
	}

	for _, combinations := range c.combinations {

		sort.Slice(combinations, func(i, j int) bool {
			return combinations[i].Dimension.IsLessThan(combinations[j].Dimension)
		})
	}
}

func (c *CombinationCatalog) GetOptimalCombination(
	leftIndex uint,
) (int, Dimension) {
	optimalRightIndex := -1
	var optimalDimension *Dimension = nil

	for rightIndex, combination := range c.combinations[leftIndex] {
		if optimalDimension == nil {
			optimalRightIndex = int(rightIndex)
			optimalDimension = &combination.Dimension
			continue
		}

		if combination.Dimension.IsLessThan(*optimalDimension) {
			optimalRightIndex = int(rightIndex)
			optimalDimension = &combination.Dimension
		}
	}

	if optimalDimension == nil {
		panic("I guess I was not expecting that! :sweat_grin:")
	}

	return optimalRightIndex, *optimalDimension
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
			_, optimalDimension := c.GetOptimalCombination(leftIndex)
			isOptimalText := ""
			if optimalDimension.Equals(combination.Dimension) {
				isOptimalText = " (optimal)"
			}
			fmt.Printf("\tWith %d = Dimensions: %dx%d%s\n", rightIndex, combination.Dimension.Wide, combination.Dimension.Long, isOptimalText)
		}
	}

}
