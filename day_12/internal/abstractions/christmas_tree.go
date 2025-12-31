package abstractions

import (
	"day_12/internal/maths"
)

type ChristmasTree struct {
	/* Size available under the tree */
	wide uint
	long uint

	/* Lists all the present configurations mapped by their ID */
	mappedPresentConfigurations map[PresentIndex]*PresentConfiguration

	Region *maths.Region
}

func NewChristmasTree(
	wide uint,
	long uint,
	presentConfigurations map[PresentIndex]uint,
) *ChristmasTree {

	return &ChristmasTree{
		wide,
		long,
		buildMappedPresentConfigurations(presentConfigurations),
		maths.NewRegion(wide, long, E),
	}
}

func buildMappedPresentConfigurations(
	presentConfigurations map[PresentIndex]uint,
) map[PresentIndex]*PresentConfiguration {

	mappedPresentConfigurations := make(map[PresentIndex]*PresentConfiguration)

	for presentIndex, presentCount := range presentConfigurations {

		mappedPresentConfigurations[presentIndex] = &PresentConfiguration{
			Index: presentIndex,
			Count: presentCount,
		}
	}

	return mappedPresentConfigurations
}

func (ct *ChristmasTree) GetWide() uint {
	return ct.wide
}

func (ct *ChristmasTree) GetLong() uint {
	return ct.long
}

func (ct *ChristmasTree) GetPresentConfiguration(
	presentIndex PresentIndex,
) *PresentConfiguration {
	return ct.mappedPresentConfigurations[presentIndex]
}

func (ct *ChristmasTree) GetPresentsCount() uint {
	presentsCount := uint(0)

	for _, presentConfiguration := range ct.mappedPresentConfigurations {
		presentsCount += presentConfiguration.Count
	}

	return presentsCount
}
