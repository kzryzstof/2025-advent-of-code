package io

import (
	"bufio"
	"day_3/internal/abstractions"
	"fmt"
	"os"
	"strconv"
)

const (
	DefaultBanksSliceCapacity = 1000
)

type BanksReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*BanksReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &BanksReader{
		inputFile,
	}, nil
}

func (p *BanksReader) Read() ([]abstractions.Bank, error) {

	banks := make([]abstractions.Bank, 0, DefaultBanksSliceCapacity)

	scanner := bufio.NewScanner(p.inputFile)

	for scanner.Scan() {
		line := scanner.Text()

		batteries := make([]abstractions.Battery, 0)

		for _, batteryVoltageRating := range line {
			batteryVoltageRatingInt, err := strconv.Atoi(string(batteryVoltageRating))

			if err != nil {
				fmt.Printf("Error converting battery voltage rating '%b' to int: %v\n", batteryVoltageRating, err)
				return nil, err
			}

			batteries = append(
				batteries,
				abstractions.Battery{Voltage: abstractions.VoltageRating(batteryVoltageRatingInt)},
			)
		}

		banks = append(banks, abstractions.Bank{Batteries: batteries})
	}

	return banks, nil
}
