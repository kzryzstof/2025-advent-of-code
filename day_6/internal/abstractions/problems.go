package abstractions

import (
	"day_6/internal/app"
)

// 10337550451352: too low
// 11759359657848: too high
// 11708563457309

type Problems struct {
	Numbers    [][]string
	Operations []string
}

func (p *Problems) ComputeTotal() (uint64, error) {

	total := uint64(0)

	for columnIndex, operation := range p.Operations {

		/* Reads all the numbers from the current column */
		cells := p.readNumbers(columnIndex)

		/* Transposes the values to get the numbers in each row */
		numbers := app.TransposeColumns(cells)

		/* Performs the operation on the numbers */
		columnTotal, err := app.Compute(operation, numbers)

		if err != nil {
			return 0, err
		}

		total += columnTotal
	}

	return total, nil
}

func (p *Problems) readNumbers(columnIndex int) []string {

	numbersCount := len(p.Numbers)

	cells := make([]string, numbersCount)

	for rowIndex := 0; rowIndex < numbersCount; rowIndex++ {
		cells[rowIndex] = p.Numbers[rowIndex][columnIndex]
	}

	return cells
}
