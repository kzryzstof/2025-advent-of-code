package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"day_3/internal/abstractions"
)

type BanksParser struct {
	inputFile     *os.File
	syncWaitGroup *sync.WaitGroup
	banksChannel  chan abstractions.Bank
	rangesCount   int
}

func NewParser(
	filePath string,
	waitGroup *sync.WaitGroup,
) (*BanksParser, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &BanksParser{
		inputFile,
		waitGroup,
		make(chan abstractions.Bank),
		0,
	}, nil
}

func (p *BanksParser) GetRangesCount() int {
	return p.rangesCount
}

func (p *BanksParser) Start() {

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

				batteries = append(batteries, abstractions.Battery{Votalge: abstractions.VoltageRating(batteryVoltageRatingInt)})
			}

			p.banksChannel <- abstractions.Bank{Batteries: batteries}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading file: %w", err)
			os.Exit(1)
		}

		close(p.banksChannel)
	}()
}

func (p *BanksParser) Banks() <-chan abstractions.Bank {
	return p.banksChannel
}
