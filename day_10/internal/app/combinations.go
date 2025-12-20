package app

import (
	"day_10/internal/abstractions"
	"fmt"
	"os"
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

		solution := rref.Solve(Verbose)

		for _, variable := range solution.Get() {
			for counter := 0; counter < int(variable.Value); counter++ {
				machine.PressGroup(int(variable.Number - 1))
			}
		}

		if !machine.IsVoltageValid() {
			fmt.Print("\nVOLTAGE ERROR!\n\n")
			machine.PrintVoltages()
			os.Exit(1)
		}

		total := float64(0)

		for _, presses := range solution.GetValues() {
			total += presses
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Processed machine %d with %d button groups: %d presses needed (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), int(total), elapsed)

		totalPresses += total
	}

	fmt.Println()

	return uint64(totalPresses)
}
