package parser

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"day_2/internal/abstractions"
)

type RangesParser struct {
	inputFile     *os.File
	syncWaitGroup *sync.WaitGroup
	rangesChannel chan abstractions.Range
}

func New(
	filePath string,
	waitGroup *sync.WaitGroup,
) (*RangesParser, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &RangesParser{
		inputFile,
		waitGroup,
		make(chan abstractions.Range),
	}, nil
}

func (p *RangesParser) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		scanner := bufio.NewScanner(p.inputFile)

		for scanner.Scan() {
			line := scanner.Text()

			fmt.Println(line)
			//p.rangesChannel <- rotation
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading file: %w", err)
			os.Exit(1)
		}

		close(p.rangesChannel)
	}()
}

func (p *RangesParser) Ranges() <-chan abstractions.Range {
	return p.rangesChannel
}
