package processor

import (
	"day_4/internal/abstractions"
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

		totalRolls := p.countRolls(section.Rows[section.RowIndex-1], spotIndex)
		totalRolls += p.countRolls(section.Rows[section.RowIndex], spotIndex) - 1 // Do not count yourself twice!

		if section.RowIndex < rowsCount {
			totalRolls += p.countRolls(section.Rows[section.RowIndex+1], spotIndex)
		}

		if totalRolls < minAccessibleRolls {
			accessibleRolls++
		}
	}

	return accessibleRolls
}

func (p *SectionsProcessor) countRolls(row abstractions.Row, spotIndex int) uint {
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

	return uint(rollsCount)
}
