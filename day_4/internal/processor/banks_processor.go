package processor

import (
	"day_4/internal/abstractions"
	"sync"
)

type BanksProcessor struct {
	banksChannel  abstractions.BanksChannel
	syncWaitGroup *sync.WaitGroup
	totalVoltage  uint
}

func NewProcessor(
	banksChannel abstractions.BanksChannel,
	waitGroup *sync.WaitGroup,
) *BanksProcessor {
	return &BanksProcessor{
		banksChannel,
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

	for bank := range p.banksChannel.Banks() {
		p.totalVoltage += bank.GetHighestVoltage()
	}
}
