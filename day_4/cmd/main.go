package main

import (
	"fmt"
	"os"
	"sync"

	"day_4/internal/parser"
	"day_4/internal/processor"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser and processor */
	waitGroup := &sync.WaitGroup{}

	sectionsParser := initializeParser(inputFile, waitGroup)
	sectionsProcessor := initializeProcessor(sectionsParser, waitGroup)

	/* Starts the parser and processor */
	sectionsParser.Start()
	sectionsProcessor.Start()

	waitGroup.Wait()

	/* Prints the total number of accessible rolls */
	fmt.Printf("Number of accessible rolls in the %d row of the department: %d\n", sectionsParser.GetRowsCount(), sectionsProcessor.GetTotalAccessibleRolls())
}

func initializeParser(
	inputFile []string,
	waitGroup *sync.WaitGroup,
) *parser.SectionsParser {
	parser, err := parser.NewParser(inputFile[0], waitGroup)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Ranges parser initialized: %v\n", parser)
	return parser
}

func initializeProcessor(
	sectionsParser *parser.SectionsParser,
	waitGroup *sync.WaitGroup,
) *processor.SectionsProcessor {
	sectionsProcessor := processor.NewProcessor(sectionsParser, waitGroup)
	fmt.Printf("Ranges processor initialized: %v\n", sectionsParser)
	return sectionsProcessor
}
