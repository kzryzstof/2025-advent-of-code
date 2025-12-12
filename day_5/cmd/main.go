package main

import (
	"fmt"
	"os"

	"day_5/internal/io"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser */
	reader := initializeReader(inputFile)

	/* Reads all the ingredients */
	ingredients, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Compacts the ranges and counts all the fresh ingredients */
	compactedFreshIngredients := ingredients.Compact()
	freshIngredientsCount := compactedFreshIngredients.Count()

	/* Prints the result */
	fmt.Printf("Number of fresh ingredients: %d\n", freshIngredientsCount)
}

func initializeReader(
	inputFile []string,
) *io.IngredientsReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
