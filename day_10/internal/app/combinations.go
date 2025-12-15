package app

import (
	"day_10/internal/abstractions"
	"fmt"
)

func ActivateMachines(
	factory *abstractions.Factory,
) uint64 {

	totalPresses := uint64(0)

	for machineIndex, machine := range factory.Machines {

		fmt.Printf("Processing machine %d with %d button groups\r", machineIndex+1, machine.GetButtonGroupsCount())

		presses, succeeded := FindShortestCombinations(
			machine.GetButtonGroupsCount(),
			machine.GetButtonGroupsCount(), /* Here we test all the combinations */
			func(buttonGroupsIndexes []int) bool {

				if len(buttonGroupsIndexes) == 3 {
					if buttonGroupsIndexes[0] == 2 && buttonGroupsIndexes[1] == 3 {
						fmt.Printf("")
					}
				}
				/* Resets the machine before testing the combination */
				machine.Reset()

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
		}
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

	var actualCombinations []int
	pressesCount := -1

	var testGroups func(currentButtonGroups []int) (int, bool)

	testGroups = func(currentButtons []int) (int, bool) {

		currentButtonCount := len(currentButtons)

		for buttonIndex := 0; buttonIndex < totalButtonGroupsCount; buttonIndex++ {
			/* Test all the combinations with the current list of buttons  */
			currentButtons[len(currentButtons)-1] = buttonIndex

			if testCombination(currentButtons) {
				if pressesCount == -1 || len(currentButtons) < pressesCount {
					actualCombinations = make([]int, len(currentButtons))
					copy(actualCombinations, currentButtons)
					pressesCount = len(currentButtons)
				}
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

				testGroups(buttonGroupsPrefix)
			}
		}

		return pressesCount, pressesCount != -1
	}

	initialButtonGroups := make([]int, 1)
	return testGroups(initialButtonGroups)
}
