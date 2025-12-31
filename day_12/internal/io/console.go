package io

import (
	"day_12/internal/abstractions"
	"fmt"
	"math"
)

func PrintCells(
	slice [][]int8,
) {
	for _, row := range slice {
		for _, cell := range row {
			if cell == abstractions.E {
				fmt.Print(".")
			} else if cell == -1 {
				fmt.Print("#")
			} else {
				fmt.Print(fmt.Sprintf("%d", cell))
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func PrintBothCells(
	leftSlice [][]int8,
	rightSlice [][]int8,
) {
	maxRows := int(math.Max(float64(len(leftSlice)), float64(len(rightSlice))))

	printRowCell := func(row []int8) {
		for _, cell := range row {
			if cell == abstractions.E {
				fmt.Print(".")
			} else {
				fmt.Print(fmt.Sprintf("%d", cell))
			}
		}
	}

	for rowIndex := 0; rowIndex < maxRows; rowIndex++ {
		if rowIndex < len(leftSlice) {
			printRowCell(leftSlice[rowIndex])
		} else {
			for colIndex := uint(0); colIndex < abstractions.MaximumShapeSize; colIndex++ {
				fmt.Print(" ")
			}
		}
		fmt.Print(" | ")
		if rowIndex < len(rightSlice) {
			printRowCell(rightSlice[rowIndex])
		}
		fmt.Println()
	}
}
