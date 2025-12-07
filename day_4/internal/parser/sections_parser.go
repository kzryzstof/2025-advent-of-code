package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	return p.banksCount
}

func (p *SectionsParser) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		scanner := bufio.NewScanner(p.inputFile)

		for scanner.Scan() {
			line := scanner.Text()

			batteries := make([]abstractions.Battery, 0)

			for _, batteryVoltageRating := range line {
				batteryVoltageRatingInt, err := strconv.Atoi(string(batteryVoltageRating))

				if err != nil {
					fmt.Printf("Error converting battery voltage rating '%s' to int: %v\n", batteryVoltageRating, err)
					os.Exit(1)
				}

				batteries = append(batteries, abstractions.Battery{Voltage: abstractions.VoltageRating(batteryVoltageRatingInt)})
			}

			p.rowsCount++
			p.sectionsChannel <- abstractions.Bank{Batteries: batteries}
		}

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
