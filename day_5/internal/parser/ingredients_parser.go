package parser

import (
	"bufio"
	"day_5/internal/abstractions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	freshIngredientsSection = iota
	availableIngredientsSection
)

type IngredientsParser struct {
	Fresh     *abstractions.FreshIngredients
	Available *abstractions.AvailableIngredients
}

func NewParser(
	filePath string,
) (*IngredientsParser, error) {

	freshIngredients, availableIngredients, err := readIngredients(filePath)

	if err != nil {
		return nil, err
	}

	return &IngredientsParser{
		freshIngredients,
		availableIngredients,
	}, nil
}

func readIngredients(
	filePath string,
) (*abstractions.FreshIngredients, *abstractions.AvailableIngredients, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, nil, err
	}

	freshIngredients := abstractions.FreshIngredients{
		Ranges: make([]abstractions.IngredientRange, 0, 1000),
	}

	availableIngredients := abstractions.AvailableIngredients{
		Ids: make([]abstractions.IngredientId, 0, 2500),
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(inputFile)

	processingSection := freshIngredientsSection

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Trim(line, " \n\t") == "" {
			processingSection = availableIngredientsSection
			continue
		}

		switch processingSection {
		case freshIngredientsSection:
			{
				rangesSlice := strings.Split(line, "-")

				if len(rangesSlice) != 2 {
					return nil, nil, fmt.Errorf("invalid range: %s", line)
				}

				from, err := strconv.ParseUint(rangesSlice[0], 10, 64)

				if err != nil {
					return nil, nil, err
				}

				to, err := strconv.ParseUint(rangesSlice[1], 10, 64)

				if err != nil {
					return nil, nil, err
				}

				if from > to {
					return nil, nil, fmt.Errorf("invalid range: %s", line)
				}

				freshIngredients.Ranges = append(freshIngredients.Ranges, abstractions.IngredientRange{
					From: abstractions.IngredientId(from),
					To:   abstractions.IngredientId(to),
				})
				break
			}
		case availableIngredientsSection:
			{
				ingredientId, err := strconv.ParseUint(line, 10, 64)

				if err != nil {
					return nil, nil, err
				}

				availableIngredients.Ids = append(availableIngredients.Ids, abstractions.IngredientId(ingredientId))
				break
			}
		}
	}

	return &freshIngredients, &availableIngredients, nil
}
