package main

import (
	"day_12/internal/abstractions"
	"day_12/internal/algorithms"
	"day_12/internal/io"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	inputFile := os.Args[1:]

	cavern, err := getCavern(inputFile)

	if err != nil {
		os.Exit(1)
	}

	packPresentsUnderChristmasTrees(cavern)

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func packPresentsUnderChristmasTrees(
	cavern *abstractions.Cavern,
) {

	failed := algorithms.PackAll(cavern, false)

	fmt.Println()

	if failed > 0 {
		fmt.Printf("Unable to place all presents under all the Christmas trees: only %d trees had enough space (out of %d)\n\n", cavern.GetChristmasTreesCount()-failed, cavern.GetChristmasTreesCount())
		return
	}

	fmt.Printf("All the presents have been placed under the Christmas trees\n\n")
}

func getCavern(
	inputFile []string,
) (*abstractions.Cavern, error) {

	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Unable to initialize the reader: %v\n", err)
		return nil, err
	}

	cavern, err := reader.Read()

	if err != nil {
		fmt.Printf("Unable to read the cavern information: %v\n", err)
		return nil, err
	}

	fmt.Printf("Found %d presents\n", cavern.GetPresentsCount())
	fmt.Printf("Found %d Christmas trees\n", cavern.GetChristmasTreesCount())

	return cavern, nil
}
