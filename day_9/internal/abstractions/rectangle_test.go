package abstractions

import "testing"

func TestNewRectangleSurface(t *testing.T) {
	tests := map[string]struct {
		a, b         *Tile
		expectedArea uint64
	}{
		"documented_use_case_1": {
			a:            &Tile{X: 2, Y: 5},
			b:            &Tile{X: 9, Y: 7},
			expectedArea: 24,
		},
		"documented_use_case_2": {
			a:            &Tile{X: 7, Y: 1},
			b:            &Tile{X: 11, Y: 7},
			expectedArea: 35,
		},
		"documented_use_case_3": {
			a:            &Tile{X: 7, Y: 3},
			b:            &Tile{X: 2, Y: 3},
			expectedArea: 6,
		},
		"documented_use_case_4": {
			a:            &Tile{X: 2, Y: 5},
			b:            &Tile{X: 11, Y: 1},
			expectedArea: 50,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rect := NewRectangle(tc.a, tc.b)

			if actualArea := rect.GetArea(); actualArea != tc.expectedArea {
				t.Fatalf("GetArea() = %d, want %d", actualArea, tc.expectedArea)
			}
		})
	}
}
