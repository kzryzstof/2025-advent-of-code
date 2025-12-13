package main

import (
	"day_8/internal/app"
	"day_8/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads the manifold from the file */
	manifold, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Simulates the beam paths */
	app.Simulate(manifold, true)

	/* Prints the result */
	fmt.Printf("The beam has created %d timelines\n", manifold.CountTimelines())
}

func initializeReader(
	inputFile []string,
) *io.ManifoldReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Parser initialized: %v\n", reader)
	return reader
}
