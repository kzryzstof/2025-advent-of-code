package io

import (
	"bufio"
	"day_4/internal/abstractions"
	"fmt"
	"os"
)

const (
	DefaultRowsSliceCapacity = 1000
)

type DepartmentParser struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*DepartmentParser, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &DepartmentParser{
		inputFile,
	}, nil
}

func (p *DepartmentParser) Read() (*abstractions.Department, error) {

	section := abstractions.Department{
		Rows: make([]abstractions.Row, 0, DefaultRowsSliceCapacity),
	}

	scanner := bufio.NewScanner(p.inputFile)

	rowIndex := uint(0)

	for scanner.Scan() {

		line := scanner.Text()

		row := abstractions.Row{
			Number: rowIndex + 1,
			Spots:  make([]abstractions.Spot, len(line)),
		}

		/* Parses the current spots */
		for spotIndex, spot := range line {
			row.Spots[spotIndex] = abstractions.Spot(spot)
		}

		section.Rows = append(section.Rows, row)
		rowIndex++
	}

	return &section, nil
}
