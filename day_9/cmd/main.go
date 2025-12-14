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

	/* Reads all the junction boxes from the playground */
	playground, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Connects pairs of junction boxes with only the specific number of cables */
	lastConnectedPair := app.ConnectJunctionBoxes(playground, 1000, true)

	if lastConnectedPair == nil {
		fmt.Println("No solution found")
		os.Exit(1)
	}

	/* Prints the result */
	fmt.Printf("The last connected pair X are: %d and %d (%d)", lastConnectedPair.A.Position.X, lastConnectedPair.B.Position.X, lastConnectedPair.A.Position.X*lastConnectedPair.B.Position.X)
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
