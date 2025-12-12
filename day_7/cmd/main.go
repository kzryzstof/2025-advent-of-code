package main

import (
	"day_7/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	manifold, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Prints the result */
	fmt.Printf("Total = %d\n", manifold)
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
