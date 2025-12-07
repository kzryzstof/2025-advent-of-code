package processor

import (
	"day_3/internal/abstractions"
	"fmt"
	"sync"
)

type BanksProcessor struct {
	banksChannel      abstractions.BanksChannel
	syncWaitGroup     *sync.WaitGroup
	invalidProductIds []int64
	totalProductId    int64
}

func NewProcessor(
	banksChannel abstractions.BanksChannel,
	waitGroup *sync.WaitGroup,
) *BanksProcessor {
	return &BanksProcessor{
		banksChannel,
		waitGroup,
		make([]int64, 0),
		0,
	}
}

func (p *BanksProcessor) GetTotalProductId() int64 {
	return p.totalProductId
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
		fmt.Printf("Processing range: %w to %s\n", bank)

	}
}
