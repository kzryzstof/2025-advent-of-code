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
			r.extractButtons(line),
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

func (r *FactoryReader) extractLightIndicators(
	line string,
) []*abstractions.Light {

	// Match any sequence of '.' or '#' between square brackets and capture the inner part.
	lightsRegex := regexp.MustCompile(`\[([.#]+)]`)

	matches := lightsRegex.FindAllStringSubmatch(line, -1)

	lights := make([]*abstractions.Light, 0, len(matches))

	for _, m := range matches {
		if len(m) < 2 {
			continue
		}

		inner := m[1]

		for i := 0; i < len(inner); i++ {
			ch := inner[i]
			lights = append(lights, abstractions.NewLight(
				ch == '#',
			))
		}
	}

	return lights
}

func (r *FactoryReader) extractButtons(
	line string,
) []*abstractions.ButtonGroup {

	// Match any sequence of numbers between parenthesis and capture the inner part.
	buttonGroupsRegex := regexp.MustCompile(`\(([\d,*]+)\)`)

	matches := buttonGroupsRegex.FindAllStringSubmatch(line, -1)

	buttonGroups := make([]*abstractions.ButtonGroup, 0, len(matches))

	for _, m := range matches {
		if len(m) < 2 {
			continue
		}

		inner := m[1] // the string between [ and ]
		lightIndicators := strings.Split(inner, ",")

		buttons := make([]*abstractions.Button, len(lightIndicators))

		// For each character inside the brackets, create a LightIndicator
		for i := 0; i < len(lightIndicators); i++ {
			lightNumber, _ := strconv.ParseInt(lightIndicators[i], 10, 64)
			buttons[i] = &abstractions.Button{CounterIndex: int(lightNumber)}
		}

		buttonGroups = append(buttonGroups, &abstractions.ButtonGroup{Buttons: buttons})
	}

	return buttonGroups
}

func (r *FactoryReader) extractVoltages(
	line string,
) []*abstractions.Voltage {

	voltagesRegex := regexp.MustCompile(`\{([\d,*]+)\}`)

	matches := voltagesRegex.FindAllStringSubmatch(line, -1)

	voltages := make([]*abstractions.Voltage, 0, len(matches))

	for _, m := range matches {
		if len(m) < 2 {
			continue
		}

		inner := m[1] // the string between [ and ]
		voltageValues := strings.Split(inner, ",")

		// For each character inside the brackets, create a LightIndicator
		for i := 0; i < len(voltageValues); i++ {
			voltage, _ := strconv.ParseUint(voltageValues[i], 10, 64)
			voltages = append(voltages, abstractions.NewVoltage(uint32(voltage)))
		}
	}

	return voltages
}
