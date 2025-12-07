package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"day_2/internal/abstractions"
)

type RangesParser struct {
	inputFile     *os.File
	syncWaitGroup *sync.WaitGroup
	rangesChannel chan abstractions.Range
	rangesCount   int
}

func NewParser(
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
		0,
	}, nil
}

func (p *RangesParser) GetRangesCount() int {
	return p.rangesCount
}

func (p *RangesParser) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

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

				p.rangesCount++

				p.rangesChannel <- abstractions.Range{From: *fromProd, To: *toProd}
			}
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
