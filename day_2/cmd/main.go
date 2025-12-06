package main

import (
	"fmt"
	"os"
	"sync"

	//"day_2/internal/abstractions"
	"day_2/internal/parser"
	"day_2/internal/processor"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/*	Initializes the dial */
	//dial := abstractions.Dial{Position: 50}

	/* 	Initializes the parser and processor */
	waitGroup := &sync.WaitGroup{}

	rangesParser := initializeParser(inputFile, waitGroup)
	rangesProcessor := initializeProcessor(rangesParser, waitGroup)

	/* Starts the parser and processor */
	rangesParser.Start()
	rangesProcessor.Start()

	waitGroup.Wait()

	/* Prints the number of times the dial ended up at position 0 */
	//fmt.Printf("Number of the times the dial ended up at position 0: %d\n", dial.GetCount())
}

func initializeParser(
	inputFile []string,
	waitGroup *sync.WaitGroup,
) *parser.RangesParser {
	rangesReader, err := parser.New(inputFile[0], waitGroup)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Ranges parser initialized: %v\n", rangesReader)
	return rangesReader
}

func initializeProcessor(
	rangesParser *parser.RangesParser,
	waitGroup *sync.WaitGroup,
) *processor.RangesProcessor {
	rangesProcessor := processor.New(rangesParser, waitGroup)
	fmt.Printf("Ranges processor initialized: %v\n", rangesParser)
	return rangesProcessor
}
