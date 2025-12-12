package main

import (
	"fmt"
	"os"

	"day_1/internal/abstractions"
	"day_1/internal/io"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/*	Initializes the dial */
	dial := abstractions.Dial{Position: 50}

	/* 	Initializes the parser */
	instructionsParser := newParser(inputFile)

	/* Reads all the instructions at once */
	instructions := instructionsParser.Read()

	for _, rotation := range instructions.Rotations {
		dial.Rotate(rotation)
	}

	/* Prints the results */
	fmt.Printf("Number of the times the dial passed by position 0: %d\n", dial.GetCount())
}

func newParser(
	inputFile []string,
) *io.InstructionsReader {
	instructionsReader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Error parsing input file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parser initialized: %v\n", instructionsReader)
	return instructionsReader
}
