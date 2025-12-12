package io

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

	problems := abstractions.Problems{
		Numbers: make([][]string, DefaultSize),
	}

	scanner := bufio.NewScanner(r.inputFile)
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
			problems.Numbers[index] = make([]string, cellsCount)

			for cellIndex, cell := range strings.SplitN(line, " ", 0) {
				/* Keep the cell as a string so that we can manipulate it later */
				problems.Numbers[index][cellIndex] = cell
			}
		}

		index++
	}

	return &problems, nil
}
