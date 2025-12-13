package abstractions

import (
	"testing"
)

func TestNewDistanceRegistry(t *testing.T) {
	tests := []struct {
		name                 string
		junctionBoxes        []*JunctionBox
		expectedClosestCount int
		validate             func(t *testing.T, registry *DistanceRegistry)
	}{
		{
			name:                 "Empty_Junction_Boxes_Returns_Empty_Registry",
			junctionBoxes:        []*JunctionBox{},
			expectedClosestCount: 0,
			validate: func(t *testing.T, registry *DistanceRegistry) {
				if len(*registry) != 0 {
					t.Errorf("Expected empty registry, got %d entries", len(*registry))
				}
			},
		},
		{
			name: "Single_Junction_Box_Returns_Empty_Registry",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
			},
			expectedClosestCount: 0,
			validate: func(t *testing.T, registry *DistanceRegistry) {
				if len(*registry) != 0 {
					t.Errorf("Expected empty registry, got %d entries", len(*registry))
				}
			},
		},
		{
			name: "Two_Junction_Boxes_Each_Has_One_Closest",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 3, Y: 4, Z: 0}},
			},
			expectedClosestCount: 2,
			validate: func(t *testing.T, registry *DistanceRegistry) {
				if len(*registry) != 2 {
					t.Errorf("Expected 2 entries in registry, got %d", len(*registry))
				}

				// Check first junction box has closest distance of 5.0
				jb1 := (*registry)[registry.getJunctionBoxes()[0]]
				if jb1 == nil {
					t.Error("Expected junction box 1 to have a closest neighbor")
				} else if jb1.Distance != 5.0 {
					t.Errorf("Expected distance 5.0, got %f", jb1.Distance)
				}
			},
		},
		{
			name: "Three_Junction_Boxes_Finds_Closest_For_Each",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 1, Y: 0, Z: 0}},
				{Position: Position{X: 10, Y: 0, Z: 0}},
			},
			expectedClosestCount: 3,
			validate: func(t *testing.T, registry *DistanceRegistry) {
				if len(*registry) != 3 {
					t.Errorf("Expected 3 entries in registry, got %d", len(*registry))
				}

				// Each junction box should have its closest neighbor recorded
				// All should have distance of 1.0 (the minimum distance between any pair)
				for i, jb := range registry.getJunctionBoxes() {
					closest := (*registry)[jb]
					if closest == nil {
						t.Errorf("Junction box %d has no closest neighbor", i)
						continue
					}
					// JB at X=0 and X=1 should have distance 1.0 to each other
					// JB at X=10 should have distance 9.0 to X=1
					if jb.Position.X == 10 {
						if closest.Distance != 9.0 {
							t.Errorf("JB at X=10: Expected closest distance 9.0, got %f", closest.Distance)
						}
					} else {
						if closest.Distance != 1.0 {
							t.Errorf("JB at X=%d: Expected closest distance 1.0, got %f", jb.Position.X, closest.Distance)
						}
					}
				}
			},
		},
		{
			name: "Four_Junction_Boxes_In_Square_Pattern",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 1, Y: 0, Z: 0}},
				{Position: Position{X: 0, Y: 1, Z: 0}},
				{Position: Position{X: 1, Y: 1, Z: 0}},
			},
			expectedClosestCount: 4,
			validate: func(t *testing.T, registry *DistanceRegistry) {
				if len(*registry) != 4 {
					t.Errorf("Expected 4 entries in registry, got %d", len(*registry))
				}

				// All junction boxes should have a closest neighbor at distance 1.0
				for _, jb := range registry.getJunctionBoxes() {
					closest := (*registry)[jb]
					if closest == nil {
						t.Errorf("Junction box at %+v has no closest neighbor", jb.Position)
					} else if closest.Distance != 1.0 {
						t.Errorf("Junction box at %+v: Expected closest distance 1.0, got %f", jb.Position, closest.Distance)
					}
				}
			},
		},
		{
			name: "Three_Dimensional_Junction_Boxes",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 1, Y: 1, Z: 1}},
				{Position: Position{X: 2, Y: 2, Z: 2}},
			},
			expectedClosestCount: 3,
			validate: func(t *testing.T, registry *DistanceRegistry) {
				if len(*registry) != 3 {
					t.Errorf("Expected 3 entries in registry, got %d", len(*registry))
				}

				jbs := registry.getJunctionBoxes()

				// Calculate expected distances
				// Distance from (0,0,0) to (1,1,1) = sqrt(3) ≈ 1.732
				// Distance from (1,1,1) to (2,2,2) = sqrt(3) ≈ 1.732
				// Distance from (0,0,0) to (2,2,2) = sqrt(12) ≈ 3.464

				expectedDist := 1.7320508075688772 // sqrt(3)

				closest1 := (*registry)[jbs[0]]
				if closest1 != nil && closest1.Distance != expectedDist {
					t.Errorf("JB1: Expected closest distance %f, got %f", expectedDist, closest1.Distance)
				}

				closest2 := (*registry)[jbs[1]]
				if closest2 != nil && closest2.Distance != expectedDist {
					t.Errorf("JB2: Expected closest distance %f, got %f", expectedDist, closest2.Distance)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := NewDistanceRegistry(tt.junctionBoxes)

			if registry == nil {
				t.Fatal("Expected non-nil registry")
			}

			if len(*registry) != tt.expectedClosestCount {
				t.Errorf("Expected %d entries in registry, got %d", tt.expectedClosestCount, len(*registry))
			}

			if tt.validate != nil {
				tt.validate(t, registry)
			}
		})
	}
}

func TestGetClosestJunctionBoxes(t *testing.T) {
	tests := []struct {
		name             string
		junctionBoxes    []*JunctionBox
		expectedDistance float64
		validate         func(t *testing.T, jb1, jb2 *JunctionBox, distance float64)
	}{
		{
			name: "Two_Junction_Boxes_Returns_The_Pair",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 3, Y: 4, Z: 0}},
			},
			expectedDistance: 5.0,
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				if jb1 == nil || jb2 == nil {
					t.Error("Expected both junction boxes to be non-nil")
				}
				if distance != 5.0 {
					t.Errorf("Expected distance 5.0, got %f", distance)
				}
			},
		},
		{
			name: "Three_Junction_Boxes_Returns_Closest_Pair",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 1, Y: 0, Z: 0}},
				{Position: Position{X: 10, Y: 0, Z: 0}},
			},
			expectedDistance: 1.0,
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				if jb1 == nil || jb2 == nil {
					t.Error("Expected both junction boxes to be non-nil")
				}
				if distance != 1.0 {
					t.Errorf("Expected distance 1.0, got %f", distance)
				}
				// The closest pair should be (0,0,0) and (1,0,0)
				if (jb1.Position.X != 0 && jb1.Position.X != 1) ||
					(jb2.Position.X != 0 && jb2.Position.X != 1) {
					t.Errorf("Expected closest pair to be at X=0 and X=1, got X=%d and X=%d",
						jb1.Position.X, jb2.Position.X)
				}
			},
		},
		{
			name: "Four_Junction_Boxes_Square_Pattern_Returns_Any_Adjacent_Pair",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 1, Y: 0, Z: 0}},
				{Position: Position{X: 0, Y: 1, Z: 0}},
				{Position: Position{X: 1, Y: 1, Z: 0}},
			},
			expectedDistance: 1.0,
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				if jb1 == nil || jb2 == nil {
					t.Error("Expected both junction boxes to be non-nil")
				}
				if distance != 1.0 {
					t.Errorf("Expected distance 1.0, got %f", distance)
				}
			},
		},
		{
			name: "Multiple_Junction_Boxes_With_One_Close_Pair",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 10, Y: 10, Z: 10}},
				{Position: Position{X: 20, Y: 20, Z: 20}},
				{Position: Position{X: 20, Y: 20, Z: 21}}, // Very close to previous
			},
			expectedDistance: 1.0,
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				if jb1 == nil || jb2 == nil {
					t.Error("Expected both junction boxes to be non-nil")
				}
				if distance != 1.0 {
					t.Errorf("Expected distance 1.0, got %f", distance)
				}
				// The closest pair should be the two at (20,20,20) and (20,20,21)
				if !((jb1.Position.X == 20 && jb1.Position.Y == 20) &&
					(jb2.Position.X == 20 && jb2.Position.Y == 20)) {
					t.Errorf("Expected closest pair to be at position (20,20,*), got (%d,%d,%d) and (%d,%d,%d)",
						jb1.Position.X, jb1.Position.Y, jb1.Position.Z,
						jb2.Position.X, jb2.Position.Y, jb2.Position.Z)
				}
			},
		},
		{
			name: "Three_Dimensional_Spacing_Returns_Closest_Pair",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 1, Y: 1, Z: 1}},
				{Position: Position{X: 2, Y: 2, Z: 2}},
				{Position: Position{X: 10, Y: 10, Z: 10}},
			},
			expectedDistance: 1.7320508075688772, // sqrt(3)
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				if jb1 == nil || jb2 == nil {
					t.Error("Expected both junction boxes to be non-nil")
				}
				expectedDist := 1.7320508075688772
				if distance != expectedDist {
					t.Errorf("Expected distance %f, got %f", expectedDist, distance)
				}
			},
		},
		{
			name: "Five_Junction_Boxes_With_Fractional_Distance",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 0, Y: 0, Z: 0}},
				{Position: Position{X: 5, Y: 5, Z: 5}},
				{Position: Position{X: 10, Y: 10, Z: 10}},
				{Position: Position{X: 2, Y: 3, Z: 0}}, // distance sqrt(13) ≈ 3.606 from origin
				{Position: Position{X: 2, Y: 3, Z: 1}}, // distance 1.0 from previous
			},
			expectedDistance: 1.0,
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				if jb1 == nil || jb2 == nil {
					t.Error("Expected both junction boxes to be non-nil")
				}
				if distance != 1.0 {
					t.Errorf("Expected distance 1.0, got %f", distance)
				}
			},
		},

		{
			name: "Documented Use Case",
			junctionBoxes: []*JunctionBox{
				{Position: Position{X: 162, Y: 817, Z: 812}},
				{Position: Position{X: 57, Y: 618, Z: 57}},
				{Position: Position{X: 906, Y: 360, Z: 560}},
				{Position: Position{X: 592, Y: 479, Z: 940}},
				{Position: Position{X: 352, Y: 342, Z: 300}},
				{Position: Position{X: 466, Y: 668, Z: 158}},
				{Position: Position{X: 542, Y: 29, Z: 236}},
				{Position: Position{X: 431, Y: 825, Z: 988}},
				{Position: Position{X: 739, Y: 650, Z: 466}},
				{Position: Position{X: 52, Y: 470, Z: 668}},
				{Position: Position{X: 216, Y: 146, Z: 977}},
				{Position: Position{X: 819, Y: 987, Z: 18}},
				{Position: Position{X: 117, Y: 168, Z: 530}},
				{Position: Position{X: 805, Y: 96, Z: 715}},
				{Position: Position{X: 346, Y: 949, Z: 466}},
				{Position: Position{X: 970, Y: 615, Z: 88}},
				{Position: Position{X: 941, Y: 993, Z: 340}},
				{Position: Position{X: 862, Y: 61, Z: 35}},
				{Position: Position{X: 984, Y: 92, Z: 344}},
				{Position: Position{X: 425, Y: 690, Z: 689}},
			},
			expectedDistance: 316.902193348099,
			validate: func(t *testing.T, jb1, jb2 *JunctionBox, distance float64) {
				expectedFirstPosition := Position{X: 162, Y: 817, Z: 812}
				if jb1.Position != expectedFirstPosition {
					t.Error("Expected junction box 1 to be different")
				}
				expectedSecondPosition := Position{X: 425, Y: 690, Z: 689}
				if jb2.Position != expectedSecondPosition {
					t.Error("Expected junction box 2 to be different")
				}
				// Use tolerance for floating point comparison
				tolerance := 0.000001
				if distance < 316.902193348099-tolerance || distance > 316.902193348099+tolerance {
					t.Errorf("Expected distance around 316.902193348099, got %f", distance)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := NewDistanceRegistry(tt.junctionBoxes)

			if registry == nil {
				t.Fatal("Expected non-nil registry")
			}

			jb1, jb2, distance := registry.GetClosestJunctionBoxes()

			if jb1 == nil || jb2 == nil {
				t.Fatal("Expected both junction boxes to be non-nil")
			}

			// Use tolerance for floating point comparison
			tolerance := 0.000001
			if distance < tt.expectedDistance-tolerance || distance > tt.expectedDistance+tolerance {
				t.Errorf("Expected distance %f, got %f", tt.expectedDistance, distance)
			}

			if tt.validate != nil {
				tt.validate(t, jb1, jb2, distance)
			}
		})
	}
}

// Helper method to get all junction boxes from registry as a slice
func (d *DistanceRegistry) getJunctionBoxes() []*JunctionBox {
	keys := make([]*JunctionBox, 0, len(*d))
	for k := range *d {
		keys = append(keys, k)
	}
	return keys
}
