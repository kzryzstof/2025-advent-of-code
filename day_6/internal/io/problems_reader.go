package io

import (
	"bufio"
	"day_6/internal/abstractions"
	"fmt"
	"os"
)

const (
	DefaultNumbersRows = 4
	DefaultColumns     = 1000
)

type ProblemsReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*ProblemsReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &ProblemsReader{
		inputFile,
	}, nil
}

func (r *ProblemsReader) Read() (*abstractions.Problems, error) {

	scanner := bufio.NewScanner(r.inputFile)

	/* Reads all the lines first because we need to figure out width of each cell first */
	/* The width will be read from the last line */
	lines := make([]string, 0, DefaultNumbersRows+1)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	/* Parses all the number now, column by column */
	problems := abstractions.Problems{
		Numbers:    make([][]string, DefaultNumbersRows),
		Operations: make([]string, 0, DefaultColumns),
	}

	numbersCount := len(lines) - 1
	operationsLine := lines[len(lines)-1]

	cellFromIndex := 0
	cellToIndex := 0
	currentOperation := int32(0)

	for characterIndex, character := range operationsLine {
		if currentOperation == 0 && character != ' ' {
			/* Initial operation */
			cellFromIndex = characterIndex
			currentOperation = character
		} else if character != ' ' || characterIndex+1 == len(operationsLine) {
			/* New operation found (or it is the last character) */
			cellToIndex = characterIndex - 1 /* not current position */ - 1 /* remove the column spacer */

			for numberIndex := 0; numberIndex < numbersCount; numberIndex++ {
				numberLine := lines[numberIndex]
				number := numberLine[cellFromIndex : cellToIndex+1]
				problems.Numbers[numberIndex] = append(problems.Numbers[numberIndex], number)
			}

			problems.Operations = append(problems.Operations, string(currentOperation))

			cellFromIndex = characterIndex
			currentOperation = character
		}
	}

	/* Parses the last number column */
	for numberIndex := 0; numberIndex < numbersCount; numberIndex++ {
		numberLine := lines[numberIndex]
		number := numberLine[cellFromIndex:len(lines[numberIndex])]
		problems.Numbers[numberIndex] = append(problems.Numbers[numberIndex], number)
	}
	problems.Operations = append(problems.Operations, string(currentOperation))

	return &problems, nil
}
