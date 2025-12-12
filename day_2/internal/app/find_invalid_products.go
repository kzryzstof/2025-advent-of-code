package app

import (
	"day_2/internal/abstractions"
	"fmt"
)

var Verbose bool // set to true to enable debug prints

func FindInvalidProductIds(
	ranges []abstractions.Range,
) uint64 {

	invalidProductSum := uint64(0)

	for _, productRange := range ranges {

		if Verbose {
			fmt.Printf("Processing products range from %s to %s\n", productRange.From.Id, productRange.To.Id)
		}

		invalidProductIds := productRange.FindInvalidProductIds()

		if len(invalidProductIds) == 0 {
			continue
		}

		for _, invalidProductId := range invalidProductIds {
			invalidProductSum += uint64(invalidProductId)
		}
	}

	return invalidProductSum
}
