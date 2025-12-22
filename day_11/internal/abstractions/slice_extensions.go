package abstractions

import "fmt"

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
) {
	fmt.Printf("%d | > ", currentCount)

	for _, sliceItem := range slice {
		fmt.Print(sliceItem + ", ")
	}

	fmt.Print("\r")
}
