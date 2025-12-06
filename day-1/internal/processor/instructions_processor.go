package processor

import (
	"day-1/internal/abstractions"
	"sync"
)

type InstructionsProcessor struct {
	rotationsChannel abstractions.RotationsChannel
	syncWaitGroup    *sync.WaitGroup
}

func New(
	rotationsChannel abstractions.RotationsChannel,
	waitGroup *sync.WaitGroup,
) *InstructionsProcessor {
	return &InstructionsProcessor{
		rotationsChannel,
		waitGroup,
	}
}

func (p *InstructionsProcessor) Start(
	dial *abstractions.Dial,
) {

	p.syncWaitGroup.Add(1)

	go func() {

		defer p.syncWaitGroup.Done()

		p.monitor(dial)
	}()
}

func (p *InstructionsProcessor) monitor(
	dial *abstractions.Dial,
) {

	for rotation := range p.rotationsChannel.Rotations() {
		dial.Rotate(rotation)
	}
}
