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
			name:     "trailing spaces in first two numbers",
			cells:    []string{"64 ", "23 ", "314"},
			expected: []uint64{4, 431, 623},
		},
		{
			name:     "leading spaces in first number",
			cells:    []string{" 51", "387", "215"},
			expected: []uint64{175, 581, 32},
		},
		{
			name:     "trailing spaces in second and third numbers",
			cells:    []string{"328", "64 ", "98 "},
			expected: []uint64{8, 248, 369},
		},
		{
			name:     "mixed leading spaces in second and third numbers",
			cells:    []string{"123", " 45", "  6"},
			expected: []uint64{356, 24, 1},
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
