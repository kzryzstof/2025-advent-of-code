package main

import (
	"day_4/internal/parser"
	"day_4/internal/processor"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser and processor */
	sectionsParser := initializeParser(inputFile)
	sectionsProcessor := initializeProcessor()

	/* Reads each row and analyzes it */
	section, hasRow := sectionsParser.ReadNextRow()

	for hasRow {
		sectionsProcessor.Analyze(section)
		section, hasRow = sectionsParser.ReadNextRow()
	}

	/* Prints the total number of accessible rolls */
	fmt.Printf("Number of accessible rolls in the %d row of the department: %d\n", sectionsParser.GetRowsCount(), sectionsProcessor.GetTotalAccessibleRolls())
}

func initializeParser(
	inputFile []string,
) *parser.SectionsParser {
	parser, err := parser.NewParser(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Section parser initialized: %v\n", parser)
	return parser
}

func initializeProcessor() *processor.SectionsProcessor {
	sectionsProcessor := processor.NewProcessor()
	fmt.Printf("Section processor initialized: %v\n", sectionsProcessor)
	return sectionsProcessor
}
