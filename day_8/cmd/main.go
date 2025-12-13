package main

import (
	"day_8/internal/app"
	"day_8/internal/io"
	"fmt"
	"os"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads all the junction boxes from the playground */
	playground, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/* Connects pairs of junction boxes with only the specific number of cables */
	circuits := app.ConnectJunctionBoxes(playground, 1000, true)

	biggestCircuits := circuits.GetBiggestCircuits(3)

	/* Prints the result */
	fmt.Printf("The elves have created %d circuits\n", circuits.Count())

	fmt.Printf(
		"Biggest circuits are %d, %d and %d junction boxes (%d)\n",
		biggestCircuits[0].Count(),
		biggestCircuits[1].Count(),
		biggestCircuits[2].Count(),
		biggestCircuits[0].Count()*biggestCircuits[1].Count()*biggestCircuits[2].Count())

	//	28594	too low
	//	990		too low
}

func initializeReader(
	inputFile []string,
) *io.PlaygroundReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
