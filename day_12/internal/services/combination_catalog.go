package services

import (
	"day_12/internal/abstractions"
	"fmt"
	"sort"
)

type combinationMetadata struct {
	presentIndex abstractions.PresentIndex
	fillRatioAvg float64
}

type CombinationCatalog struct {
	combinations map[abstractions.PresentIndex][]abstractions.Combination
	metadata     []combinationMetadata
}

func NewCombinationCatalog() *CombinationCatalog {
	return &CombinationCatalog{
		combinations: make(map[abstractions.PresentIndex][]abstractions.Combination),
	}
}

func (c *CombinationCatalog) StoreNewShape(
	leftIndex abstractions.PresentIndex,
	rightIndex abstractions.PresentIndex,
	shape abstractions.Shape,
) {
	if _, ok := c.combinations[leftIndex]; !ok {
		c.combinations[leftIndex] = make([]abstractions.Combination, 0)
	}

	var combination *abstractions.Combination = nil

	removedIndex := -1

	for existingIndex, existingCombination := range c.combinations[leftIndex] {
		if existingCombination.OtherIndex == rightIndex {
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
			abstractions.Combination{
				Index:      leftIndex,
				OtherIndex: rightIndex,
				Shape:      shape,
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
			if combination.Shape.FillRatio == 1 {
				/*
					Adds a lot more weight to combinations with a fill ratio of 1
					It helps prioritize the packing of the combination with more fill ratio of 1
				*/
				totalFillRatio += 1
			}
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
				return combinations[i].Index != combinations[j].OtherIndex
			}

			return combinations[i].Shape.IsMoreOptimalThan(combinations[j].Shape)
		})
	}
}

func (c *CombinationCatalog) GetCombinationsOrderByFillRatio() []abstractions.PresentIndex {

	sortedCombinations := make([]abstractions.PresentIndex, 0)

	for _, combinations := range c.metadata {
		sortedCombinations = append(sortedCombinations, combinations.presentIndex)
	}

	return sortedCombinations
}

func (c *CombinationCatalog) GetOptimalCombination(
	leftIndex abstractions.PresentIndex,
) (abstractions.PresentIndex, abstractions.Shape) {

	var optimalRightIndex = abstractions.NoPresentIndex()

	var optimalShape *abstractions.Shape = nil

	for _, combination := range c.combinations[leftIndex] {
		if optimalShape == nil {
			optimalRightIndex = combination.OtherIndex
			optimalShape = &combination.Shape
			continue
		}

		if combination.Shape.IsMoreOptimalThan(*optimalShape) {
			optimalRightIndex = combination.OtherIndex
			optimalShape = &combination.Shape
		}
	}

	if optimalShape == nil {
		panic("I guess I was not expecting that! :sweat_grin:")
	}

	return optimalRightIndex, *optimalShape
}

func (c *CombinationCatalog) GetOptimalCombinations(
	leftIndex abstractions.PresentIndex,
) []abstractions.Combination {

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
