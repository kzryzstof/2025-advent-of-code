package main

import (
	"day_5/internal/parser"
	"day_5/internal/processor"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser and processor */
	initializeParser(inputFile)
	//sectionsProcessor := initializeProcessor()

	///* Reads each row and analyzes it */
	//accessibleRollsFound := true
	//rowsCount := sectionsParser.GetRowsCount()
	//loopNumber := 1
	//
	//for accessibleRollsFound {
	//	/* Keeps looping until no more accessible roll is found */
	//	accessibleRollsFound = false
	//	fmt.Printf("LOOP %d...\n", loopNumber)
	//	for rowIndex := uint(0); rowIndex < rowsCount; rowIndex++ {
	//
	//		if sectionsProcessor.Analyze(sectionsParser.Section, rowIndex) {
	//			accessibleRollsFound = true
	//		}
	//	}
	//	loopNumber++
	//}
	//
	///* Prints the total number of accessible rolls */
	//fmt.Printf("Number of accessible rolls in the %d row of the department: %d\n", rowsCount, sectionsProcessor.GetTotalAccessibleRolls())
}

func initializeParser(
	inputFile []string,
) *parser.IngredientsParser {
	ingredientsParser, err := parser.NewParser(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Parser initialized: %v\n", ingredientsParser)
	return ingredientsParser
}

func initializeProcessor() *processor.DepartmentProcessor {
	sectionsProcessor := processor.NewProcessor()
	fmt.Printf("Section processor initialized: %v\n", sectionsProcessor)
	return sectionsProcessor
}
