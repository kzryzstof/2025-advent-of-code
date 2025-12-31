package abstractions

import (
	"day_12/internal/maths"
)

type ChristmasTree struct {
	Index ChristmasTreeIndex

	/* Size available under the tree */
	Dimension maths.Dimension

	/* Lists all the present configurations mapped by their ID */
	mappedPresentConfigurations map[PresentIndex]*PresentConfiguration

	Region *maths.Region
}

func NewChristmasTree(
	index ChristmasTreeIndex,
	wide uint,
	long uint,
	presentConfigurations map[PresentIndex]uint,
) *ChristmasTree {

	return &ChristmasTree{
		index,
		maths.Dimension{
			Wide: wide,
			Long: long,
		},
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
