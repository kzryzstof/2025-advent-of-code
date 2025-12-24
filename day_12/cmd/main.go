package main

import (
	"day_12/internal/algorithms"
	"day_12/internal/io"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	elapsed := time.Since(startTime)

	/* 	Initializes the reader */
	reader := initializeReader(inputFile)

	/* Reads all the presents and trees from the cavern */
	cavern, err := reader.Read()

	if err != nil {
		os.Exit(1)
	}

	/*
					Some precomputation could be useful.

					To reduce the complexity of polygons by grouping them upfront.

					For instance, polygons 0 and 2 could easily be combined:

					0:		2:			-> Super Shape 0-2
					.##		###			   .##		###     		..#
					##.		.##			   ##.		.## -> (flip) 	.##
					#..		..#			   #..		..#    			###

			 							 -> Super Shape 0-2 is 3x4 now.
				  						  .##
										  ###
										  ###
									      ###

			 						     -> Super Shape 2-2 is 3x4 (which one missing spot)
										  ###
										  ###
										  ###
									      ###

					Since we can know the new dimensions, we could just quickly an roughly
					figure out if there is enough space without trying all the solutions.

					Default
					3x3

					Index 0		Index 1		Index 2		Index 3		Index 4		Index5
					0+2	-> 3x4													5+5	-> 	4x4


		###
		####
		####
		 ###
	*/
	fmt.Printf("Found %d presents\n", cavern.GetPresentsCount())
	fmt.Printf("Found %d Christmas trees\n", cavern.GetChristmasTreesCount())

	algorithms.ComputePermutations(cavern.GetPresents())

	/* Prints the result */
	fmt.Printf("Execution time: %v\n", elapsed)
}

func initializeReader(
	inputFile []string,
) *io.CavernReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Reader initialized: %v\n", reader)
	return reader
}
