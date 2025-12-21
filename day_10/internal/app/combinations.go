package app

import (
	"day_10/internal/abstractions"
	"day_10/internal/algorithms"
	"fmt"
	"os"
	"time"
)

const (
	Verbose = false
)

func ActivateMachines(
	factory *abstractions.Factory,
) int64 {

	totalPresses := int64(0)

	for machineIndex, machine := range factory.Machines {

		startTime := time.Now()
		fmt.Printf("Processing machine %d with %d button groups\r", machineIndex+1, machine.GetButtonGroupsCount())

		/*
			Transforms the list of button groups into
			an augmented matrix (variables) and the list of voltages into
			a vector form (constants)
		*/
		augmentedMatrix := abstractions.ToAugmentedMatrix(machine)

		/*
			Uses an HNF to solve the system of equations with integers
		*/
		hnf := algorithms.ToHermiteNormalForm(augmentedMatrix, Verbose)

		solution := hnf.Solve(Verbose)

		/*
			Although the solution is supposed to be valid,
			testing the solution provided helpful feedback
			to catch mistakes in the algorithm
		*/
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

		total := int64(0)

		for _, presses := range solution.GetValues() {
			total += presses
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Processed machine %d with %d button groups: %d presses needed (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), total, elapsed)

		totalPresses += total
	}

	fmt.Println()

	return totalPresses
}
