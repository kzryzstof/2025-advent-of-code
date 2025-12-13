package app

//
//import (
//	"testing"
//)
//
//func TestNewDistanceRegistry(t *testing.T) {
//	tests := []struct {
//		name                 string
//		junctionBoxes        []*JunctionBox
//		expectedClosestCount int
//		validate             func(t *testing.T, registry *DistanceRegistry)
//	}{
//		{
//			name: "Documented use case",
//			junctionBoxes: []*JunctionBox{
//				{Position: Position{X: 162, Y: 817, Z: 812}},
//				{Position: Position{X: 57, Y: 618, Z: 57}},
//				{Position: Position{X: 906, Y: 360, Z: 560}},
//				{Position: Position{X: 592, Y: 479, Z: 940}},
//				{Position: Position{X: 352, Y: 342, Z: 300}},
//				{Position: Position{X: 466, Y: 668, Z: 158}},
//				{Position: Position{X: 542, Y: 29, Z: 236}},
//				{Position: Position{X: 431, Y: 825, Z: 988}},
//				{Position: Position{X: 739, Y: 650, Z: 466}},
//				{Position: Position{X: 52, Y: 470, Z: 668}},
//				{Position: Position{X: 216, Y: 146, Z: 977}},
//				{Position: Position{X: 819, Y: 987, Z: 18}},
//				{Position: Position{X: 117, Y: 168, Z: 530}},
//				{Position: Position{X: 805, Y: 96, Z: 715}},
//				{Position: Position{X: 346, Y: 949, Z: 466}},
//				{Position: Position{X: 970, Y: 615, Z: 88}},
//				{Position: Position{X: 941, Y: 993, Z: 340}},
//				{Position: Position{X: 862, Y: 61, Z: 35}},
//				{Position: Position{X: 984, Y: 92, Z: 344}},
//				{Position: Position{X: 425, Y: 690, Z: 689}},
//			},
//			expectedClosestCount: 0,
//			validate: func(t *testing.T, registry *DistanceRegistry) {
//				if len(*registry) != 0 {
//					t.Errorf("Expected empty registry, got %d entries", len(*registry))
//				}
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			registry := NewDistanceRegistry(tt.junctionBoxes)
//
//			if registry == nil {
//				t.Fatal("Expected non-nil registry")
//			}
//
//			if len(*registry) != tt.expectedClosestCount {
//				t.Errorf("Expected %d entries in registry, got %d", tt.expectedClosestCount, len(*registry))
//			}
//
//			if tt.validate != nil {
//				tt.validate(t, registry)
//			}
//		})
//	}
//}
//
//// Helper method to get all junction boxes from registry as a slice
//func (d *DistanceRegistry) getJunctionBoxes() []*JunctionBox {
//	keys := make([]*JunctionBox, 0, len(*d))
//	for _, k := range *d {
//		keys = append(keys, k.A)
//	}
//	return keys
//}
