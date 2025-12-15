package app

import "day_10/internal/abstractions"

func ActivateMachines(
	factory *abstractions.Factory,
) uint64 {

	totalPresses := uint64(0)

	for _, machine := range factory.Machines {
		presses, succeeded := FindShortestCombinations(
			machine.GetButtonGroupsCount(),
			0, /* <-- Not sure about this parameter!! */
			func(buttonGroupsIndexes []int) bool {

				/* Resets the machine before testing the combination */
				machine.Reset()

				/* Presses all the button groups in the combination */
				for _, buttonGroupIndex := range buttonGroupsIndexes {
					machine.PressGroup(buttonGroupIndex)
				}

				/* Tests if the machine is activated */
				return machine.IsOn()
			},
		)

		if succeeded {
			totalPresses += uint64(presses)
		}
	}

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

	var testGroups func(currentButtonGroups []int) (int, bool)

	testGroups = func(currentButtons []int) (int, bool) {

		currentButtonCount := len(currentButtons)

		for buttonIndex := 1; buttonIndex <= totalButtonGroupsCount; buttonIndex++ {
			/* Test all the combinations with the current list of buttons  */
			currentButtons[len(currentButtons)-1] = buttonIndex

			if testCombination(currentButtons) {
				return len(currentButtons), true
			}
		}

		if currentButtonCount < maximumCombinationLength {

			/* Creates a new list of buttons for the next iteration that will have one more button added */
			buttonGroupsPrefix := make([]int, currentButtonCount+1)

			/* Makes sure to include the current list */
			copy(buttonGroupsPrefix, currentButtons)

			/* Loops to test all the combinations with one more button group */
			for buttonIndex := 1; buttonIndex <= totalButtonGroupsCount; buttonIndex++ {

				buttonGroupsPrefix[currentButtonCount-1] = buttonIndex

				pressesCount, succeeded := testGroups(buttonGroupsPrefix)

				if succeeded {
					return pressesCount, true
				}
			}
		}

		return -1, false
	}

	initialButtonGroups := make([]int, 1)
	return testGroups(initialButtonGroups)
}
