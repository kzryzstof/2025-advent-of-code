package app

import (
	"day_7/internal/abstractions"
	"strings"
	"testing"
)

// helper to build a manifold from a single string where spaces separate rows
// and '.' represents Empty, '^' a Splitter, 'S' a StartingPoint.
func buildManifoldFromSingleLine(input string) *abstractions.Manifold {
	// each space separates rows
	rows := strings.Split(input, " ")

	locations := make([][]string, len(rows))
	var tachyons []*abstractions.Tachyon

	for r, row := range rows {
		locations[r] = make([]string, len(row))
		for c, ch := range row {
			cell := string(ch)
			locations[r][c] = cell
			if cell == abstractions.StartingPoint {
				// create a moving tachyon at this position
				pos := abstractions.Position{RowIndex: r, ColIndex: c}
				t := &abstractions.Tachyon{Position: pos}
				t.Start()
				tachyons = append(tachyons, t)
			}
		}
	}

	return &abstractions.Manifold{
		Locations: locations,
		Tachyons:  tachyons,
	}
}

func TestSimulate_WithGivenManifoldLayout(t *testing.T) {
	tests := []struct {
		name           string
		layout         string
		expectedSplits uint
	}{
		{
			name:           "given splitter pattern without explicit start",
			layout:         ".......S....... ............... .......^....... ............... ......^.^...... ............... .....^.^.^..... ............... ....^.^...^.... ............... ...^.^...^.^... ............... ..^...^.....^.. ............... .^.^.^.^.^...^. ...............",
			expectedSplits: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manifold := buildManifoldFromSingleLine(tt.layout)

			manifold.Draw()

			Simulate(manifold)

			if tt.expectedSplits != manifold.GetSplitsCount() {
				t.Fatalf("expected at least one starting tachyon, got 0")
			}

			// Run the simulation; the primary assertion is that it terminates
			// without panic or infinite loop.
			Simulate(manifold)
		})
	}
}
