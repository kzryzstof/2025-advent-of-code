package app

import (
	"day_10/internal/abstractions"
	"fmt"
	"os"
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

		presses, succeeded := FindShortestCombinations(
			machine.GetButtonGroupsCount(),
			machine.GetButtonGroupsCount(), /* Here we test all the combinations */
			func(buttonGroupsIndexes []int) bool {

				fmt.Printf("Processing machine %d with %d button groups [ ", machineIndex+1, machine.GetButtonGroupsCount())

				for _, buttonGroupIndex := range buttonGroupsIndexes {
					fmt.Printf("%d ", buttonGroupIndex)
				}

				fmt.Printf("]\r")

				/* Resets the machine before testing the combination */
				machine.CloseLights()

				/* Presses all the button groups in the combination */
				for _, buttonGroupIndex := range buttonGroupsIndexes {
					machine.PressGroup(buttonGroupIndex)
				}

				/* Tests if the machine is activated */
				return machine.IsActivated()
			},
		)

		if succeeded {
			totalPresses += uint64(presses)
		} else {
			elapsed := time.Since(startTime)
			fmt.Printf("Failed to activate machine %d with %d button groups (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), elapsed)
			os.Exit(1)
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Processed machine %d with %d button groups: %d pressed needed (%v)\n", machineIndex+1, machine.GetButtonGroupsCount(), presses, elapsed)
	}

	fmt.Println()

	return totalPresses
}

func FindShortestCombinations(
	/* Indicates the maximum number of button groups are in the combination to test */
	maximumCombinationLength int,
	/* Indicates the maximum number of button groups to test */
	totalButtonGroupsCount int,
	/* Function to call to test the machine after pressing of all the button groups in the collection */
	testCombination func([]int) bool,
) (int, bool) {

	/* The recursive function is not optimal */

	var testGroups func(currentButtonGroups []int, currentNumberToTest int) (int, bool)

	testGroups = func(currentButtons []int, currentNumberToTest int) (int, bool) {

		currentButtonCount := len(currentButtons)

		canTest := currentButtonCount == currentNumberToTest

		for buttonIndex := 0; buttonIndex < totalButtonGroupsCount; buttonIndex++ {

			/* Test all the combinations with the current list of buttons */
			currentButtons[len(currentButtons)-1] = buttonIndex

			if canTest && testCombination(currentButtons) {
				return currentButtonCount, true
			}
		}

		if currentButtonCount < maximumCombinationLength {

			/* Creates a new list of buttons for the next iteration that will have one more button added */
			buttonGroupsPrefix := make([]int, currentButtonCount+1)

			/* Makes sure to include the current list */
			copy(buttonGroupsPrefix, currentButtons)

			/* Loops to test all the combinations with one more button group */
			for buttonIndex := 0; buttonIndex < totalButtonGroupsCount; buttonIndex++ {

				buttonGroupsPrefix[currentButtonCount-1] = buttonIndex

				pressedCount, succeeded := testGroups(buttonGroupsPrefix, currentNumberToTest)

				if succeeded {
					return pressedCount, true
				}
			}
		}

		return -1, false
	}

	count := 1

	for count <= maximumCombinationLength {

		initialButtonGroups := make([]int, 1)

		pressesCount, succeeded := testGroups(initialButtonGroups, count)

		if succeeded {
			return pressesCount, true
		}

		count++
	}

	return -1, false
}
