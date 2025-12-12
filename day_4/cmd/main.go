package main

import (
	"day_4/internal/app"
	"day_4/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser and processor */
	reader := initializeReader(inputFile)

	/* Reads all the rows */
	department, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	forklift := app.NewForklift(false)
	forklift.RemoveRolls(department)

	/* Prints the total number of accessible rolls */
	fmt.Printf("Number of accessible rolls in the %d rows of the department: %d\n", len(department.Rows), forklift.GetAccessedRollsCount())
}

func initializeReader(
	inputFile []string,
) *io.DepartmentParser {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
