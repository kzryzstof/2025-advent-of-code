package main

import (
	"day_12/internal/io"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	elapsed := time.Since(startTime)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads all the presents and trees from the cavern */
	cavern, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Found %d presents\n", cavern.GetPresentsCount())
	fmt.Printf("Found %d Christmas trees\n", cavern.GetChristmasTreesCount())

	/* Prints the result */
	fmt.Printf("Execution time: %v\n", elapsed)
}

func initializeReader(
	inputFile []string,
) *io.CavernReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
