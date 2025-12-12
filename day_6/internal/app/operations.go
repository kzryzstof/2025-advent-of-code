package app

import "fmt"

func Compute(
	operation string,
	numbers []uint64,
) (uint64, error) {
	/* Performs the operation on the numbers */
	switch operation {
	case `*`:
		return applyOperation(numbers, func(a, b uint64) uint64 { return a * b }), nil
	case `+`:
		return applyOperation(numbers, func(a, b uint64) uint64 { return a + b }), nil
	}

	return 0, fmt.Errorf("invalid operation: %s", operation)
}

func applyOperation(numbers []uint64, operation func(uint64, uint64) uint64) uint64 {

	total := numbers[0]

	for _, number := range numbers[1:] {
		total = operation(total, number)
	}

	return total
}
