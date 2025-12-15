package main

import (
	"day_10/internal/app"
	"day_10/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads all the machines from the factory */
	factory, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	pressesCount := app.ActivateMachines(factory)

	/* Prints the result */
	fmt.Printf("The factory has %d machines. All of them have been activated with %d presses", len(factory.Machines), pressesCount)
}

func initializeReader(
	inputFile []string,
) *io.FactoryReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
