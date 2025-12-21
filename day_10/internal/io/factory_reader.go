package io

import (
	"bufio"
	"day_10/internal/abstractions"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	DefaultMachinesCapacity = 250
	All                     = -1
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

		machines = append(machines, abstractions.NewMachine(
			r.extractButtonGroups(line),
			r.extractVoltages(line),
		))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &abstractions.Factory{
		Machines: machines,
	}, nil
}

func (r *FactoryReader) extractButtonGroups(
	line string,
) []*abstractions.ButtonGroup {

	// Match any sequence of numbers between parenthesis and capture the inner part.
	buttonGroupsRegex := regexp.MustCompile(`\(([\d,*]+)\)`)

	matches := buttonGroupsRegex.FindAllStringSubmatch(line, All)

	buttonGroups := make([]*abstractions.ButtonGroup, 0, len(matches))

	for _, m := range matches {
		if len(m) < 2 {
			continue
		}

		buttonNumbers := strings.Split(m[1], ",")

		buttons := make([]*abstractions.Button, len(buttonNumbers))

		// For each character inside the brackets, create a LightIndicator
		for i := 0; i < len(buttonNumbers); i++ {
			buttonNumber, _ := strconv.ParseInt(buttonNumbers[i], 10, 32)
			buttons[i] = abstractions.NewButton(int(buttonNumber))
		}

		buttonGroups = append(buttonGroups, &abstractions.ButtonGroup{Buttons: buttons})
	}

	return buttonGroups
}

func (r *FactoryReader) extractVoltages(
	line string,
) []*abstractions.Voltage {

	voltagesRegex := regexp.MustCompile(`\{([\d,*]+)}`)

	matches := voltagesRegex.FindAllStringSubmatch(line, All)

	voltages := make([]*abstractions.Voltage, 0, len(matches))

	for _, m := range matches {

		if len(m) < 2 {
			continue
		}

		voltageValues := strings.Split(m[1], ",")

		for i := 0; i < len(voltageValues); i++ {
			voltage, _ := strconv.ParseUint(voltageValues[i], 10, 64)
			voltages = append(voltages, abstractions.NewVoltage(int64(voltage)))
		}
	}

	return voltages
}
