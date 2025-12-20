package app

import (
	"day_10/internal/abstractions"
	"fmt"
	"time"
)

const (
	Verbose = false
)

func ActivateMachines(
	factory *abstractions.Factory,
) uint64 {

	totalPresses := float64(0)

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
		rref := abstractions.ToReducedRowEchelonForm(augmentedMatrix, Verbose)

		total := float64(0)

		for _, presses := range rref.Solve(Verbose) {
			total += presses
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Processed machine %d with %d button groups: %f pressed needed (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), total, elapsed)

		totalPresses += total
	}

	fmt.Println()

	return uint64(totalPresses)
}
