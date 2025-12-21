package abstractions

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
