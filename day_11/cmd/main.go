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

	fmt.Printf("%d devices read\n", len(devices))

	requiredNodes := []string{"svr", "fft", "dac", "out"}

	graph := abstractions.BuildGraph(devices, requiredNodes)
	fmt.Printf("Graph built\n")

	from := "svr"
	to := "out"

	pathsCount := graph.CountPaths(from, to, requiredNodes)

	elapsed := time.Since(startTime)

	/* Prints the result */
	fmt.Printf("Graph has %d path from '%s' to '%s'\n", pathsCount, from, to)

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
