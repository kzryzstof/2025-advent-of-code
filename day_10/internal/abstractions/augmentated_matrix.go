package abstractions

type AugmentedMatrix struct {
	Matrix *Matrix
	Vector *Vector
}

func ToAugmentedMatrix(
	machine *Machine,
) *AugmentedMatrix {
	groups := machine.GetButtonGroups()
	voltages := machine.GetVoltages()

	/* Creates the matrix made of the variables */
	groupsMatrix := NewMatrix(
		len(voltages),
		len(groups),
	)

	for groupIndex, group := range groups {
		for _, button := range group.Buttons {
			groupsMatrix.Set(button.CounterIndex, groupIndex, 1)
		}
	}

	/* Creates the vector made of the result */
	voltagesVector := NewVector(len(voltages))

	for voltageIndex, voltage := range voltages {
		voltagesVector.Set(voltageIndex, float64(voltage.GetValue()))
	}

	return &AugmentedMatrix{
		Matrix: groupsMatrix,
		Vector: voltagesVector,
	}
}
