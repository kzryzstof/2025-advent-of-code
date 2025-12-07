package abstractions

import "testing"

func TestRange_FindInvalidProductIds(t *testing.T) {
	tests := []struct {
		name                      string
		from                      string
		to                        string
		expectedInvalidProductIds []int64
	}{
		{
			name:                      "Range from 11 to 22",
			from:                      "11",
			to:                        "22",
			expectedInvalidProductIds: []int64{11, 22},
		},
		{
			name:                      "Range from 95 to 115",
			from:                      "95",
			to:                        "115",
			expectedInvalidProductIds: []int64{99, 111},
		},
		{
			name:                      "Range from 998 to 1012",
			from:                      "998",
			to:                        "1012",
			expectedInvalidProductIds: []int64{999, 1010},
		},
		{
			name:                      "Range from 1188511880 to 1188511890",
			from:                      "1188511880",
			to:                        "1188511890",
			expectedInvalidProductIds: []int64{1188511885},
		},
		{
			name:                      "Range from 222220 to 222224",
			from:                      "222220",
			to:                        "222224",
			expectedInvalidProductIds: []int64{222222},
		},
		{
			name:                      "Range from 1698522 to 1698528",
			from:                      "1698522",
			to:                        "1698528",
			expectedInvalidProductIds: []int64{},
		},
		{
			name:                      "Range from 446443 to 446449",
			from:                      "446443",
			to:                        "446449",
			expectedInvalidProductIds: []int64{446446},
		},
		{
			name:                      "Range from 38593856 to 38593862",
			from:                      "38593856",
			to:                        "38593862",
			expectedInvalidProductIds: []int64{38593859},
		},
		{
			name:                      "Range from 565653 to 565659",
			from:                      "565653",
			to:                        "565659",
			expectedInvalidProductIds: []int64{565656},
		},
		{
			name:                      "Range from 2121212118 to 2121212124",
			from:                      "2121212118",
			to:                        "2121212124",
			expectedInvalidProductIds: []int64{2121212121},
		},
		{
			name:                      "Range from 824824821 to 824824827",
			from:                      "824824821",
			to:                        "824824827",
			expectedInvalidProductIds: []int64{824824824},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromProd, err := NewProduct(tt.from)
			if err != nil {
				t.Fatalf("failed to create from product: %v", err)
			}
			toProd, err := NewProduct(tt.to)
			if err != nil {
				t.Fatalf("failed to create to product: %v", err)
			}

			r := Range{From: *fromProd, To: *toProd}
			actual := r.FindInvalidProductIds()

			if len(actual) != len(tt.expectedInvalidProductIds) {
				t.Fatalf("expected %d invalid ids, got %d", len(tt.expectedInvalidProductIds), len(actual))
			}

			for i, productId := range actual {
				if productId != tt.expectedInvalidProductIds[i] {
					t.Errorf("at index %d: expected id %q, got %q", i, tt.expectedInvalidProductIds[i], productId)
				}
			}
		})
	}
}
