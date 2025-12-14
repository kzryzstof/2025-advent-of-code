package main

import (
	"day_9/internal/app"
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

	fmt.Printf("There are %d red tiles in the movie theater", len(movieTheater.RedTiles))

	biggestRectangle := app.ArrangeTiles(movieTheater)

	/* Prints the result */
	fmt.Printf("The biggest rectangle has an area of %d", biggestRectangle.GetArea())
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
