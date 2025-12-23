package io

import (
	"bufio"
	"day_12/internal/abstractions"
	"fmt"
	"os"
	"strings"
)

const (
	DefaultPresentsCount = 6
)

type CavernReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*CavernReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &CavernReader{
		inputFile,
	}, nil
}

func (r *CavernReader) Read() (*abstractions.Cavern, error) {

	scanner := bufio.NewScanner(r.inputFile)

	presents := r.extractPresents(scanner)

	return abstractions.NewCavern(
		presents,
	), nil
}

func (r *CavernReader) extractPresents(
	scanner *bufio.Scanner,
) map[uint]*abstractions.Present {

	presents := make(map[uint]*abstractions.Present)

	for presentIndex := uint(0); presentIndex < DefaultPresentsCount; presentIndex++ {
		/* Skips the index */
		scanner.Scan()

		/* Extract the shape */
		var b strings.Builder
		for row := 0; row < 3; row++ {
			scanner.Scan()
			b.WriteString(scanner.Text())
		}

		/* Reads the empty */
		scanner.Scan()
		scanner.Text()

		shape := make([][]byte, 3)

		for row := 0; row < 3; row++ {
			shape[row] = make([]byte, 3)
			for col := 0; col < 3; col++ {
				if b.String()[row*3+col] == '#' {
					shape[row][col] = byte(1)
				} else {
					shape[row][col] = byte(0)
				}
			}
		}

		presents[presentIndex] = abstractions.NewPresent(
			shape,
			3,
			3,
		)
	}

	return presents
}
