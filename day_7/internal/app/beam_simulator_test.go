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
				t := abstractions.NewTachyon(
					pos,
				)
				tachyons = append(tachyons, t)
			}
		}
	}

	return abstractions.NewManifold(
		locations,
		tachyons,
	)
}

func TestSimulate_WithGivenManifoldLayout(t *testing.T) {
	tests := []struct {
		name              string
		layout            string
		expectedTimelines uint64
	}{
		{
			name: "1 level of splits",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"...............",
			expectedTimelines: uint64(2),
		},
		{
			name: "2 levels of splits",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^.^...... " +
				"............... ",
			expectedTimelines: uint64(4),
		},
		{
			name: "3 levels of splits (1)",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^........ " +
				"............... " +
				".....^......... " +
				"............... ",
			expectedTimelines: uint64(4),
		},
		{
			name: "3 levels of splits (2)",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^........ " +
				"............... " +
				".....^.^....... " +
				"............... ",
			expectedTimelines: uint64(5),
		},
		{
			name: "3 levels of splits (3)",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^........ " +
				"............... " +
				".....^......... " +
				"........^...... " +
				"...............",
			expectedTimelines: uint64(5),
		},
		{
			name: "3 levels of splits",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^.^...... " +
				"............... " +
				".....^.^.^..... " +
				"............... ",
			expectedTimelines: uint64(8),
		},
		{
			name: "4 levels of splits",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^.^...... " +
				"............... " +
				".....^.^.^..... " +
				"............... " +
				"....^.^...^.... " +
				"............... ",
			expectedTimelines: uint64(13),
		},
		{
			name: "documented used case",
			layout: "" +
				".......S....... " +
				"............... " +
				".......^....... " +
				"............... " +
				"......^.^...... " +
				"............... " +
				".....^.^.^..... " +
				"............... " +
				"....^.^...^.... " +
				"............... " +
				"...^.^...^.^... " +
				"............... " +
				"..^...^.....^.. " +
				"............... " +
				".^.^.^.^.^...^. " +
				"...............",
			expectedTimelines: uint64(40),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manifold := buildManifoldFromSingleLine(tt.layout)

			manifold.Draw()

			Simulate(manifold, false)

			manifold.Draw()

			actualTimelines := manifold.CountTimelines()

			if tt.expectedTimelines != actualTimelines {
				t.Fatalf("expected %d timelines, got %d", tt.expectedTimelines, actualTimelines)
			}
		})
	}
}
