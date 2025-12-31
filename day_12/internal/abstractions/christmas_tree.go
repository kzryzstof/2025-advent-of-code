package abstractions

import (
	"day_12/internal/maths"
	"sort"
)

type ChristmasTree struct {
	/* Size available under the tree */
	wide uint
	long uint

	/* Lists all the present configurations in descending order of count */
	sortedPresentConfigurations []*PresentConfiguration
	mappedPresentConfigurations map[uint]*PresentConfiguration

	Region *maths.Region
}

func NewChristmasTree(
	wide uint,
	long uint,
	presentConfigurations map[uint]uint,
) *ChristmasTree {

	mappedPresentConfigurations := make(map[uint]*PresentConfiguration)
	sortedPresentConfigurations := make([]*PresentConfiguration, 0)

	for presentIndex, presentCount := range presentConfigurations {

		presentConfiguration := &PresentConfiguration{
			Index: presentIndex,
			Count: presentCount,
		}

		sortedPresentConfigurations = append(sortedPresentConfigurations, presentConfiguration)
		mappedPresentConfigurations[presentIndex] = presentConfiguration
	}

	sort.Slice(sortedPresentConfigurations, func(i, j int) bool {
		return sortedPresentConfigurations[i].Count > sortedPresentConfigurations[j].Count
	})

	return &ChristmasTree{
		wide,
		long,
		sortedPresentConfigurations,
		mappedPresentConfigurations,
		maths.NewRegion(wide, long, E),
	}
}

func (ct *ChristmasTree) GetWide() uint {
	return ct.wide
}

func (ct *ChristmasTree) GetLong() uint {
	return ct.long
}

func (ct *ChristmasTree) GetPresentConfiguration(
	presentIndex uint,
) *PresentConfiguration {
	return ct.mappedPresentConfigurations[presentIndex]
}

func (ct *ChristmasTree) GetPresentsCount() uint {
	presentsCount := uint(0)

	for _, presentConfiguration := range ct.sortedPresentConfigurations {
		presentsCount += presentConfiguration.Count
	}

	return presentsCount
}
