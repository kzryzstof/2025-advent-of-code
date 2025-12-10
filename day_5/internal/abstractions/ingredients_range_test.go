package abstractions

import (
	"reflect"
	"testing"
)

func TestFreshIngredients_Compact(t *testing.T) {
	type testCase struct {
		name     string
		input    FreshIngredients
		expected []IngredientRange
	}

	tests := []testCase{
		{
			name: "empty ranges",
			input: FreshIngredients{
				Ranges: []IngredientRange{},
			},
			expected: []IngredientRange{},
		},
		{
			name: "single range",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 5},
				},
			},
			expected: []IngredientRange{
				{From: 1, To: 5},
			},
		},
		{
			name: "non overlapping ranges stay unchanged",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 3},
					{From: 5, To: 7},
				},
			},
			expected: []IngredientRange{
				{From: 1, To: 3},
				{From: 5, To: 7},
			},
		},
		{
			name: "overlapping ranges are compacted",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 3},
					{From: 2, To: 5},
				},
			},
			// adjust this expectation to whatever behavior Compact
			// is supposed to implement once finished
			expected: []IngredientRange{
				{From: 1, To: 5},
			},
		},
		{
			name: "touching ranges (no gap)",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 3},
					{From: 4, To: 6},
				},
			},
			// if Compact should merge touching ranges, expect [1,6],
			// otherwise expect two ranges; adjust accordingly
			expected: []IngredientRange{
				{From: 1, To: 6},
			},
		},
		{
			name: "touching ranges (no gap)",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 4, To: 6},
					{From: 1, To: 3},
				},
			},
			// if Compact should merge touching ranges, expect [1,6],
			// otherwise expect two ranges; adjust accordingly
			expected: []IngredientRange{
				{From: 1, To: 6},
			},
		},
		{
			name: "inside ranges",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 6},
					{From: 2, To: 5},
				},
			},
			// if Compact should merge touching ranges, expect [1,6],
			// otherwise expect two ranges; adjust accordingly
			expected: []IngredientRange{
				{From: 1, To: 6},
			},
		},
		{
			name: "documented use case",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 3, To: 5},
					{From: 10, To: 14},
					{From: 16, To: 20},
					{From: 12, To: 18},
				},
			},
			// if Compact should merge touching ranges, expect [1,6],
			// otherwise expect two ranges; adjust accordingly
			expected: []IngredientRange{
				{From: 3, To: 5},
				{From: 10, To: 20},
			},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotPtr := tc.input.Compact()
			if gotPtr == nil {
				t.Fatalf("Compact() returned nil slice pointer")
			}
			got := *gotPtr

			if !reflect.DeepEqual(got.Ranges, tc.expected) {
				t.Fatalf("Compact() = %#v, expected %#v", got, tc.expected)
			}
		})
	}
}

func TestFreshIngredients_Count(t *testing.T) {
	type testCase struct {
		name     string
		input    FreshIngredients
		expected uint64
	}

	tests := []testCase{
		{
			name: "empty ranges",
			input: FreshIngredients{
				Ranges: []IngredientRange{},
			},
			expected: 0,
		},
		{
			name: "single range",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 5},
				},
			},
			expected: 5,
		},
		{
			name: "non overlapping ranges stay unchanged",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 3},
					{From: 5, To: 7},
				},
			},
			expected: 6,
		},
		{
			name: "overlapping ranges are compacted",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 3},
					{From: 2, To: 5},
				},
			},
			expected: 7,
		},
		{
			name: "touching ranges (no gap)",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 3},
					{From: 4, To: 6},
				},
			},
			expected: 6,
		},
		{
			name: "touching ranges (no gap)",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 4, To: 6},
					{From: 1, To: 3},
				},
			},
			expected: 6,
		},
		{
			name: "inside ranges",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 1, To: 6},
					{From: 2, To: 5},
				},
			},
			expected: 10,
		},
		{
			name: "documented use case",
			input: FreshIngredients{
				Ranges: []IngredientRange{
					{From: 3, To: 5},
					{From: 10, To: 20},
				},
			},
			expected: 14,
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.input.Count()

			if got != tc.expected {
				t.Fatalf("Count() = %#v, expected %#v", got, tc.expected)
			}
		})
	}
}
