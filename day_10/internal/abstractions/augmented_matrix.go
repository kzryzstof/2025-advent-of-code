package abstractions

type AugmentedMatrix struct {
	Matrix *Matrix
}

func ToAugmentedMatrix(
	machine *Machine,
) *AugmentedMatrix {
	groups := machine.GetButtonGroups()
	voltages := machine.GetVoltages()

	groupsMatrix := NewMatrix(
		len(voltages),
		len(groups)+1,
	)

	/* Injects the button groups as coefficient of the matrix */
	for groupIndex, group := range groups {
		for _, button := range group.Buttons {
			groupsMatrix.Set(button.CounterIndex, groupIndex, 1)
		}
	}

	/* Injects the voltages as constant terms of the matrix on the last column */
	lastColumn := len(groups)
	for voltageIndex, voltage := range voltages {
		groupsMatrix.Set(voltageIndex, lastColumn, float64(voltage.GetValue()))
	}

	/* Now we have the augmented matrix */
	return &AugmentedMatrix{
		Matrix: groupsMatrix,
	}
}
