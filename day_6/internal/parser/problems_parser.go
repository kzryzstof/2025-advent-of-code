package parser

import (
	"bufio"
	"day_6/internal/abstractions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DefaultSize = 4
)

type ProblemsParser struct {
	Problems *abstractions.Problems
}

func NewParser(
	filePath string,
) (*ProblemsParser, error) {

	problems, err := readProblems(filePath)

	if err != nil {
		return nil, err
	}

	return &ProblemsParser{
		problems,
	}, nil
}

func readProblems(
	filePath string,
) (*abstractions.Problems, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	problems := abstractions.Problems{
		Numbers: make([][]uint64, DefaultSize),
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)
	index := 0

	for scanner.Scan() {

		line := scanner.Text()

		columns := strings.Fields(line)

		cellsCount := len(columns)

		_, err := strconv.ParseUint(columns[0], 10, 64)

		if err != nil {
			problems.Operations = make([]string, cellsCount)

			/* This is supposedly the last line */
			for cellIndex, operation := range columns {
				problems.Operations[cellIndex] = operation
			}
		} else {
			problems.Numbers[index] = make([]uint64, cellsCount)

			for cellIndex, cell := range columns {
				number, err := strconv.ParseUint(cell, 10, 64)

				if err != nil {
					fmt.Printf("Error parsing number: %v\n", err)
					return nil, err
				}

				problems.Numbers[index][cellIndex] = number
			}
		}

		index++
	}

	return &problems, nil
}
