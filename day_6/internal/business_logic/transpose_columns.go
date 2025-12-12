package business_logic

import (
	"day_6/internal/extensions"
	"fmt"
	"strconv"
)

func TransposeColumns(cells []string) []uint64 {

	/* Indicates how number to build */
	numbersCount := 0

	/* Temporary accumulator for the digits */
	accumulator := make([][]string, 0)

	for _, number := range cells {
		/* Converts the number to digits, and reverse the order */
		digits := toDigitsReversed(number)
		/* The count of numbers to produce depends on the longest number */
		numbersCount = max(numbersCount, len(digits))
		accumulator = append(accumulator, digits)
	}

	/* Rebuilds the numbers, digit by digit */
	numbers := make([]uint64, numbersCount)

	for numberIndex := 0; numberIndex < numbersCount; numberIndex++ {

		numberText := ""

		for digitIndex := 0; digitIndex < len(accumulator); digitIndex++ {
			if numberIndex >= len(accumulator[digitIndex]) {
				continue
			}
			numberText += accumulator[digitIndex][numberIndex]
		}

		value, err := strconv.ParseUint(numberText, 10, 64)

		if err != nil {
			panic(fmt.Errorf("Error parsing number '%s': %v\n", numberText, err))
		}

		numbers[numberIndex] = value
	}

	extensions.Reverse(numbers)

	return numbers
}

func toDigitsReversed(number string) []string {
	digits := extensions.ToDigits(number)
	return digits
}
