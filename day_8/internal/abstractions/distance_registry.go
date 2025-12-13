package abstractions

import (
	"math"
	"slices"
)

type ClosestJunctionBox struct {
	FromJunctionBox *JunctionBox
	ToJunctionBox   *JunctionBox
	Distance        float64
}
type DistanceRegistry map[*JunctionBox]*ClosestJunctionBox

func NewDistanceRegistry(
	junctionBoxes []*JunctionBox,
) *DistanceRegistry {

	distanceRegistry := make(DistanceRegistry)

	/* Compute the distances between all the junction boxes */

	for fromIndex, fromJunctionBox := range junctionBoxes {

		for toIndex, toJunctionBox := range junctionBoxes {

			if fromIndex == toIndex {
				continue
			}

			distance := fromJunctionBox.Distance(*toJunctionBox)

			/* Only keep the closest junction box */

			value, ok := distanceRegistry[fromJunctionBox]

			if ok && value.Distance <= distance {
				continue
			}

			/* Avoid adding an existing relation in the opposite direction */
			existingRelation := distanceRegistry[toJunctionBox]
			if existingRelation != nil && existingRelation.ToJunctionBox == fromJunctionBox {
				continue
			}

			/* Add or replace the closest junction box */
			distanceRegistry[fromJunctionBox] = &ClosestJunctionBox{
				FromJunctionBox: fromJunctionBox,
				ToJunctionBox:   toJunctionBox,
				Distance:        distance,
			}
		}
	}

	return &distanceRegistry
}

func (d *DistanceRegistry) GetOrderedDistances() []*ClosestJunctionBox {

	closestJunctionBoxes := make([]*ClosestJunctionBox, 0)

	for _, closestJunctionBox := range *d {
		closestJunctionBoxes = append(closestJunctionBoxes, closestJunctionBox)
	}

	slices.SortFunc(closestJunctionBoxes, func(i, j *ClosestJunctionBox) int {
		if i.Distance < j.Distance {
			return -1
		}
		if i.Distance > j.Distance {
			return 1
		}
		return 0
	})

	return closestJunctionBoxes
}

func (d *DistanceRegistry) GetClosestJunctionBoxes() (*JunctionBox, *JunctionBox, float64) {

	var selectedJunctionBoxes *JunctionBox
	closestDistance := math.MaxFloat64

	for _, closestJunctionBox := range *d {
		if closestJunctionBox.Distance < closestDistance {
			selectedJunctionBoxes = closestJunctionBox.ToJunctionBox
			closestDistance = closestJunctionBox.Distance
		}
	}

	return selectedJunctionBoxes, (*d)[selectedJunctionBoxes].ToJunctionBox, closestDistance
}
