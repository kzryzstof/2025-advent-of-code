package main

import (
	"day_5/internal/parser"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser and processor */
	ingredientsParser := initializeParser(inputFile)

	/* Founds out all the fresh ingredients */
	freshIngredientsCount := 0

	for _, ingredientId := range ingredientsParser.Available.Ids {
		if ingredientsParser.Fresh.IsFresh(ingredientId) {
			freshIngredientsCount++
		}
	}

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
