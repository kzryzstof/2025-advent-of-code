package abstractions

import (
	"slices"
)

type JunctionBoxPair struct {
	A        *JunctionBox
	B        *JunctionBox
	Distance float64
}

func Order(
	unorderedPairs []*JunctionBoxPair,
) []*JunctionBoxPair {

	orderedPairs := make([]*JunctionBoxPair, 0, len(unorderedPairs))

	for _, pair := range unorderedPairs {
		orderedPairs = append(orderedPairs, pair)
	}

	slices.SortFunc(orderedPairs, func(i, j *JunctionBoxPair) int {
		if i.Distance < j.Distance {
			return -1
		}
		if i.Distance > j.Distance {
			return 1
		}
		return 0
	})

	return orderedPairs
}
