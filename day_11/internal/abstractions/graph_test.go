package abstractions

import "testing"

func TestGraph_BuildGraph(t *testing.T) {
	tests := []struct {
		name               string
		devices            []*Device
		fromNode           string
		toNode             string
		requiredNodes      []string
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
			fromNode:           "you",
			toNode:             "out",
			requiredNodes:      []string{},
			expectedPathsCount: 5,
		},
		{
			name: "1.2-eight_paths_case",
			devices: []*Device{
				NewDevice("svr", []string{"aaa", "bbb"}),
				NewDevice("aaa", []string{"fft"}),
				NewDevice("bbb", []string{"tty"}),
				NewDevice("fft", []string{"ccc"}),
				NewDevice("tty", []string{"ccc"}),
				NewDevice("ccc", []string{"ddd", "eee"}),
				NewDevice("ddd", []string{"hub"}),
				NewDevice("eee", []string{"dac"}),
				NewDevice("hub", []string{"fff"}),
				NewDevice("dac", []string{"fff"}),
				NewDevice("fff", []string{"ggg", "hhh"}),
				NewDevice("ggg", []string{"out"}),
				NewDevice("hhh", []string{"out"}),
			},
			fromNode:           "svr",
			toNode:             "out",
			requiredNodes:      []string{"fft", "dac"},
			expectedPathsCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := BuildGraph(tt.devices, tt.requiredNodes)

			actualPathsCount := graph.CountPaths(
				tt.fromNode,
				tt.toNode,
				tt.requiredNodes,
			)

			if actualPathsCount != tt.expectedPathsCount {
				t.Errorf("Expected %d paths, got %d", tt.expectedPathsCount, actualPathsCount)
			}
		})
	}
}
