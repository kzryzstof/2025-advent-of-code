package extensions

import (
	"reflect"
	"testing"
)

func TestToDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single digit",
			input:    "5",
			expected: []string{"5"},
		},
		{
			name:     "two digits",
			input:    "42",
			expected: []string{"4", "2"},
		},
		{
			name:     "three digits",
			input:    "123",
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "number with zero",
			input:    "105",
			expected: []string{"1", "0", "5"},
		},
		{
			name:     "number starting with zero",
			input:    "007",
			expected: []string{"0", "0", "7"},
		},
		{
			name:     "large number",
			input:    "9876543210",
			expected: []string{"9", "8", "7", "6", "5", "4", "3", "2", "1", "0"},
		},
		{
			name:     "all same digits",
			input:    "7777",
			expected: []string{"7", "7", "7", "7"},
		},
		{
			name:     "all zeros",
			input:    "0000",
			expected: []string{"0", "0", "0", "0"},
		},
		{
			name:     "very large number",
			input:    "12345678901234567890",
			expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test: returns []string of digit characters
			result := ToDigits(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToDigits(%q) = %v, expected %v", tt.input, result, tt.expected)
			}

			// Additional validation: check that the length matches
			if len(result) != len(tt.expected) {
				t.Errorf("ToDigits(%q) returned slice of length %d, expected %d",
					tt.input, len(result), len(tt.expected))
			}

			// Validate that most significant digit is first (leftmost digit in the string)
			if len(tt.input) > 0 {
				firstChar := string(tt.input[0])
				if result[0] != firstChar {
					t.Errorf("ToDigits(%q) first element = %q, expected most significant digit %q",
						tt.input, result[0], firstChar)
				}
			}
		})
	}
}
