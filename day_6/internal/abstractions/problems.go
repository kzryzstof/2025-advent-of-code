package abstractions

type Problems struct {
	Numbers    [][]uint64
	Operations []string
}

func (p *Problems) ComputeTotal() uint64 {

	rowsCount := len(p.Numbers)
	total := uint64(0)

	for columnIndex, operation := range p.Operations {

		columnTotal := uint64(0)

		switch operation {
		case "*":
			for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
				number := p.Numbers[rowIndex][columnIndex]
				if rowIndex == 0 {
					columnTotal = number
				} else {
					columnTotal *= number
				}
			}
		case "+":
			for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
				number := p.Numbers[rowIndex][columnIndex]
				if rowIndex == 0 {
					columnTotal = number
				} else {
					columnTotal += number
				}
			}
		}

		total += columnTotal
	}

	return total
}
