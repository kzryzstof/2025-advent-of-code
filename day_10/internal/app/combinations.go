package app

import (
	"day_10/internal/abstractions"
	"fmt"
	"time"
)

const (
	// TilesValidationConcurrency on macOS M1 Pro, 10 seems to be a nice sweet spot
	TestConcurrency = 10
)

func ActivateMachines(
	factory *abstractions.Factory,
) uint64 {

	totalPresses := uint64(0)

	for machineIndex, machine := range factory.Machines {

		startTime := time.Now()
		fmt.Printf("Processing machine %d with %d button groups\r", machineIndex+1, machine.GetButtonGroupsCount())

		/* Integer Linear Programming Fun Time! */

		/*
			1. I have to transform the list of button groups into
			an augmented matrix (variables) and the list of voltages into
			a vector form (solutions)
		*/
		augmentedMatrix := abstractions.ToAugmentedMatrix(machine)

		/*	2. I use Gaussian elimination to solve the system of equations */
		rref := abstractions.ToReducedRowEchelonForm(augmentedMatrix, false)

		total := 0

		for _, presses := range rref.Solve(false) {
			total += int(presses)
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Processed machine %d with %d button groups: %d pressed needed (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), total, elapsed)

		totalPresses += uint64(total)
	}

	fmt.Println()

	return totalPresses
}
