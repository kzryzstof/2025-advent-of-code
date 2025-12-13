package main

import (
	"day_8/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads all the junction boxes from the playground */
	playground, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Prints the result */
	fmt.Printf("The playground has created %v junction boxes\n", len(playground.JunctionBoxes))
}

func initializeReader(
	inputFile []string,
) *io.PlaygroundReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
