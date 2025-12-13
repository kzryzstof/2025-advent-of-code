package app

import (
	"day_8/internal/abstractions"
	"testing"
)

func TestCreateCircuits(t *testing.T) {
	tests := []struct {
		name                 string
		playground           *abstractions.Playground
		availableCablesCount uint
		expectedCircuitCount int
		expectedSize         int
		validate             func(t *testing.T, circuits *abstractions.Circuits)
	}{
		{
			name: "Documented Use Case",
			playground: &abstractions.Playground{
				JunctionBoxes: []*abstractions.JunctionBox{
					{Position: abstractions.Position{X: 162, Y: 817, Z: 812}},
					{Position: abstractions.Position{X: 57, Y: 618, Z: 57}},
					{Position: abstractions.Position{X: 906, Y: 360, Z: 560}},
					{Position: abstractions.Position{X: 592, Y: 479, Z: 940}},
					{Position: abstractions.Position{X: 352, Y: 342, Z: 300}},
					{Position: abstractions.Position{X: 466, Y: 668, Z: 158}},
					{Position: abstractions.Position{X: 542, Y: 29, Z: 236}},
					{Position: abstractions.Position{X: 431, Y: 825, Z: 988}},
					{Position: abstractions.Position{X: 739, Y: 650, Z: 466}},
					{Position: abstractions.Position{X: 52, Y: 470, Z: 668}},
					{Position: abstractions.Position{X: 216, Y: 146, Z: 977}},
					{Position: abstractions.Position{X: 819, Y: 987, Z: 18}},
					{Position: abstractions.Position{X: 117, Y: 168, Z: 530}},
					{Position: abstractions.Position{X: 805, Y: 96, Z: 715}},
					{Position: abstractions.Position{X: 346, Y: 949, Z: 466}},
					{Position: abstractions.Position{X: 970, Y: 615, Z: 88}},
					{Position: abstractions.Position{X: 941, Y: 993, Z: 340}},
					{Position: abstractions.Position{X: 862, Y: 61, Z: 35}},
					{Position: abstractions.Position{X: 984, Y: 92, Z: 344}},
					{Position: abstractions.Position{X: 425, Y: 690, Z: 689}},
				},
			},
			availableCablesCount: 10,
			expectedSize:         40,
			expectedCircuitCount: 11,
			validate: func(t *testing.T, circuits *abstractions.Circuits) {
				if circuits.Count() != 11 {
					t.Errorf("Expected 11 circuits, got %d", circuits.Count())
				}

				// Verify all 20 junction boxes are accounted for
				totalBoxes := 0
				for _, circuit := range circuits.GetAll() {
					totalBoxes += circuit.Count()
				}
				if totalBoxes != 20 {
					t.Errorf("Expected total of 20 junction boxes across all circuits, got %d", totalBoxes)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			circuits := CreateCircuits(tt.playground, tt.availableCablesCount, true)

			if circuits == nil {
				t.Fatal("Expected non-nil circuits slice")
			}

			if circuits.Count() != tt.expectedCircuitCount {
				t.Errorf("Expected %d circuits, got %d", tt.expectedCircuitCount, circuits.Count())
			}

			// Verify all circuits are the expected size
			biggestCircuits := circuits.GetBiggestCircuits(3)
			actualSize := biggestCircuits[0].Count() * biggestCircuits[1].Count() * biggestCircuits[2].Count()
			if actualSize != tt.expectedSize {
				t.Errorf("Expected size of %d, got %d", tt.expectedSize, actualSize)
			}

			if tt.validate != nil {
				tt.validate(t, circuits)
			}
		})
	}
}
