package parser

import (
	"bufio"
	"day_5/internal/abstractions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IngredientsParser struct {
	Fresh *abstractions.FreshIngredients
}

func NewParser(
	filePath string,
) (*IngredientsParser, error) {

	freshIngredients, err := readIngredients(filePath)

	if err != nil {
		return nil, err
	}

	return &IngredientsParser{
		freshIngredients,
	}, nil
}

func readIngredients(
	filePath string,
) (*abstractions.FreshIngredients, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	freshIngredients := abstractions.FreshIngredients{
		Ranges: make([]abstractions.IngredientRange, 0, 500),
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Trim(line, " \n\t") == "" {
			break
		}

		rangesSlice := strings.Split(line, "-")

		if len(rangesSlice) != 2 {
			return nil, fmt.Errorf("invalid range: %s", line)
		}

		from, err := strconv.ParseUint(rangesSlice[0], 10, 64)

		if err != nil {
			return nil, err
		}

		to, err := strconv.ParseUint(rangesSlice[1], 10, 64)

		if err != nil {
			return nil, err
		}

		if from > to {
			return nil, fmt.Errorf("invalid range: %s", line)
		}

		fromIngredientId := abstractions.IngredientId(from)
		toIngredientId := abstractions.IngredientId(to)

		newRange := abstractions.IngredientRange{
			From: fromIngredientId,
			To:   toIngredientId,
		}

		freshIngredients.Ranges = append(freshIngredients.Ranges, newRange)
	}

	return &abstractions.FreshIngredients{Ranges: freshIngredients.Ranges}, nil
}
