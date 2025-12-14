package io

import (
	"bufio"
	"day_10/internal/abstractions"
	"fmt"
	"os"
	"regexp"
)

const (
	DefaultMachinesCapacity = 250
)

type FactoryReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*FactoryReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &FactoryReader{
		inputFile,
	}, nil
}

func (r *FactoryReader) Read() (*abstractions.Factory, error) {

	scanner := bufio.NewScanner(r.inputFile)

	machines := make([]*abstractions.Machine, 0, DefaultMachinesCapacity)

	for scanner.Scan() {
		line := scanner.Text()

		/* Parses the light indicators */
		lightIndicators := r.extractLightIndicators(line)

		machines = append(machines, &abstractions.Machine{
			LightIndicators: lightIndicators,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &abstractions.Factory{
		Machines: machines,
	}, nil
}

func (r *FactoryReader) extractLightIndicators(
	line string,
) []*abstractions.LightIndicator {

	// Match any sequence of '.' or '#' between square brackets and capture the inner part.
	// Final regexp pattern: \[([.#]+)]
	lightsRegex := regexp.MustCompile(`\[([.#]+)]`)

	matches := lightsRegex.FindAllStringSubmatch(line, -1)

	lightIndicators := make([]*abstractions.LightIndicator, 0, len(matches))

	for _, m := range matches {
		if len(m) < 2 {
			continue
		}

		inner := m[1] // the string between [ and ]

		// For each character inside the brackets, create a LightIndicator
		for i := 0; i < len(inner); i++ {
			ch := inner[i]
			lightIndicators = append(lightIndicators, &abstractions.LightIndicator{
				IsOn: ch == '#',
			})
		}
	}

	return lightIndicators
}
