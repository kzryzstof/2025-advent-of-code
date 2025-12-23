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
) {
	fmt.Printf("%d | > %d nodes\r", currentCount, len(slice))
}
