package app

import (
	"day_8/internal/abstractions"
	"fmt"
	"slices"
)

func CreateCircuits(
	playground *abstractions.Playground,
	availableCablesCount uint,
	verbose bool,
) []*abstractions.Circuit {

	circuits := make([]*abstractions.Circuit, 0, len(playground.JunctionBoxes))

	/* Creates circuits with single junction boxes */
	for _, junctionBox := range playground.JunctionBoxes {
		circuits = append(circuits, abstractions.NewCircuit(junctionBox))
	}

	/* Pre-compute all the closest distances up-front */
	distanceRegistry := abstractions.NewDistanceRegistry(playground.JunctionBoxes)

	/* Connects pairs of junction boxes */
	orderedDistances := distanceRegistry.GetOrderedDistances()

	for _, closestJunctionBoxes := range orderedDistances {

		if availableCablesCount == 0 {
			if verbose {
				fmt.Println("No more available cables to connect junction boxes")
			}
			break
		}

		if verbose {
			fmt.Printf("%d available cables to connect junction boxes\n", availableCablesCount)
		}

		if verbose {
			fmt.Printf(
				"\tShortest connection between junction boxes: %v %v (distance: %f)\n",
				closestJunctionBoxes.FromJunctionBox.Position,
				closestJunctionBoxes.ToJunctionBox.Position,
				closestJunctionBoxes.Distance,
			)
		}

		fromCircuit := abstractions.GetCircuit(circuits, closestJunctionBoxes.FromJunctionBox)
		toCircuit := abstractions.GetCircuit(circuits, closestJunctionBoxes.ToJunctionBox)

		if fromCircuit == toCircuit {
			/* Already in the same circuit */
			if verbose {
				fmt.Printf("\tJunction boxes are in the same circuit\n")
			}
			continue
		}

		if toCircuit.HasSingleJunctionBox() {

			if verbose {
				fmt.Printf(
					"\tJunction box is not in a circuit: %v\n",
					closestJunctionBoxes.ToJunctionBox.Position,
				)
			}

			fromCircuit.AddJunctionBox(closestJunctionBoxes.ToJunctionBox)

			circuits = slices.DeleteFunc(circuits, func(c *abstractions.Circuit) bool {
				return c == toCircuit
			})
		} else if fromCircuit.HasSingleJunctionBox() {

			if verbose {
				fmt.Printf(
					"\tJunction box is not in a circuit: %v\n",
					closestJunctionBoxes.FromJunctionBox.Position,
				)
			}

			toCircuit.AddJunctionBox(closestJunctionBoxes.FromJunctionBox)

			circuits = slices.DeleteFunc(circuits, func(c *abstractions.Circuit) bool {
				return c == fromCircuit
			})
		} else {

			/*
				for _, fromJunctionBox := range fromCircuit.GetJunctionBoxes() {
					toCircuit.AddJunctionBox(fromJunctionBox)
				}

				circuits = slices.DeleteFunc(circuits, func(c *abstractions.Circuit) bool {
					return c == fromCircuit
				})
			*/
		}

		availableCablesCount--
	}

	return circuits
}
