package app

import (
	"day_8/internal/abstractions"
	"fmt"
)

func ConnectJunctionBoxes(
	playground *abstractions.Playground,
	availableConnectionsCount uint,
	verbose bool,
) *abstractions.JunctionBoxPair {

	circuits := abstractions.NewCircuits()

	/* Creates circuits with single junction boxes */
	for _, junctionBox := range playground.JunctionBoxes {
		circuits.Add(abstractions.NewCircuit(junctionBox))
	}

	/* Pairs all the junction boxes */
	unorderedPairs := Pair(playground.JunctionBoxes)

	/* Connects pairs of junction boxes as long as there are available cables */
	orderedPairs := abstractions.Order(unorderedPairs)

	logPairs(orderedPairs, 30, verbose)

	var lastPair *abstractions.JunctionBoxPair

	for _, pair := range orderedPairs {

		if circuits.Count() == 1 {
			logAllConnectionsMade(verbose)
			break
		}

		logNewConnection(verbose)
		logPairDistance(*pair, verbose)

		circuitA := circuits.Get(pair.A)
		circuitB := circuits.Get(pair.B)

		if circuitA == circuitB {

			/* Already in the same circuit */
			logSameCircuit(verbose)

		} else if circuitA.IsDisconnected() {

			/* Connects the junction box from circuit A to circuit B */
			logJunctionBoxNotInAnyCircuit(*pair.A, verbose)
			circuits.Connect(circuitA, circuitB)
			logCircuits(circuits, circuitB, verbose)

		} else if circuitB.IsDisconnected() {

			/* Connects the junction box from circuit B to circuit A */
			logJunctionBoxNotInAnyCircuit(*pair.B, verbose)
			circuits.Connect(circuitB, circuitA)
			logCircuits(circuits, circuitA, verbose)

		} else {

			/* Merge the circuits together with one cable */
			circuits.Merge(circuitA, circuitB)
			logCircuits(circuits, circuitB, verbose)

		}

		if circuits.Count() == 1 {
			lastPair = pair
		}
	}

	logDetailsCircuits(circuits, verbose)

	return lastPair
}

/* Logging */

func logAllConnectionsMade(
	verbose bool,
) {
	if verbose {
		fmt.Println("No more junction box available.")
	}
}

func logPairDistance(
	pair abstractions.JunctionBoxPair,
	verbose bool,
) {
	if verbose {
		fmt.Printf(
			"\tShortest distance: %f. Pair: %v %v\n",
			pair.Distance,
			pair.A.Position,
			pair.B.Position,
		)
	}
}

func logNewConnection(
	verbose bool,
) {
	if verbose {
		fmt.Printf("New connection\n")
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

func logPairs(
	pairs []*abstractions.JunctionBoxPair,
	count int,
	verbose bool,
) {
	if verbose {

		fmt.Printf("There are %02d pairs | Showing the first %d\n", len(pairs), count)

		for pairIndex, pair := range pairs {
			if pairIndex >= count {
				break
			}
			fmt.Printf("%02d | Distance: %f | %v - %v\n", pairIndex+1, pair.Distance, pair.A.Position, pair.B.Position)
		}

		fmt.Println()
	}
}
