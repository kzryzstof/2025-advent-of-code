package main

import (
	"day_9/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads all the red tiles from the movie theater */
	movieTheater, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Prints the result */
	fmt.Printf("There are %d red tiles in the movie theater", len(movieTheater.RedTiles))
}

func initializeReader(
	inputFile []string,
) *io.RedTilesReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
