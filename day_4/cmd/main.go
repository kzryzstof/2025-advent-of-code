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

	banksParser := initializeParser(inputFile, waitGroup)
	banksProcessor := initializeProcessor(banksParser, waitGroup)

	/* Starts the parser and processor */
	banksParser.Start()
	banksProcessor.Start()

	waitGroup.Wait()

	/* Prints the total number of consisting of the sum of all the product IDs */
	fmt.Printf("Sum of all the highest voltage from the %d banks: %d\n", banksParser.GetBanksCount(), banksProcessor.GetTotalVoltage())
}

func initializeParser(
	inputFile []string,
	waitGroup *sync.WaitGroup,
) *parser.BanksParser {
	rangesReader, err := parser.NewParser(inputFile[0], waitGroup)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Ranges parser initialized: %v\n", rangesReader)
	return rangesReader
}

func initializeProcessor(
	banksParser *parser.BanksParser,
	waitGroup *sync.WaitGroup,
) *processor.BanksProcessor {
	banksProcessor := processor.NewProcessor(banksParser, waitGroup)
	fmt.Printf("Ranges processor initialized: %v\n", banksParser)
	return banksProcessor
}
