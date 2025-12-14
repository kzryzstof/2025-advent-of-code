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
			expectedArea: 24,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mt := NewMovieTheater(tc.redTiles)

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
