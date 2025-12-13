package app

import (
	"day_8/internal/abstractions"
	"fmt"
)

func ConnectJunctionBoxes(
	playground *abstractions.Playground,
	availableCablesCount uint,
	verbose bool,
) *abstractions.Circuits {

	circuits := abstractions.NewCircuits()

	/* Creates circuits with single junction boxes */
	for _, junctionBox := range playground.JunctionBoxes {
		circuits.Add(abstractions.NewCircuit(junctionBox))
	}

	/* Pre-compute all the closest distances up-front */
	distanceRegistry := abstractions.NewDistanceRegistry(playground.JunctionBoxes)

	/* Connects pairs of junction boxes */
	orderedDistances := distanceRegistry.GetOrderedDistances()

	for _, closestJunctionBoxes := range orderedDistances {

		if availableCablesCount == 0 {
			logNoMoreCables(verbose)
			break
		}

		logRemainingCables(availableCablesCount, verbose)
		logShortestConnection(*closestJunctionBoxes, verbose)

		fromCircuit := circuits.Get(closestJunctionBoxes.FromJunctionBox)
		toCircuit := circuits.Get(closestJunctionBoxes.ToJunctionBox)

		if fromCircuit == toCircuit {
			/* Already in the same circuit */
			logSameCircuit(verbose)
			continue
		}

		if toCircuit.HasSingleJunctionBox() {

			/* Include the TO junction box in the FROM circuit */
			logJunctionBoxNotInAnyCircuit(*closestJunctionBoxes.ToJunctionBox, verbose)
			fromCircuit.AddJunctionBox(closestJunctionBoxes.ToJunctionBox)
			circuits.Remove(toCircuit)

		} else if fromCircuit.HasSingleJunctionBox() {

			/* Include the FROM junction box in the TO circuit */
			logJunctionBoxNotInAnyCircuit(*closestJunctionBoxes.FromJunctionBox, verbose)
			toCircuit.AddJunctionBox(closestJunctionBoxes.FromJunctionBox)
			circuits.Remove(fromCircuit)

		} else {
			//	Do nothing?
		}

		availableCablesCount--
	}

	return circuits
}

/* Logging */

func logNoMoreCables(
	verbose bool,
) {
	if verbose {
		fmt.Println("No more available cables to connect junction boxes")
	}
}

func logShortestConnection(
	closestJunctionBoxes abstractions.ClosestJunctionBox,
	verbose bool,
) {
	if verbose {
		fmt.Printf(
			"\tShortest connection between junction boxes: %v %v (distance: %f)\n",
			closestJunctionBoxes.FromJunctionBox.Position,
			closestJunctionBoxes.ToJunctionBox.Position,
			closestJunctionBoxes.Distance,
		)
	}
}

func logRemainingCables(
	availableCablesCount uint,
	verbose bool,
) {
	if verbose {
		fmt.Printf("%d available cables to connect junction boxes\n", availableCablesCount)
	}
}

func logSameCircuit(
	verbose bool,
) {
	if verbose {
		fmt.Printf("\tJunction boxes are in the same circuit\n")
	}
}

func logJunctionBoxNotInAnyCircuit(
	junctionBox abstractions.JunctionBox,
	verbose bool,
) {
	if verbose {
		fmt.Printf(
			"\tJunction box is not in a circuit: %v\n",
			junctionBox.Position,
		)
	}
}
