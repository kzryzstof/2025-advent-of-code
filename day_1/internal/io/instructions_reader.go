package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"day_1/internal/abstractions"
)

const (
	DefaultInstructionsSliceCapacity = 1000
)

type InstructionsReader struct {
	inputFile *os.File
}

func New(
	filePath string,
) (*InstructionsReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &InstructionsReader{
		inputFile,
	}, nil
}

func (r *InstructionsReader) Read() *abstractions.Instructions {

	rotations := make([]abstractions.Rotation, 0, DefaultInstructionsSliceCapacity)

	scanner := bufio.NewScanner(r.inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		direction, err := extractDirection(line)

		if err != nil {
			fmt.Printf("error converting orientation from %s: %w", line, err)
			os.Exit(1)
		}

		distance, err := extractDistance(line)

		if err != nil {
			fmt.Printf("error converting distance '%d' to int: %w", distance, err)
			os.Exit(1)
		}

		rotations = append(
			rotations,
			abstractions.Rotation{
				Direction: direction,
				Distance:  distance,
			})
	}

	return &abstractions.Instructions{
		Rotations: rotations,
	}
}

func extractDistance(
	line string,
) (int, error) {

	distance := line[1:]

	distanceInt, err := strconv.Atoi(distance)

	if err != nil {
		fmt.Printf("Error converting distance '%s' to int: %v\n", distance, err)
		return 0, err
	}

	if distanceInt < 0 {
		fmt.Printf("Unexpected distance '%d'\n", distanceInt)
		return 0, err
	}

	return distanceInt, err
}

func extractDirection(
	line string,
) (abstractions.Direction, error) {

	if line[0] == 'L' {
		return abstractions.Left, nil
	}

	if line[0] == 'R' {
		return abstractions.Right, nil
	}

	return abstractions.Left, fmt.Errorf("Invalid direction '%c'\n", line[0])
}
