package io

import (
	"bufio"
	"day_10/internal/abstractions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DefaultJunctionBoxesCapacity = 1000
)

type PlaygroundReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*PlaygroundReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &PlaygroundReader{
		inputFile,
	}, nil
}

func (r *PlaygroundReader) Read() (*abstractions.Playground, error) {

	scanner := bufio.NewScanner(r.inputFile)

	junctionBoxes := make([]*abstractions.JunctionBox, 0, DefaultJunctionBoxesCapacity)

	for scanner.Scan() {
		line := scanner.Text()

		coordinates := strings.Split(line, ",")

		x, err := strconv.ParseUint(coordinates[0], 10, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting coordinates '%s': %w", line, err)
		}

		y, err := strconv.ParseUint(coordinates[1], 10, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting coordinates '%s': %w", line, err)
		}

		z, err := strconv.ParseUint(coordinates[2], 10, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting coordinates '%s': %w", line, err)
		}

		junctionBoxes = append(
			junctionBoxes,
			&abstractions.JunctionBox{
				Position: abstractions.Position{
					X: x,
					Y: y,
					Z: z,
				},
			})
	}

	return &abstractions.Playground{
		JunctionBoxes: junctionBoxes,
	}, nil
}
