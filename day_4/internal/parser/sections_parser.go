package parser

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"day_4/internal/abstractions"
)

type SectionsParser struct {
	inputFile       *os.File
	syncWaitGroup   *sync.WaitGroup
	sectionsChannel chan abstractions.Section
	rowsCount       int
}

func NewParser(
	filePath string,
	waitGroup *sync.WaitGroup,
) (*SectionsParser, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &SectionsParser{
		inputFile,
		waitGroup,
		make(chan abstractions.Section),
		0,
	}, nil
}

func (p *SectionsParser) GetRowsCount() int {
	return p.rowsCount
}

func (p *SectionsParser) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		scanner := bufio.NewScanner(p.inputFile)

		section := abstractions.Section{Rows: make([]abstractions.Row, 3), RowIndex: 1}

		analysisStarted := false
		rowIndexToFill := 1

		for scanner.Scan() {
			line := scanner.Text()

			if p.rowsCount == 0 {
				p.allocateRows(section, line)
				/* We start filling the second row, the first one being empty here */
				/* The analysis cannot start yet since the row below is not filled yet */
			} else if p.rowsCount == 1 {
				/* No copy at this moment */
				/* From now on, the second row will be the one to be filled */
				rowIndexToFill = 2
				/* The analysis can start now */
				analysisStarted = true
			} else {
				/* Copies the second row to the first one, and the third to the second one*/
				p.copyRow(section, 1, 0)
				p.copyRow(section, 2, 1)
			}

			/* Parses the current spots */
			for spotIndex, spot := range line {
				section.Rows[rowIndexToFill].Spots[spotIndex] = abstractions.Spot(spot)
			}

			section.Rows[rowIndexToFill].Number = p.rowsCount + 1

			p.rowsCount++

			/* Sends the section for analysis */
			if analysisStarted {
				section.RowIndex = 1
				p.sectionsChannel <- section
			}
		}

		/* Do not forget to analyze the last row */
		section.RowIndex = 2
		p.sectionsChannel <- section

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading file: %w", err)
			os.Exit(1)
		}

		close(p.sectionsChannel)
	}()
}

func (p *SectionsParser) copyRow(
	section abstractions.Section,
	fromRowIndex uint,
	toRowIndex uint,
) {
	section.Rows[toRowIndex].Number = section.Rows[fromRowIndex].Number

	for spotIndex, spot := range section.Rows[fromRowIndex].Spots {
		section.Rows[toRowIndex].Spots[spotIndex] = spot
	}
}

func (p *SectionsParser) allocateRows(
	section abstractions.Section,
	line string,
) {
	section.Rows[0] = abstractions.Row{Number: 0, Spots: make([]abstractions.Spot, len(line))}
	section.Rows[1] = abstractions.Row{Number: 1, Spots: make([]abstractions.Spot, len(line))}
	section.Rows[2] = abstractions.Row{Number: 2, Spots: make([]abstractions.Spot, len(line))}
}

func (p *SectionsParser) Sections() <-chan abstractions.Section {
	return p.sectionsChannel
}
