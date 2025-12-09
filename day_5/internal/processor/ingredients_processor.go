package processor

import (
	"day_5/internal/abstractions"
)

const (
	minAccessibleRolls = 4
)

type DepartmentProcessor struct {
	rollsAccessed uint
}

func NewProcessor() *DepartmentProcessor {
	return &DepartmentProcessor{
		0,
	}
}

func (p *DepartmentProcessor) GetTotalAccessibleRolls() uint {
	return p.rollsAccessed
}

func (p *DepartmentProcessor) Analyze(
	section *abstractions.Department,
	rowIndex uint,
) bool {

	//fmt.Printf("Processor | Row %03d\n", section.Rows[rowIndex].Number)
	accessibleRolls := p.countAccessibleRolls(section, rowIndex)
	p.rollsAccessed += accessibleRolls

	return accessibleRolls > 0
}

func (p *DepartmentProcessor) countAccessibleRolls(
	section *abstractions.Department,
	rowIndex uint,
) uint {

	accessibleRolls := uint(0)
	rowsCount := uint(len(section.Rows)) - 1

	for spotIndex, spot := range section.Rows[rowIndex].Spots {
		if spot != abstractions.Roll {
			continue
		}

		topRowIndex := rowIndex - 1
		currentRowIndex := rowIndex
		bottomRowIndex := rowIndex + 1

		surroundingRolls := uint(0)

		/* The top row is available only if the current row is not the first one */
		if rowIndex > 0 {
			surroundingRolls += p.countRolls(&section.Rows[topRowIndex], spotIndex)
		}

		surroundingRolls += p.countRolls(&section.Rows[currentRowIndex], spotIndex) - 1 // Do not count yourself twice!

		if rowIndex < rowsCount {
			surroundingRolls += p.countRolls(&section.Rows[bottomRowIndex], spotIndex)
		}

		if surroundingRolls < minAccessibleRolls {
			//fmt.Printf("            Row %03d | Found accessible roll at spot %03d (Surrounding rolls %d) \n", section.Rows[rowIndex].Number, spotIndex, surroundingRolls)
			accessibleRolls++
			section.Rows[currentRowIndex].Spots[spotIndex] = abstractions.Empty
		}
	}

	return accessibleRolls
}

func (p *DepartmentProcessor) countRolls(
	row *abstractions.Row,
	spotIndex int,
) uint {
	spotsCount := len(row.Spots)
	rollsCount := uint(0)

	for offset := -1; offset <= 1; offset++ {
		index := spotIndex + offset
		if index < 0 || index >= spotsCount {
			continue
		}

		if row.Spots[index] == abstractions.Roll {
			rollsCount++
		}
	}

	return rollsCount
}
