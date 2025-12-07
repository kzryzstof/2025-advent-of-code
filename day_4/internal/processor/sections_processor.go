package processor

import (
	"day_4/internal/abstractions"
	"fmt"
	"sync"
)

const (
	minAccessibleRolls = 4
)

type SectionsProcessor struct {
	sectionsChannel abstractions.SectionChannel
	syncWaitGroup   *sync.WaitGroup
	rollsAccessed   uint
}

func NewProcessor(
	sectionsChannel abstractions.SectionChannel,
	waitGroup *sync.WaitGroup,
) *SectionsProcessor {
	return &SectionsProcessor{
		sectionsChannel,
		waitGroup,
		0,
	}
}

func (p *SectionsProcessor) GetTotalAccessibleRolls() uint {
	return p.rollsAccessed
}

func (p *SectionsProcessor) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		p.monitor()
	}()
}

func (p *SectionsProcessor) monitor() {

	for section := range p.sectionsChannel.Sections() {
		accessibleRolls := p.countAccessibleRolls(section)
		p.rollsAccessed += accessibleRolls
	}
}

func (p *SectionsProcessor) countAccessibleRolls(
	section abstractions.Section,
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
			fmt.Printf("Current row %d | Spot %d | Counting rolls on row %d \n", section.Rows[currentRowIndex].Number, spotIndex, section.Rows[topRowIndex].Number)
			surroundingRolls += p.countRolls(section.Rows[topRowIndex], spotIndex)
		}

		fmt.Printf("Current row %d | Spot %d | Counting rolls on row %d\n", section.Rows[currentRowIndex].Number, spotIndex, section.Rows[currentRowIndex].Number)
		surroundingRolls += p.countRolls(section.Rows[currentRowIndex], spotIndex) - 1 // Do not count yourself twice!

		if section.RowIndex < rowsCount {
			fmt.Printf("Current row %d | Spot %d | Counting rolls on row %d\n", section.Rows[currentRowIndex].Number, spotIndex, section.Rows[bottomRowIndex].Number)
			surroundingRolls += p.countRolls(section.Rows[bottomRowIndex], spotIndex)
		}

		if surroundingRolls < minAccessibleRolls {
			accessibleRolls++
		}
	}

	return accessibleRolls
}

func (p *SectionsProcessor) countRolls(
	row abstractions.Row,
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
