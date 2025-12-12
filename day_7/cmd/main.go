package main

import (
	"day_7/internal/app"
	"day_7/internal/io"
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

	// manifold.Draw()

	/* Simulates the beam paths */
	app.Simulate(manifold)

	manifold.Draw()

	/* Prints the result */
	fmt.Printf("There are %d beams\n", len(manifold.Tachyons))
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
