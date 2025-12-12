package main

import (
	"day_6/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Read all the problems */
	problems, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Computes the total */
	total, err := problems.ComputeTotal()

	if err != nil {
		fmt.Printf("Error computing total: %v\n", err)
		os.Exit(1)
	}

	/* Prints the result */
	fmt.Printf("Total = %d\n", total)
}

func initializeReader(
	inputFile []string,
) *io.ProblemsReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Parser initialized: %v\n", reader)
	return reader
}
