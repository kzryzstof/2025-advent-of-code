package abstractions

import (
	"fmt"
)

const (
	Reset   = "\033[0m"
	Reverse = "\033[7m" // Inverts foreground and background
)

func AddOnce(
	slice []string,
	text string,
) []string {
	for _, sliceItem := range slice {
		if sliceItem == text {
			return slice
		}
	}

	return append(slice, text)
}

func Print(
	slice []string,
	currentCount uint,
	encounteredNodes []string,
) {
	fmt.Printf("%d | > %d nodes", currentCount, len(slice))

	//for _, sliceItem := range slice {
	//	if slices.Contains(encounteredNodes, sliceItem) {
	//		/* Invert colors for encountered nodes */
	//		fmt.Print(Reverse + sliceItem + Reset + ", ")
	//	} else {ß
	//		fmt.Print(sliceItem + ", ")
	//	}
	//}

	/*
		A
		|
		B
		| \
		C D
		| /
		E
		|
		F

	*/

	fmt.Print("\r")
}
