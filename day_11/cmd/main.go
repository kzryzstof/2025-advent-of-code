package main

import (
	"day_11/internal/abstractions"
	"day_11/internal/io"
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
	devices, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	graph := abstractions.BuildGraph(devices)

	from := "you"
	to := "out"

	pathsCount := graph.CountPaths(from, to)

	elapsed := time.Since(startTime)

	/* Prints the result */
	fmt.Printf("The room has %d devices. Graph has %d path from '%s' to '%s' .\n", len(devices), pathsCount, from, to)

	fmt.Printf("Execution time: %v\n", elapsed)
}

func initializeReader(
	inputFile []string,
) *io.DevicesReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
