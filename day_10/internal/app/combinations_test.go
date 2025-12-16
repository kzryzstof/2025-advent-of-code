package app

import (
	"fmt"
	"testing"
)

func TestComputeCombinations(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "ListIsEmpty_EmptyListReturned",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindShortestCombinations(11, 11, func(ints []int) bool {
				fmt.Printf("Testing groups: [")
				for _, i := range ints {
					fmt.Printf("%d ", i)
				}
				fmt.Println("]")
				return false
			})
		})
	}
}
