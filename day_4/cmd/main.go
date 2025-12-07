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

	/*	Initializes the dial */
	//dial := abstractions.Dial{Position: 50}

	/* 	Initializes the parser and processor */
	waitGroup := &sync.WaitGroup{}

	sectionsParser := initializeParser(inputFile, waitGroup)
	banksProcessor := initializeProcessor(sectionsParser, waitGroup)

	/* Starts the parser and processor */
	sectionsParser.Start()
	banksProcessor.Start()

	waitGroup.Wait()

	/* Prints the total number of consisting of the sum of all the product IDs */
	fmt.Printf("Sum of all the highest voltage from the %d rows: %d\n", sectionsParser.GetRowsCount(), banksProcessor.GetTotalVoltage())
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
	banksParser *parser.SectionsParser,
	waitGroup *sync.WaitGroup,
) *processor.BanksProcessor {
	banksProcessor := processor.NewProcessor(banksParser, waitGroup)
	fmt.Printf("Ranges processor initialized: %v\n", banksParser)
	return banksProcessor
}
