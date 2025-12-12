package main

import (
	"fmt"
	"os"

	"day_2/internal/app"
	"day_2/internal/io"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	rangesParser := initializeReader(inputFile)

	/* Reads all the ranges */
	ranges := rangesParser.Read()

	invalidProductIdsSum := app.FindInvalidProductIds(ranges)

	/* Prints the result */
	fmt.Printf("Sum of all the invalid product IDs found in %d ranges: %d\n", len(ranges), invalidProductIdsSum)
}

func initializeReader(
	inputFile []string,
) *io.RangesReader {

	rangesReader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Ranges reader initialized: %v\n", rangesReader)
	return rangesReader
}
