package processor

import (
	"day_2/internal/abstractions"
	"fmt"
	"sync"
)

type RangesProcessor struct {
	rangesChannel abstractions.RangesChannel
	syncWaitGroup *sync.WaitGroup
}

func New(
	rangesChannel abstractions.RangesChannel,
	waitGroup *sync.WaitGroup,
) *RangesProcessor {
	return &RangesProcessor{
		rangesChannel,
		waitGroup,
	}
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
		fmt.Printf("Range: %v\n", r)
	}
}
