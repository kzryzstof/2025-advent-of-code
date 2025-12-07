package processor

import (
	"day_4/internal/abstractions"
	"fmt"
)

const (
	minAccessibleRolls = 4
)

type SectionsProcessor struct {
	rollsAccessed uint
}

func NewProcessor() *SectionsProcessor {
	return &SectionsProcessor{
		0,
	}
}

func (p *SectionsProcessor) GetTotalAccessibleRolls() uint {
	return p.rollsAccessed
}

func (p *SectionsProcessor) Analyze(section *abstractions.Section) {

	fmt.Printf("Processor | Row %03d | Received\n", section.Rows[section.RowIndex].Number)
	accessibleRolls := p.countAccessibleRolls(section)
	p.rollsAccessed += accessibleRolls
	fmt.Printf("Processor | Rows %03d | Pocessed\n", section.Rows[section.RowIndex].Number)
}

func (p *SectionsProcessor) countAccessibleRolls(
	section *abstractions.Section,
) uint {

	accessibleRolls := uint(0)
	rowsCount := len(section.Rows) - 1

	for spotIndex, spot := range section.Rows[section.RowIndex].Spots {
		if spot != abstractions.Roll {
			continue
		}

		topRowIndex := section.RowIndex - 1
		currentRowIndex := section.RowIndex
		bottomRowIndex := section.RowIndex + 1

		surroundingRolls := uint(0)

		if topRowIndex >= 0 {
			surroundingRolls += p.countRolls(&section.Rows[topRowIndex], spotIndex)
		}

		surroundingRolls += p.countRolls(&section.Rows[currentRowIndex], spotIndex) - 1 // Do not count yourself twice!

		if section.RowIndex < rowsCount {
			surroundingRolls += p.countRolls(&section.Rows[bottomRowIndex], spotIndex)
		}

		if surroundingRolls < minAccessibleRolls {
			fmt.Printf("            Row %03d | Found accessible roll at spot %03d (Surrounding rolls %d) \n", section.Rows[section.RowIndex].Number, spotIndex, surroundingRolls)
			accessibleRolls++
		}
	}

	return accessibleRolls
}

func (p *SectionsProcessor) countRolls(
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
