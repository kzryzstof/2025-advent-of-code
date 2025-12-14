package app

import (
	"day_9/internal/abstractions"
	"testing"
)

func TestArrangeTiles(t *testing.T) {
	tests := map[string]struct {
		redTiles     []*abstractions.Tile
		expectedArea uint64
	}{
		"no tiles": {
			redTiles:     []*abstractions.Tile{},
			expectedArea: 1,
		},
		"single tile": {
			redTiles:     []*abstractions.Tile{{X: 1, Y: 1}},
			expectedArea: 1,
		},
		"two tiles horizontal": {
			redTiles:     []*abstractions.Tile{{X: 0, Y: 0}, {X: 3, Y: 0}},
			expectedArea: 4,
		},
		"two tiles vertical": {
			redTiles:     []*abstractions.Tile{{X: 0, Y: 0}, {X: 0, Y: 4}},
			expectedArea: 5,
		},
		"rectangle from multiple tiles": {
			redTiles: []*abstractions.Tile{
				{X: 0, Y: 0},
				{X: 4, Y: 0},
				{X: 0, Y: 3},
				{X: 4, Y: 3},
			},
			expectedArea: 20,
		},
		"document use case": {
			redTiles: []*abstractions.Tile{
				{X: 7, Y: 1, Color: abstractions.Red},
				{X: 11, Y: 1, Color: abstractions.Red},
				{X: 11, Y: 7, Color: abstractions.Red},
				{X: 9, Y: 7, Color: abstractions.Red},
				{X: 9, Y: 5, Color: abstractions.Red},
				{X: 2, Y: 5, Color: abstractions.Red},
				{X: 2, Y: 3, Color: abstractions.Red},
				{X: 7, Y: 3, Color: abstractions.Red},
			},
			expectedArea: 50,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mt := abstractions.NewMovieTheater(tc.redTiles)

			rect := ArrangeTiles(mt)

			if rect == nil {
				if tc.expectedArea != 0 {
					t.Fatalf("ArrangeTiles() = nil, want area %d", tc.expectedArea)
				}
				return
			}

			if actualArea := rect.GetArea(); actualArea != tc.expectedArea {
				t.Fatalf("ArrangeTiles() area = %d, want %d", actualArea, tc.expectedArea)
			}
		})
	}
}
