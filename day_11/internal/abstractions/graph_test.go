package abstractions

import "testing"

func TestGraph_BuildGraph(t *testing.T) {
	tests := []struct {
		name               string
		devices            []*Device
		expectedPathsCount uint
	}{
		{
			name: "1.1-documented_use_case",
			devices: []*Device{
				NewDevice("you", []string{"bbb", "ccc"}),
				NewDevice("bbb", []string{"ddd", "eee"}),
				NewDevice("ccc", []string{"ddd", "eee", "fff"}),
				NewDevice("ddd", []string{"ggg"}),
				NewDevice("eee", []string{"out"}),
				NewDevice("fff", []string{"out"}),
				NewDevice("ggg", []string{"out"}),
			},
			expectedPathsCount: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := BuildGraph(tt.devices)

			actualPathsCount := graph.CountPaths("you", "out")

			if actualPathsCount != tt.expectedPathsCount {
				t.Errorf("Expected %d paths, got %d", tt.expectedPathsCount, actualPathsCount)
			}
		})
	}
}
