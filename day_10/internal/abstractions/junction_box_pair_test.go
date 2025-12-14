package abstractions

import (
	"testing"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		name                 string
		unorderedPairs       []*JunctionBoxPair
		expectedOrderedPairs []*JunctionBoxPair
	}{
		{
			name:                 "ListIsEmpty_EmptyListReturned",
			unorderedPairs:       []*JunctionBoxPair{},
			expectedOrderedPairs: []*JunctionBoxPair{},
		},
		{
			name: "ListHasOnePair_PairReturned",
			unorderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
			},
			expectedOrderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
			},
		},
		{
			name: "ListHasTwoOrderedPairs_OrderedReturnedUnchanged",
			unorderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 5, Y: 0, Z: 0}},
					Distance: 5.0,
				},
			},
			expectedOrderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 5, Y: 0, Z: 0}},
					Distance: 5.0,
				},
			},
		},
		{
			name: "ListHasTwoUnorderedPairs_OrderedReturned",
			unorderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 5, Y: 0, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
			},
			expectedOrderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 5, Y: 0, Z: 0}},
					Distance: 5.0,
				},
			},
		},
		{
			name: "ListHasMultipleUnorderedPairs_OrderedReturned",
			unorderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 10, Y: 0, Z: 0}},
					Distance: 10.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 3, Y: 4, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 3, Y: 0, Z: 0}},
					Distance: 3.0,
				},
			},
			expectedOrderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 3, Y: 0, Z: 0}},
					Distance: 3.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 3, Y: 4, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 10, Y: 0, Z: 0}},
					Distance: 10.0,
				},
			},
		},
		{
			name: "ListHasPairWithSameDistance_OrderUnchangedReturned",
			unorderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 3, Y: 4, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 0, Y: 5, Z: 0}},
					Distance: 5.0,
				},
			},
			expectedOrderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 3, Y: 4, Z: 0}},
					Distance: 5.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 0, Y: 5, Z: 0}},
					Distance: 5.0,
				},
			},
		},
		{
			name: "ListHasUnorderedPairsWithFractionalDistances_OrderListReturned",
			unorderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 1, Z: 1}},
					Distance: 1.7320508075688772, // sqrt(3)
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 1, Z: 0}},
					Distance: 1.4142135623730951, // sqrt(2)
				},
			},
			expectedOrderedPairs: []*JunctionBoxPair{
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 0, Z: 0}},
					Distance: 1.0,
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 1, Z: 0}},
					Distance: 1.4142135623730951, // sqrt(2)
				},
				{
					A:        &JunctionBox{Position: Position{X: 0, Y: 0, Z: 0}},
					B:        &JunctionBox{Position: Position{X: 1, Y: 1, Z: 1}},
					Distance: 1.7320508075688772, // sqrt(3)
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Store original for comparison
			expectedOrderedLen := len(tt.expectedOrderedPairs)

			actualOrderedPairs := Order(tt.unorderedPairs)

			if actualOrderedPairs == nil {
				t.Fatal("Expected non-nil ordered pairs slice")
			}

			if len(actualOrderedPairs) != len(tt.expectedOrderedPairs) {
				t.Errorf("Expected %d pairs, got %d", expectedOrderedLen, len(actualOrderedPairs))
			}

			for i := 0; i < expectedOrderedLen; i++ {
				if actualOrderedPairs[i].A.Position != tt.expectedOrderedPairs[i].A.Position {
					t.Errorf("%d | A is incorrect | %v != %v",
						i, actualOrderedPairs[i].A, tt.expectedOrderedPairs[i].A)
				}

				if actualOrderedPairs[i].B.Position != tt.expectedOrderedPairs[i].B.Position {
					t.Errorf("%d | B is incorrect | %v != %v",
						i, actualOrderedPairs[i].B, tt.expectedOrderedPairs[i].B)
				}

				if actualOrderedPairs[i].Distance != tt.expectedOrderedPairs[i].Distance {
					t.Errorf("%d | Distance is incorrect | %v != %v",
						i, actualOrderedPairs[i].Distance, tt.expectedOrderedPairs[i].Distance)
				}
			}
		})
	}
}
