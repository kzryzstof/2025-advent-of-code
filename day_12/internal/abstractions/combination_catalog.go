package abstractions

import "fmt"

/*
	Stores combinations of presents (stored by index)
*/

type CombinationCatalog struct {
	combinations map[uint]map[uint]Dimension
}

func NewCombinationCatalog() *CombinationCatalog {
	return &CombinationCatalog{
		combinations: make(map[uint]map[uint]Dimension),
	}
}

func (c *CombinationCatalog) StoreNewCombination(
	leftIndex uint,
	rightIndex uint,
	dimension Dimension,
) {
	if _, ok := c.combinations[leftIndex]; !ok {
		c.combinations[leftIndex] = make(map[uint]Dimension)
	}

	storedDimension := c.combinations[leftIndex][rightIndex]

	if storedDimension.IsEmpty() || dimension.IsLessThan(storedDimension) {
		c.combinations[leftIndex][rightIndex] = dimension
	}
}

func (c *CombinationCatalog) GetOptimalCombination(
	leftIndex uint,
) (int, Dimension) {
	optimalRightIndex := -1
	var optimalDimension *Dimension = nil

	for rightIndex, dimension := range c.combinations[leftIndex] {
		if optimalDimension == nil {
			optimalRightIndex = int(rightIndex)
			optimalDimension = &dimension
			continue
		}

		if dimension.IsLessThan(*optimalDimension) {
			optimalRightIndex = int(rightIndex)
			optimalDimension = &dimension
		}
	}

	if optimalDimension == nil {
		panic("I guess I was not expecting that! :sweat_grin:")
	}

	return optimalRightIndex, *optimalDimension
}

func (c *CombinationCatalog) PrintOptimalCombinations() {

	for leftIndex, combinations := range c.combinations {
		fmt.Printf("Present %d\n", leftIndex)
		for rightIndex, combination := range combinations {
			_, optimalCombination := c.GetOptimalCombination(leftIndex)
			isOptimalText := ""
			if optimalCombination.Equals(combination) {
				isOptimalText = " (optimal)"
			}
			fmt.Printf("\tWith %d = Dimensions: %dx%d%s\n", rightIndex, combination.Wide, combination.Long, isOptimalText)
		}
	}

}
