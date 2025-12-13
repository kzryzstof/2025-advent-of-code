package app

import (
	"day_8/internal/abstractions"
)

func Pair(
	junctionBoxes []*abstractions.JunctionBox,
) []*abstractions.JunctionBoxPair {

	unorderedPairs := make([]*abstractions.JunctionBoxPair, 0)

	/* Creates pairs of junction boxes, with the computed distances between the two */

	for fromIndex, fromJunctionBox := range junctionBoxes {

		for toIndex := fromIndex + 1; toIndex < len(junctionBoxes); toIndex++ {

			toJunctionBox := junctionBoxes[toIndex]

			distance := fromJunctionBox.MeasureDistance(*toJunctionBox)

			unorderedPairs = append(
				unorderedPairs,
				&abstractions.JunctionBoxPair{
					A:        fromJunctionBox,
					B:        toJunctionBox,
					Distance: distance,
				},
			)
		}
	}

	return unorderedPairs
}
