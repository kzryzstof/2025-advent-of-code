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

	/* Pre-computes all the closest distances up-front */
	distanceRegistry := abstractions.NewDistanceRegistry(playground.JunctionBoxes)

	/* Connects pairs of junction boxes as long as there are available cables */
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
			circuits.AddTo(toCircuit, fromCircuit)
			logCircuits(circuits, fromCircuit, verbose)

		} else if fromCircuit.HasSingleJunctionBox() {

			/* Include the FROM junction box in the TO circuit */
			logJunctionBoxNotInAnyCircuit(*closestJunctionBoxes.FromJunctionBox, verbose)
			circuits.AddTo(fromCircuit, toCircuit)
			logCircuits(circuits, toCircuit, verbose)

		} else {
			/* Merge the circuits together with one cable */
			circuits.Merge(fromCircuit, toCircuit)
			logCircuits(circuits, toCircuit, verbose)
		}

		availableCablesCount--
	}

	logDetailsCircuits(circuits, verbose)

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
			"\tShortest distance: %f. Connecting junction boxes: %v %v\n",
			closestJunctionBoxes.Distance,
			closestJunctionBoxes.FromJunctionBox.Position,
			closestJunctionBoxes.ToJunctionBox.Position,
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

func logJunctionBoxAlreadyInCircuits(
	fromJunctionBox abstractions.JunctionBox,
	toJunctionBox abstractions.JunctionBox,
	verbose bool,
) {
	if verbose {
		fmt.Printf(
			"\tJunction boxes are already each in a circuit: %v %v. Skipping\n",
			fromJunctionBox.Position,
			toJunctionBox.Position,
		)
	}
}

func logCircuits(
	circuits *abstractions.Circuits,
	circuit *abstractions.Circuit,
	verbose bool,
) {
	if verbose {
		fmt.Printf("\tCircuit has now %d junction boxes\n", circuit.Count())
		fmt.Printf("\tThere are now %d circuits in total\n", circuits.Count())
	}
}

func logDetailsCircuits(
	circuits *abstractions.Circuits,
	verbose bool,
) {
	if verbose {
		for circuitIndex, circuit := range circuits.GetAll() {
			fmt.Printf("\tCircuit %d has now %d junction boxes\n", circuitIndex+1, circuit.Count())
		}
	}
}
