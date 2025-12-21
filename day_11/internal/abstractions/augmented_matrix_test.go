package abstractions

import "testing"

func TestToAugmentedMatrix(t *testing.T) {
	tests := []struct {
		name         string
		buttonGroups []*ButtonGroup
		voltages     []*Voltage
		/* Each row of the matrix represents a counter and each cell a button group */
		/* 1 = button group affects the specific counter. 0 = button group doesn't affect the specific counter */
		expectedMatrix [][]int64
	}{
		{
			name: "1.1-documented_use_case",
			buttonGroups: []*ButtonGroup{
				{Buttons: []*Button{NewButton(3)}},
				{Buttons: []*Button{NewButton(1), NewButton(3)}},
				{Buttons: []*Button{NewButton(2)}},
				{Buttons: []*Button{NewButton(2), NewButton(3)}},
				{Buttons: []*Button{NewButton(0), NewButton(2)}},
				{Buttons: []*Button{NewButton(0), NewButton(1)}},
			},
			voltages: []*Voltage{
				NewVoltage(3),
				NewVoltage(5),
				NewVoltage(4),
				NewVoltage(7),
			},
			expectedMatrix: [][]int64{
				{0, 0, 0, 0, 1, 1, 3},
				{0, 1, 0, 0, 0, 1, 5},
				{0, 0, 1, 1, 1, 0, 4},
				{1, 1, 0, 1, 0, 0, 7},
			},
		},
		{
			name: "1.2-documented_use_case",
			buttonGroups: []*ButtonGroup{
				{Buttons: []*Button{NewButton(0), NewButton(2), NewButton(3), NewButton(4)}},
				{Buttons: []*Button{NewButton(2), NewButton(3)}},
				{Buttons: []*Button{NewButton(0), NewButton(4)}},
				{Buttons: []*Button{NewButton(0), NewButton(1), NewButton(2)}},
				{Buttons: []*Button{NewButton(1), NewButton(2), NewButton(3), NewButton(4)}},
			},
			voltages: []*Voltage{
				NewVoltage(7),
				NewVoltage(5),
				NewVoltage(12),
				NewVoltage(7),
				NewVoltage(2),
			},
			expectedMatrix: [][]int64{
				{1, 0, 1, 1, 0, 7},
				{0, 0, 0, 1, 1, 5},
				{1, 1, 0, 1, 1, 12},
				{1, 1, 0, 0, 1, 7},
				{1, 0, 1, 0, 1, 2},
			},
		},
		{
			name: "1.3-documented_use_case",
			buttonGroups: []*ButtonGroup{
				{Buttons: []*Button{NewButton(0), NewButton(1), NewButton(2), NewButton(3), NewButton(4)}},
				{Buttons: []*Button{NewButton(0), NewButton(3), NewButton(4)}},
				{Buttons: []*Button{NewButton(0), NewButton(1), NewButton(2), NewButton(4), NewButton(5)}},
				{Buttons: []*Button{NewButton(1), NewButton(2)}},
			},
			voltages: []*Voltage{
				NewVoltage(10),
				NewVoltage(11),
				NewVoltage(11),
				NewVoltage(5),
				NewVoltage(10),
				NewVoltage(5),
			},
			expectedMatrix: [][]int64{
				{1, 1, 1, 0, 10},
				{1, 0, 1, 1, 11},
				{1, 0, 1, 1, 11},
				{1, 1, 0, 0, 5},
				{1, 1, 1, 0, 10},
				{0, 0, 1, 0, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machine := NewMachine(tt.buttonGroups, tt.voltages)
			result := ToAugmentedMatrix(machine)

			Print(result.Matrix)

			// Verify matrix dimensions
			if result.Matrix.Rows() != len(tt.expectedMatrix) {
				t.Errorf("Matrix rows mismatch: got %d, want %d", result.Matrix.Rows(), len(tt.expectedMatrix))
			}
			if result.Matrix.Cols() != len(tt.expectedMatrix[0]) {
				t.Errorf("Matrix cols mismatch: got %d, want %d", result.Matrix.Cols(), len(tt.expectedMatrix[0]))
			}

			// Verify matrix values
			for row := 0; row < result.Matrix.Rows(); row++ {
				for col := 0; col < result.Matrix.Cols(); col++ {
					got := result.Matrix.Get(row, col)
					want := tt.expectedMatrix[row][col]
					if got != want {
						t.Errorf("Matrix[%d][%d] = %d, want %d", row, col, got, want)
					}
				}
			}
		})
	}
}
