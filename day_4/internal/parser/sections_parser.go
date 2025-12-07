package parser

import (
	"bufio"
	"day_4/internal/abstractions"
	"fmt"
	"os"
)

type SectionsParser struct {
	rowsCount       uint
	section         *abstractions.Section
	analysisStarted bool
	rowIndexToFill  uint
	scanner         *bufio.Scanner
}

func NewParser(
	filePath string,
) (*SectionsParser, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	scanner := bufio.NewScanner(inputFile)

	return &SectionsParser{
		0,
		&abstractions.Section{Rows: make([]abstractions.Row, 3), RowIndex: 1},
		false,
		uint(1),
		scanner,
	}, nil
}

func (p *SectionsParser) GetRowsCount() uint {
	return p.rowsCount
}

func (p *SectionsParser) ReadNextRow() (*abstractions.Section, bool) {

	if eof := p.readRow(); !eof {
		return nil, false
	}

	/* The first row is filled but the second one is needed to start the analysis */
	if p.rowsCount == 1 {
		if eof := p.readRow(); !eof {
			return nil, false
		}
	}

	return p.section, true
}

func (p *SectionsParser) readRow() bool {

	eof := p.scanner.Scan()

	if !eof {
		/* No more rows but we still have the last row to analyze */
		if p.section.RowIndex == 2 {
			return false
		}
		p.section.RowIndex = 2
		return true
	}

	line := p.scanner.Text()

	fmt.Printf("Parser | %03d | Read line | %s", p.rowsCount+1, line)

	if p.rowsCount == 0 {
		p.allocateRows(line)
		/* We start filling the second row, the first one being empty here */
		/* The analysis cannot start yet since the row below is not filled yet */
	} else if p.rowsCount == 1 {
		/* No copy at this moment */
		/* From now on, the second row will be the one to be filled */
		p.rowIndexToFill = 2
		/* The analysis can start now */
		p.analysisStarted = true
	} else {
		/* Copies the second row to the first one, and the third to the second one*/
		p.copyRow(1, 0)
		p.copyRow(2, 1)
	}

	/* Parses the current spots */
	for spotIndex, spot := range line {
		p.section.Rows[p.rowIndexToFill].Spots[spotIndex] = abstractions.Spot(spot)
	}

	p.section.Rows[p.rowIndexToFill].Number = p.rowsCount + 1

	p.rowsCount++

	return true
}

func (p *SectionsParser) copyRow(
	fromRowIndex uint,
	toRowIndex uint,
) {
	p.section.Rows[toRowIndex].Number = p.section.Rows[fromRowIndex].Number

	for spotIndex, spot := range p.section.Rows[fromRowIndex].Spots {
		p.section.Rows[toRowIndex].Spots[spotIndex] = spot
	}
}

func (p *SectionsParser) allocateRows(
	line string,
) {
	p.section.Rows[0] = abstractions.Row{Number: 0, Spots: make([]abstractions.Spot, len(line))}
	p.section.Rows[1] = abstractions.Row{Number: 1, Spots: make([]abstractions.Spot, len(line))}
	p.section.Rows[2] = abstractions.Row{Number: 2, Spots: make([]abstractions.Spot, len(line))}
}
