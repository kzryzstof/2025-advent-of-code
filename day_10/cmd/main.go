package main

import (
	"day_10/internal/app"
	"day_10/internal/io"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

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

	elapsed := time.Since(startTime)

	//	20195 too high!!
	//	20172 too high!!
	//	20117: too high!!
	//	20011: not the right value
	//	20030

	/* Prints the result */
	fmt.Printf("The factory has %d machines. All of them have been activated with %d presses\n", len(factory.Machines), pressesCount)

	fmt.Printf("Execution time: %v\n", elapsed)
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
