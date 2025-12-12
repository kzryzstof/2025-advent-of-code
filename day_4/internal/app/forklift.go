package app

import (
	"day_4/internal/abstractions"
	"fmt"
)

const (
	minAccessibleRolls = 4
)

type Forklift struct {
	verbose       bool
	rollsAccessed uint
}

func NewForklift(verbose bool) *Forklift {
	return &Forklift{
		verbose,
		0,
	}
}

func (p *Forklift) GetAccessedRollsCount() uint {
	return p.rollsAccessed
}

func (p *Forklift) RemoveRolls(
	department *abstractions.Department,
) {

	accessibleRollsFound := true
	rowsCount := uint(len(department.Rows))
	loopNumber := 1

	for accessibleRollsFound {
		/* Keeps looping until no more accessible roll is found */
		accessibleRollsFound = false

		if p.verbose {
			fmt.Printf("LOOP %d...\n", loopNumber)
		}

		for rowIndex := uint(0); rowIndex < rowsCount; rowIndex++ {

			rollsAccessed := p.countAccessibleRolls(department, rowIndex)
			p.rollsAccessed += rollsAccessed

			if rollsAccessed > 0 {
				accessibleRollsFound = true
			}
		}
		loopNumber++
	}
}

func (p *Forklift) countAccessibleRolls(
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
			if p.verbose {
				fmt.Printf("            Row %03d | Found accessible roll at spot %03d (Surrounding rolls %d) \n", section.Rows[rowIndex].Number, spotIndex, surroundingRolls)
			}
			accessibleRolls++
			section.Rows[currentRowIndex].Spots[spotIndex] = abstractions.Empty
		}
	}

	return accessibleRolls
}

func (p *Forklift) countRolls(
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
