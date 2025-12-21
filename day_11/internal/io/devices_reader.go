package io

import (
	"bufio"
	"day_11/internal/abstractions"
	"fmt"
	"os"
	"strings"
)

const (
	DefaultDevicesCapacity = 1000
	All                    = -1
)

type DevicesReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*DevicesReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &DevicesReader{
		inputFile,
	}, nil
}

func (r *DevicesReader) Read() ([]*abstractions.Device, error) {

	scanner := bufio.NewScanner(r.inputFile)

	devices := make([]*abstractions.Device, 0, DefaultDevicesCapacity)

	for scanner.Scan() {
		line := scanner.Text()

		devices = append(devices, r.extractDevice(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *DevicesReader) extractDevice(
	line string,
) *abstractions.Device {

	nameOutputs := strings.Split(line, ":")
	name := nameOutputs[0]
	outputs := strings.Split(nameOutputs[1][1:], " ")

	return abstractions.NewDevice(
		name,
		outputs,
	)
}
