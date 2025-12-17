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
			a matrix form (variables) and the list of voltages into
			a vector form (solutions)
		*/
		groupsMatrix, voltagesVector := ToMatrix(machine)

		/*	2. I use Gaussian elimination to solve the system of equations */
		abstractions.Reduce(groupsMatrix, voltagesVector)

		fmt.Println("%v %v", groupsMatrix, voltagesVector)

		elapsed := time.Since(startTime)
		fmt.Printf("Processed machine %d with %d button groups: %d pressed needed (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), 0, elapsed)
	}

	fmt.Println()

	return totalPresses
}

func ToMatrix(
	machine *abstractions.Machine,
) (*abstractions.Matrix, *abstractions.Vector) {
	groups := machine.GetButtonGroups()
	voltages := machine.GetVoltages()

	/* Creates the matrix made of the variables */
	groupsMatrix := abstractions.NewMatrix(len(groups), len(voltages))

	for groupIndex, group := range groups {
		for _, button := range group.Buttons {
			groupsMatrix.Set(button.CounterIndex, groupIndex, 1)
		}
	}

	/* Creates the vector made of the result */
	voltagesVector := abstractions.NewVector(len(voltages))

	for voltageIndex, voltage := range voltages {
		voltagesVector.Set(voltageIndex, float64(voltage.GetValue()))
	}

	return groupsMatrix, voltagesVector
}
