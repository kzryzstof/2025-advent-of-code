package processor

import (
	"day_4/internal/abstractions"
	"fmt"
	"sync"
)

type BanksProcessor struct {
	sectionsChannel abstractions.SectionChannel
	syncWaitGroup   *sync.WaitGroup
	totalVoltage    uint
}

func NewProcessor(
	sectionsChannel abstractions.SectionChannel,
	waitGroup *sync.WaitGroup,
) *BanksProcessor {
	return &BanksProcessor{
		sectionsChannel,
		waitGroup,
		0,
	}
}

func (p *BanksProcessor) GetTotalVoltage() uint {
	return p.totalVoltage
}

func (p *BanksProcessor) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		p.monitor()
	}()
}

func (p *BanksProcessor) monitor() {

	for section := range p.sectionsChannel.Sections() {
		fmt.Println(section)
	}
}
