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
	ingredientsParser := initializeParser(inputFile)

	/* Compacts the ranges */
	compactedFreshIngredients := ingredientsParser.Fresh.Compact()

	freshIngredientsCount := compactedFreshIngredients.Count()

	/* Prints the result */
	fmt.Printf("Number of fresh ingredients: %d\n", freshIngredientsCount)
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
