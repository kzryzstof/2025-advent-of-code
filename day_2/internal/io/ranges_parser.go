package io

import (
	"bufio"
	"day_2/internal/abstractions"
	"fmt"
	"os"
	"strings"
)

const (
	DefaultRangesCapacity = 1000
)

type RangesReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*RangesReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &RangesReader{
		inputFile,
	}, nil
}

func (p *RangesReader) Read() []abstractions.Range {

	ranges := make([]abstractions.Range, 0, DefaultRangesCapacity)

	scanner := bufio.NewScanner(p.inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		for _, rangeStr := range strings.Split(line, ",") {
			productIds := strings.Split(rangeStr, "-")

			if len(productIds) != 2 {
				fmt.Printf("invalid range: %s", rangeStr)
				os.Exit(1)
			}

			fromProd, err := abstractions.NewProduct(productIds[0])
			if err != nil {
				fmt.Printf("failed to create from product: %v", err)
				os.Exit(1)
			}

			toProd, err := abstractions.NewProduct(productIds[1])
			if err != nil {
				fmt.Printf("failed to create to product: %v", err)
				os.Exit(1)
			}

			ranges = append(
				ranges,
				abstractions.Range{From: *fromProd, To: *toProd},
			)
		}
	}

	return ranges
}
