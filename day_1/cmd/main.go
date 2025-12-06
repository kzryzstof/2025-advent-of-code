package main

import (
	"fmt"
	"os"
	"sync"

	"day_1/internal/abstractions"
	"day_1/internal/parser"
	"day_1/internal/processor"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/*	Initializes the dial */
	dial := abstractions.Dial{Position: 50}

	/* 	Initializes the parser and processor */
	waitGroup := &sync.WaitGroup{}

	instructionsParser := initializeParser(inputFile, waitGroup)
	instructionsProcessor := initializeProcessor(instructionsParser, waitGroup)

	/* Starts the parser and processor */
	instructionsParser.Start()
	instructionsProcessor.Start(&dial)

	waitGroup.Wait()

	/* Prints the number of times the dial ended up at position 0 */
	fmt.Printf("Number of the times the dial ended up at position 0: %d\n", dial.GetCount())
}

func initializeParser(
	inputFile []string,
	waitGroup *sync.WaitGroup,
) *parser.InstructionsParser {
	instructionsReader, err := parser.New(inputFile[0], waitGroup)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Instructions parser initialized: %v\n", instructionsReader)
	return instructionsReader
}

func initializeProcessor(
	parser *parser.InstructionsParser,
	waitGroup *sync.WaitGroup,
) *processor.InstructionsProcessor {
	instructionsProcessor := processor.New(parser, waitGroup)
	fmt.Printf("Instructions processor initialized: %v\n", instructionsProcessor)
	return instructionsProcessor
}
