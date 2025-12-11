package main

import (
	"day_6/internal/parser"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser */
	problemsParser := initializeParser(inputFile)

	total := problemsParser.Problems.ComputeTotal()

	/* Prints the result */
	fmt.Printf("Total = %d\n", total)
}

func initializeParser(
	inputFile []string,
) *parser.ProblemsParser {
	problemsParser, err := parser.NewParser(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Parser initialized: %v\n", problemsParser)
	return problemsParser
}
