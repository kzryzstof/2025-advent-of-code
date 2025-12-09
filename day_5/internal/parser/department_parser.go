package parser

import (
	"bufio"
	"day_5/internal/abstractions"
	"fmt"
	"os"
)

type DepartmentParser struct {
	rowsCount uint
	Section   *abstractions.Department
}

func NewParser(
	filePath string,
) (*DepartmentParser, error) {

	rowsCount, err := countRows(filePath)

	if err != nil {
		return nil, err
	}

	section, err := readAllRows(filePath, rowsCount)

	if err != nil {
		return nil, err
	}

	return &DepartmentParser{
		rowsCount,
		section,
	}, nil
}

func (p *DepartmentParser) GetRowsCount() uint {
	return p.rowsCount
}

func readAllRows(
	filePath string,
	rowsCount uint,
) (*abstractions.Department, error) {

	section := abstractions.Department{
		Rows: make([]abstractions.Row, rowsCount),
	}

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	rowIndex := uint(0)

	for scanner.Scan() {

		line := scanner.Text()

		fmt.Printf("Parser | %03d | Read line | %s\n", rowIndex+1, line)

		section.Rows[rowIndex] = abstractions.Row{
			Number: rowIndex + 1,
			Spots:  make([]abstractions.Spot, len(line)),
		}

		/* Parses the current spots */
		for spotIndex, spot := range line {
			section.Rows[rowIndex].Spots[spotIndex] = abstractions.Spot(spot)
		}

		rowIndex++
	}

	return &section, nil
}

func countRows(
	filePath string,
) (uint, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return 0, err
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	rowsCount := uint(0)

	for scanner.Scan() {
		rowsCount++
	}

	return rowsCount, nil
}
