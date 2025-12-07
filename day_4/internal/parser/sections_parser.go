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

		for scanner.Scan() {
			line := scanner.Text()

			/* Initializes the three rows with the spots upfront */
			currentRowIndex := 0
			rowToAnalyze := -1

			if p.rowsCount == 0 && len(section.Rows[0].Spots) == 0 {
				section.Rows[0] = abstractions.Row{Spots: make([]abstractions.Spot, len(line))}
				section.Rows[1] = abstractions.Row{Spots: make([]abstractions.Spot, len(line))}
				section.Rows[2] = abstractions.Row{Spots: make([]abstractions.Spot, len(line))}
				currentRowIndex = 1
			} else if p.rowsCount == 1 {
				currentRowIndex = 2
				rowToAnalyze = 1
			} else {
				/* Copies the second row to the first one */
				for spotIndex, spot := range section.Rows[1].Spots {
					section.Rows[0].Spots[spotIndex] = spot
				}
				/* Copies the third row to the second one */
				for spotIndex, spot := range section.Rows[2].Spots {
					section.Rows[1].Spots[spotIndex] = spot
				}
				currentRowIndex = 2
				rowToAnalyze = 1
			}

			/* Parses the current spots */
			for spotIndex, spot := range line {
				section.Rows[currentRowIndex].Spots[spotIndex] = abstractions.Spot(spot)
			}

			p.rowsCount++

			/* Sends the section for analysis */
			if rowToAnalyze != -1 {
				section.RowIndex = rowToAnalyze
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

func (p *SectionsParser) Sections() <-chan abstractions.Section {
	return p.sectionsChannel
}
