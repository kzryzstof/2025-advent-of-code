package main

import (
	"day_3/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	banksReader := initializeReader(inputFile)

	/* Reads all the banks */
	banks, err := banksReader.Read()

	if err != nil {
		os.Exit(1)
	}

	totalVoltage := uint(0)

	for _, bank := range banks {
		totalVoltage += bank.GetHighestVoltage()
	}

	/* Prints the result */
	fmt.Printf("Sum of all the highest voltage from the %d banks: %d\n", len(banks), totalVoltage)
}

func initializeReader(
	inputFile []string,
) *io.BanksReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
