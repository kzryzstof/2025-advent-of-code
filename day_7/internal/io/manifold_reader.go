package io

import (
	"bufio"
	"day_7/internal/abstractions"
	"fmt"
	"os"
)

const (
	DefaultSize = 142
)

type ManifoldReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*ManifoldReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &ManifoldReader{
		inputFile,
	}, nil
}

func (r *ManifoldReader) Read() (*abstractions.Manifold, error) {

	scanner := bufio.NewScanner(r.inputFile)

	locations := make([][]abstractions.Location, 0, DefaultSize)

	/* Reads all the lines first because we need to figure out width of each cell first */
	/* The width will be read from the last line */

	for scanner.Scan() {
		line := scanner.Text()

		locationRow := make([]abstractions.Location, 0, len(line))

		for _, char := range line {
			locationRow = append(locationRow, abstractions.Location(char))
		}

		locations = append(locations, locationRow)
	}

	return &abstractions.Manifold{
		Locations: locations,
	}, nil
}
