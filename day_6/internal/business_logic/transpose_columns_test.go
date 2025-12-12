package business_logic

import (
	"reflect"
	"testing"
)

func TestTransposeColumns(t *testing.T) {
	tests := []struct {
		name     string
		cells    []string
		expected []uint64
	}{
		{
			name:     "All digits in a single column",
			cells:    []string{"1", "2", "3", "4"},
			expected: []uint64{1234},
		},
		{
			name:     "Different length numbers - 10, 2, 3, 4",
			cells:    []string{"10", "2", "3", "4"},
			expected: []uint64{0, 1234},
		},
		{
			name:     "Different length numbers - 1, 2, 3, 40",
			cells:    []string{"1", "2", "3", "40"},
			expected: []uint64{0, 1234},
		},
		{
			name:     "Different length numbers - 1, 2, 3, 04",
			cells:    []string{"1", "2", "3", "04"},
			expected: []uint64{4, 1230},
		},
		{
			name:     "example case with 9, 15, 84, 942",
			cells:    []string{"9", "15", "84", "942"},
			expected: []uint64{2, 544, 9189},
		},
		{
			name:     "empty second cell",
			cells:    []string{"9", "", "84", "942"},
			expected: []uint64{2, 44, 989},
		},
		{
			name:     "empty third cell",
			cells:    []string{"1", "2", "", "3"},
			expected: []uint64{123},
		},
		{
			name:     "three numbers 64, 23, 314",
			cells:    []string{"64", "23", "314"},
			expected: []uint64{4, 431, 623},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransposeColumns(tt.cells)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("transpose(%v) = %v, expected %v", tt.cells, result, tt.expected)
			}
		})
	}
}
