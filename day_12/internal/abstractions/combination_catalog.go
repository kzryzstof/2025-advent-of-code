package abstractions

import (
	"fmt"
	"sort"
)

/*
	Stores combinations of presents (stored by index)
*/

type combinationMetadata struct {
	presentIndex uint
	fillRatioAvg float64
}

type CombinationCatalog struct {
	combinations map[uint][]Combination
	metadata     []combinationMetadata
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

	c.updateFillRatios()
	c.sort()
}

func (c *CombinationCatalog) updateFillRatios() {

	c.metadata = make([]combinationMetadata, 0)

	for leftIndex, combinations := range c.combinations {
		totalFillRatio := float64(0)
		fillRatioCount := float64(0)
		for _, combination := range combinations {
			totalFillRatio += combination.Shape.FillRatio
			fillRatioCount++
		}
		c.metadata = append(c.metadata, combinationMetadata{leftIndex, totalFillRatio / fillRatioCount})
	}

	sort.Slice(c.metadata, func(i, j int) bool { return c.metadata[i].fillRatioAvg > c.metadata[j].fillRatioAvg })
}

func (c *CombinationCatalog) sort() {

	for _, combinations := range c.combinations {
		sort.Slice(combinations, func(i, j int) bool {
			isEqual := combinations[i].Shape.FillRatio == combinations[j].Shape.FillRatio

			if isEqual {
				return combinations[i].PresentIndex != combinations[j].OtherPresentIndex
			}

			return combinations[i].Shape.IsMoreOptimalThan(combinations[j].Shape)
		})
	}
}

func (c *CombinationCatalog) GetCombinations() []uint {

	sortedCombinations := make([]uint, 0)

	for _, combinations := range c.metadata {
		sortedCombinations = append(sortedCombinations, combinations.presentIndex)
	}

	return sortedCombinations
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

	for _, metadata := range c.metadata {
		leftIndex := metadata.presentIndex
		combinations := c.combinations[leftIndex]
		totalFillRatio := float64(0)
		fmt.Printf("Present %d\n", leftIndex)
		for rightIndex, combination := range combinations {
			_, optimalShape := c.GetOptimalCombination(leftIndex)
			isOptimalText := ""
			if optimalShape.Dimension.Equals(combination.Shape.Dimension) {
				isOptimalText = " (optimal)"
			}
			totalFillRatio += combination.Shape.FillRatio
			fmt.Printf("\tWith %d = Dimensions: %dx%d with fill ratio %.2f%s\n", rightIndex, combination.Shape.Dimension.Wide, combination.Shape.Dimension.Long, combination.Shape.FillRatio, isOptimalText)
		}
		fmt.Printf("\tAvg fill ratio: %.2f\n", totalFillRatio/6)
	}

	fmt.Print("\n")
}
