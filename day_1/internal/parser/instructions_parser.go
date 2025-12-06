package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"day_1/internal/abstractions"
)

type InstructionsParser struct {
	inputFile        *os.File
	syncWaitGroup    *sync.WaitGroup
	rotationsChannel chan abstractions.Rotation
}

func New(
	filePath string,
	waitGroup *sync.WaitGroup,
) (*InstructionsParser, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &InstructionsParser{
		inputFile,
		waitGroup,
		make(chan abstractions.Rotation),
	}, nil
}

func (p *InstructionsParser) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		scanner := bufio.NewScanner(p.inputFile)

		for scanner.Scan() {
			line := scanner.Text()

			if len(line) < 2 {
				continue
			}

			direction, err := p.extractDirection(line)

			if err != nil {
				fmt.Printf("error converting orientation from %s: %w", line, err)
				os.Exit(1)
			}

			distance, err := p.extractDistance(line)

			if err != nil {
				fmt.Printf("error converting distance '%d' to int: %w", distance, err)
				os.Exit(1)
			}

			rotation := abstractions.Rotation{
				Direction: direction,
				Distance:  distance,
			}

			p.rotationsChannel <- rotation
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading file: %w", err)
			os.Exit(1)
		}

		close(p.rotationsChannel)
	}()
}

func (p *InstructionsParser) extractDistance(line string) (int, error) {
	distance := line[1:]

	distanceInt, err := strconv.Atoi(distance)

	if err != nil {
		fmt.Printf("Error converting distance '%s' to int: %v\n", distance, err)
		return 0, err
	}

	if distanceInt < 0 {
		fmt.Printf("Unexpected distance '%d'\n", distanceInt)
		return 0, err
	}

	return distanceInt, err
}

func (p *InstructionsParser) extractDirection(line string) (abstractions.Direction, error) {

	if line[0] == 'L' {
		return abstractions.Left, nil
	}

	if line[0] == 'R' {
		return abstractions.Right, nil
	}

	return abstractions.Left, fmt.Errorf("Invalid direction '%c'\n", line[0])
}

func (p *InstructionsParser) Rotations() <-chan abstractions.Rotation {
	return p.rotationsChannel
}
