package extensions

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("uint64 slices", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []uint64
			expected []uint64
		}{
			{
				name:     "empty slice",
				input:    []uint64{},
				expected: []uint64{},
			},
			{
				name:     "single element",
				input:    []uint64{1},
				expected: []uint64{1},
			},
			{
				name:     "two elements",
				input:    []uint64{1, 2},
				expected: []uint64{2, 1},
			},
			{
				name:     "odd number of elements",
				input:    []uint64{1, 2, 3, 4, 5},
				expected: []uint64{5, 4, 3, 2, 1},
			},
			{
				name:     "even number of elements",
				input:    []uint64{1, 2, 3, 4},
				expected: []uint64{4, 3, 2, 1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Make a copy to preserve original for error messages
				original := make([]uint64, len(tt.input))
				copy(original, tt.input)

				Reverse(tt.input)

				if !reflect.DeepEqual(tt.input, tt.expected) {
					t.Errorf("Reverse(%v) = %v, expected %v", original, tt.input, tt.expected)
				}
			})
		}
	})

	t.Run("string slices", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []string
			expected []string
		}{
			{
				name:     "empty slice",
				input:    []string{},
				expected: []string{},
			},
			{
				name:     "single word",
				input:    []string{"hello"},
				expected: []string{"hello"},
			},
			{
				name:     "multiple words",
				input:    []string{"a", "b", "c", "d"},
				expected: []string{"d", "c", "b", "a"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				original := make([]string, len(tt.input))
				copy(original, tt.input)

				Reverse(tt.input)

				if !reflect.DeepEqual(tt.input, tt.expected) {
					t.Errorf("Reverse(%v) = %v, expected %v", original, tt.input, tt.expected)
				}
			})
		}
	})

	t.Run("int slices", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{
				name:     "negative and positive numbers",
				input:    []int{-5, -1, 0, 1, 5},
				expected: []int{5, 1, 0, -1, -5},
			},
			{
				name:     "duplicates",
				input:    []int{1, 2, 2, 3, 3, 3},
				expected: []int{3, 3, 3, 2, 2, 1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				original := make([]int, len(tt.input))
				copy(original, tt.input)

				Reverse(tt.input)

				if !reflect.DeepEqual(tt.input, tt.expected) {
					t.Errorf("Reverse(%v) = %v, expected %v", original, tt.input, tt.expected)
				}
			})
		}
	})
}
