package io

import (
	"bufio"
	"day_9/internal/abstractions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DefaultRedTilesCapacity = 500
)

type RedTilesReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*RedTilesReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &RedTilesReader{
		inputFile,
	}, nil
}

func (r *RedTilesReader) Read() (*abstractions.MovieTheater, error) {

	scanner := bufio.NewScanner(r.inputFile)

	redTiles := make([]*abstractions.Tile, 0, DefaultRedTilesCapacity)

	for scanner.Scan() {
		line := scanner.Text()

		coordinates := strings.Split(line, ",")

		if len(coordinates) != 2 {
			return nil, fmt.Errorf("invalid coordinates: %s", line)
		}

		x, err := strconv.ParseUint(coordinates[0], 10, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting X coordinates '%s': %w", line, err)
		}

		y, err := strconv.ParseUint(coordinates[1], 10, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting Y coordinates '%s': %w", line, err)
		}

		redTiles = append(
			redTiles,
			&abstractions.Tile{
				X: x,
				Y: y,
			})
	}

	return &abstractions.MovieTheater{
		RedTiles: redTiles,
	}, nil
}
