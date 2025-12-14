package abstractions

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name             string
		position         Position
		otherPosition    Position
		expectedDistance float64
	}{
		{
			name:             "Same_Position_Returns_Zero_Distance",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 0, Z: 0},
			expectedDistance: 0.0,
		},
		{
			name:             "Unit_Distance_On_X_Axis",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 1, Y: 0, Z: 0},
			expectedDistance: 1.0,
		},
		{
			name:             "Unit_Distance_On_Y_Axis",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 1, Z: 0},
			expectedDistance: 1.0,
		},
		{
			name:             "Unit_Distance_On_Z_Axis",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 0, Z: 1},
			expectedDistance: 1.0,
		},
		{
			name:             "Distance_On_X_Axis_Multiple_Units",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 5, Y: 0, Z: 0},
			expectedDistance: 5.0,
		},
		{
			name:             "Distance_On_Y_Axis_Multiple_Units",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 10, Z: 0},
			expectedDistance: 10.0,
		},
		{
			name:             "Distance_On_Z_Axis_Multiple_Units",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 0, Z: 7},
			expectedDistance: 7.0,
		},
		{
			name:             "3_4_5_Right_Triangle_On_XY_Plane",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 3, Y: 4, Z: 0},
			expectedDistance: 5.0,
		},
		{
			name:             "5_12_13_Right_Triangle_On_XY_Plane",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 5, Y: 12, Z: 0},
			expectedDistance: 13.0,
		},
		{
			name:             "Diagonal_In_2D_Square_XY",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 1, Y: 1, Z: 0},
			expectedDistance: math.Sqrt(2),
		},
		{
			name:             "Diagonal_In_2D_Square_XZ",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 1, Y: 0, Z: 1},
			expectedDistance: math.Sqrt(2),
		},
		{
			name:             "Diagonal_In_2D_Square_YZ",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 1, Z: 1},
			expectedDistance: math.Sqrt(2),
		},
		{
			name:             "Diagonal_In_3D_Unit_Cube",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 1, Y: 1, Z: 1},
			expectedDistance: math.Sqrt(3),
		},
		{
			name:             "Diagonal_In_3D_Cube_Size_2",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 2, Y: 2, Z: 2},
			expectedDistance: math.Sqrt(12),
		},
		{
			name:             "Distance_Is_Symmetric_Case_1",
			position:         Position{X: 1, Y: 2, Z: 3},
			otherPosition:    Position{X: 4, Y: 6, Z: 8},
			expectedDistance: math.Sqrt(9 + 16 + 25), // sqrt(50)
		},
		{
			name:             "Distance_Is_Symmetric_Case_2",
			position:         Position{X: 4, Y: 6, Z: 8},
			otherPosition:    Position{X: 1, Y: 2, Z: 3},
			expectedDistance: math.Sqrt(9 + 16 + 25), // sqrt(50)
		},
		{
			name:             "Large_Coordinates_Case_1",
			position:         Position{X: 100, Y: 200, Z: 300},
			otherPosition:    Position{X: 100, Y: 200, Z: 300},
			expectedDistance: 0.0,
		},
		{
			name:             "Large_Coordinates_Case_2",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 100, Y: 100, Z: 100},
			expectedDistance: math.Sqrt(30000),
		},
		{
			name:             "Real_World_Example_From_Test_Data",
			position:         Position{X: 162, Y: 817, Z: 812},
			otherPosition:    Position{X: 57, Y: 618, Z: 57},
			expectedDistance: math.Sqrt(105*105 + 199*199 + 755*755),
		},
		{
			name:             "Adjacent_Positions_On_X_Axis",
			position:         Position{X: 5, Y: 10, Z: 15},
			otherPosition:    Position{X: 6, Y: 10, Z: 15},
			expectedDistance: 1.0,
		},
		{
			name:             "Adjacent_Positions_On_Y_Axis",
			position:         Position{X: 5, Y: 10, Z: 15},
			otherPosition:    Position{X: 5, Y: 11, Z: 15},
			expectedDistance: 1.0,
		},
		{
			name:             "Adjacent_Positions_On_Z_Axis",
			position:         Position{X: 5, Y: 10, Z: 15},
			otherPosition:    Position{X: 5, Y: 10, Z: 16},
			expectedDistance: 1.0,
		},
		{
			name:             "Complex_3D_Distance",
			position:         Position{X: 10, Y: 20, Z: 30},
			otherPosition:    Position{X: 13, Y: 24, Z: 33},
			expectedDistance: math.Sqrt(9 + 16 + 9), // sqrt(34)
		},
		{
			name:             "Maximum_X_Coordinate_Difference",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 1000, Y: 0, Z: 0},
			expectedDistance: 1000.0,
		},
		{
			name:             "Maximum_Y_Coordinate_Difference",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 1000, Z: 0},
			expectedDistance: 1000.0,
		},
		{
			name:             "Maximum_Z_Coordinate_Difference",
			position:         Position{X: 0, Y: 0, Z: 0},
			otherPosition:    Position{X: 0, Y: 0, Z: 1000},
			expectedDistance: 1000.0,
		},
		{
			name:             "All_Coordinates_Different",
			position:         Position{X: 10, Y: 20, Z: 30},
			otherPosition:    Position{X: 40, Y: 50, Z: 60},
			expectedDistance: math.Sqrt(900 + 900 + 900), // sqrt(2700)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distance := tt.position.Distance(tt.otherPosition)

			// Use tolerance for floating point comparison
			tolerance := 0.000001
			if math.Abs(distance-tt.expectedDistance) > tolerance {
				t.Errorf("Expected distance %f, got %f", tt.expectedDistance, distance)
			}

			// Verify distance is non-negative
			if distance < 0 {
				t.Errorf("Distance should never be negative, got %f", distance)
			}

			// Verify symmetry: distance(A, B) == distance(B, A)
			reverseDistance := tt.otherPosition.Distance(tt.position)
			if math.Abs(distance-reverseDistance) > tolerance {
				t.Errorf("Distance should be symmetric: forward=%f, reverse=%f", distance, reverseDistance)
			}
		})
	}
}
