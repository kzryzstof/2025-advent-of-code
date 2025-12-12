package extensions

func ToDigits(number string) []string {
	digits := make([]string, len(number))

	for index, digit := range number {
		digits[index] = string(digit)
	}

	return digits
}
