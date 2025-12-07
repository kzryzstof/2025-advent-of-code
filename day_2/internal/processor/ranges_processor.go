package processor

import (
	"day_2/internal/abstractions"
	"fmt"
	"sync"
)

type RangesProcessor struct {
	rangesChannel     abstractions.RangesChannel
	syncWaitGroup     *sync.WaitGroup
	invalidProductIds []int64
	totalProductId    int64
}

func NewProcessor(
	rangesChannel abstractions.RangesChannel,
	waitGroup *sync.WaitGroup,
) *RangesProcessor {
	return &RangesProcessor{
		rangesChannel,
		waitGroup,
		make([]int64, 0),
		0,
	}
}

func (p *RangesProcessor) GetTotalProductId() int64 {
	return p.totalProductId
}

func (p *RangesProcessor) Start() {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		p.monitor()
	}()
}

func (p *RangesProcessor) monitor() {

	for r := range p.rangesChannel.Ranges() {
		fmt.Printf("Processing range: %s to %s\n", r.From.Id, r.To.Id)

		invalidProductIds := r.FindInvalidProductIds()

		if len(invalidProductIds) == 0 {
			fmt.Printf("\tNo invalid product IDs found\n")
			continue
		}

		for _, invalidProductId := range invalidProductIds {
			p.invalidProductIds = append(p.invalidProductIds, invalidProductId)
			p.totalProductId += invalidProductId
		}
	}
}
