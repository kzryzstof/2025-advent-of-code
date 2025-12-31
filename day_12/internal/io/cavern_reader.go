package io

import (
	"bufio"
	"day_12/internal/abstractions"
	"day_12/internal/maths"
	"fmt"
	"os"
	"strings"
)

const (
	DefaultPresentsCount = 6
)

type CavernReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*CavernReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &CavernReader{
		inputFile,
	}, nil
}

func (r *CavernReader) Read() (*abstractions.Cavern, error) {

	scanner := bufio.NewScanner(r.inputFile)

	presents := r.extractPresents(scanner)
	christmasTrees, err := r.extractChristmasTrees(scanner)

	if err != nil {
		return nil, err
	}

	return abstractions.NewCavern(
		abstractions.NewPresents(presents),
		christmasTrees,
	), nil
}

func (r *CavernReader) extractPresents(
	scanner *bufio.Scanner,
) map[abstractions.PresentIndex]*abstractions.Present {

	presents := make(map[abstractions.PresentIndex]*abstractions.Present)

	for presentIndex := uint(0); presentIndex < DefaultPresentsCount; presentIndex++ {
		/* Skips the index */
		scanner.Scan()

		/* Extract the shape */
		var b strings.Builder
		for row := 0; row < 3; row++ {
			scanner.Scan()
			b.WriteString(scanner.Text())
		}

		/* Reads the empty */
		scanner.Scan()
		scanner.Text()

		shape := make([][]int8, 3)

		empty, occupied := 0, 0

		for row := 0; row < 3; row++ {
			shape[row] = make([]int8, 3)
			for col := 0; col < 3; col++ {
				if b.String()[row*3+col] == '#' {
					shape[row][col] = int8(1)
					occupied++
				} else {
					shape[row][col] = abstractions.E
					empty++
				}
			}
		}

		fillRatio := float64(occupied) / float64(occupied+empty)

		fmt.Printf("Present %d: %d occupied, %d empty, fill ratio: %.2f\n", presentIndex+1, occupied, empty, fillRatio)

		presents[abstractions.PresentIndex(presentIndex)] = abstractions.NewPresent(
			abstractions.PresentIndex(presentIndex),
			abstractions.NewShape(
				maths.Dimension{
					Wide: 3,
					Long: 3,
				},
				shape,
			),
		)
	}

	return presents
}

func (r *CavernReader) extractChristmasTrees(
	scanner *bufio.Scanner,
) ([]*abstractions.ChristmasTree, error) {

	christmasTrees := make([]*abstractions.ChristmasTree, 0)

	christmasTreesCount := 1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var wide, long uint
		var presents0Count, presents1Count, presents2Count, presents3Count, presents4Count, presents5Count uint
		_, err := fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &wide, &long, &presents0Count, &presents1Count, &presents2Count, &presents3Count, &presents4Count, &presents5Count)

		if err != nil {
			return nil, fmt.Errorf("Error reading Christmas trees: %v\n", err)
		}

		christmasTrees = append(christmasTrees, abstractions.NewChristmasTree(
			abstractions.ChristmasTreeIndex(christmasTreesCount),
			wide,
			long,
			map[abstractions.PresentIndex]uint{
				0: presents0Count,
				1: presents1Count,
				2: presents2Count,
				3: presents3Count,
				4: presents4Count,
				5: presents5Count,
			}),
		)

		christmasTreesCount++
	}

	return christmasTrees, nil
}
