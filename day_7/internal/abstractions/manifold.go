package abstractions

import (
	"fmt"
	"os"
)

const (
	Empty         string = "."
	Splitter             = "^"
	Beam                 = "|"
	StartingPoint        = "S"
	Invalid              = " "
)

type Manifold struct {
	Locations   [][]string
	Tachyons    []*Tachyon
	splitsCount uint
	/* Lists all the splitters encountered in the manifold sorted by RowIndex, ColIndex
	each time one is encountered, timelines are created */
	timelines [][]uint64
}

func NewManifold(
	locations [][]string,
	tachyons []*Tachyon,
) *Manifold {
	return &Manifold{
		Locations: locations,
		Tachyons:  tachyons,
		timelines: make([][]uint64, len(locations)),
	}
}

func (m *Manifold) GetNextPosition(
	tachyon *Tachyon,
	direction Direction,
) string {
	newPosition := tachyon.Position.MoveTo(direction)
	return m.GetLocation(newPosition)
}

func (m *Manifold) CanMove(
	t Tachyon,
	direction Direction,
) bool {
	return m.isWithinBoundary(
		t.Position,
		direction,
	)
}

func (m *Manifold) IsNextLocationEmpty(
	t Tachyon,
	direction Direction,
) bool {
	nextLocation := m.GetNextPosition(&t, direction)
	return nextLocation == Empty
}

func (m *Manifold) IsNextLocationSplitter(
	t Tachyon,
	direction Direction,
) bool {
	nextLocation := m.GetNextPosition(&t, direction)
	return nextLocation == Splitter
}

func (m *Manifold) IsNextLocationOtherTachyon(
	t Tachyon,
	direction Direction,
) bool {
	nextLocation := m.GetNextPosition(&t, direction)
	return nextLocation == Beam
}

func (m *Manifold) Merge(
	t *Tachyon,
	direction Direction,
) {
	/* Merges the specified tachyon with the one at the next position */
	nextPosition := t.Position.MoveTo(direction)

	existingTachyon := m.GetTachyonAt(nextPosition)

	if existingTachyon == nil {
		fmt.Println("No existing tachyon found at position", t.Position)
		os.Exit(1)
	}

	/* Merges the two tachyons (and keep track of the "freaking" beams) */
	t.MergeTo(existingTachyon)

	/* No need to track the position of this tachyon anymore */
	t.Stop()
}

func (m *Manifold) GetLocation(
	position Position,
) string {

	if position.RowIndex < 0 || position.RowIndex >= len(m.Locations) {
		return Invalid
	}

	if position.ColIndex < 0 || position.ColIndex >= len(m.Locations[position.RowIndex]) {
		return Invalid
	}

	return m.Locations[position.RowIndex][position.ColIndex]
}

func (m *Manifold) isWithinBoundary(
	position Position,
	direction Direction,
) bool {

	newPosition := position.MoveTo(direction)

	if newPosition.RowIndex < 0 || newPosition.RowIndex >= len(m.Locations) {
		return false
	}

	if newPosition.ColIndex < 0 || newPosition.ColIndex >= len(m.Locations[newPosition.RowIndex]) {
		return false
	}

	return true
}

func (m *Manifold) Draw() {
	for _, row := range m.Locations {
		for _, location := range row {
			print(location)
		}
		println()
	}
}

func (m *Manifold) SetBeamAt(
	position Position,
	direction Direction,
) (bool, Position) {

	if !m.isWithinBoundary(position, direction) {
		return false, position
	}

	newPosition := position.MoveTo(direction)

	m.Locations[newPosition.RowIndex][newPosition.ColIndex] = Beam
	return true, newPosition
}

func (m *Manifold) SplitBeamAt(
	tachyon *Tachyon,
	direction Direction,
) {

	m.splitsCount++

	/* Loki split the timeline. Again! */
	splitPosition := tachyon.Position.MoveTo(direction)
	m.createTimeline(*tachyon, splitPosition)

	/* Adds a new tachyon and move it to the right if there is no other tachyon */
	newTachyonDirection := Direction{
		RowDelta: direction.RowDelta,
		ColDelta: direction.ColDelta + 1,
	}

	newTachyon := m.splitTachyon(tachyon)

	if !m.IsNextLocationOtherTachyon(*tachyon, newTachyonDirection) {
		newTachyon.Move(
			m,
			newTachyonDirection,
		)
	} else {
		/* There is already a tachyon at this position? No biggies: merge with it! */
		m.Merge(newTachyon, newTachyonDirection)
	}

	/* Takes the existing tachyon and move it to the left */

	existingTachyonDirection := Direction{
		RowDelta: direction.RowDelta,
		ColDelta: direction.ColDelta - 1,
	}

	if !m.IsNextLocationOtherTachyon(*tachyon, existingTachyonDirection) {
		tachyon.Move(
			m,
			existingTachyonDirection,
		)
	} else {
		/* There is already a tachyon at this position as well? Again, no biggies: merge with it! */
		m.Merge(tachyon, existingTachyonDirection)
	}
}

func (m *Manifold) splitTachyon(
	tachyon *Tachyon,
) *Tachyon {

	/* Creates a new tachyon out the existing one (with all the properties) */
	newTachyon := tachyon.Split()

	m.Tachyons = append(m.Tachyons, newTachyon)

	return newTachyon
}

func (m *Manifold) AreTachyonsMoving() bool {

	movingTachyons := false

	for _, tachyon := range m.Tachyons {
		movingTachyons = movingTachyons || tachyon.IsMoving()
	}

	return movingTachyons
}

func (m *Manifold) createTimeline(
	tachyon Tachyon,
	position Position,
) {
	/* A new MCU timeline is created each time a splitter is encountered */

	if m.timelines[position.RowIndex] == nil {
		m.timelines[position.RowIndex] = make([]uint64, len(m.Locations[position.RowIndex]))
	}

	/* Important: note how we know how many beams hit the splitter at the same
	while we only simulate the trajectory of one tachyon */
	m.timelines[position.RowIndex][position.ColIndex] += tachyon.GetMergedBeams()
}

func (m *Manifold) CountTimelines() uint64 {

	/* We start with our timeline, the sacred one */
	totalTimelines := uint64(1)

	for _, timeline := range m.timelines {

		if timeline == nil {
			continue
		}

		for _, beamsCount := range timeline {
			if beamsCount == 0 {
				continue
			}

			/* Each time a beam hit the splitter, a new timeline is created */
			totalTimelines += beamsCount
		}
	}

	return totalTimelines
}

func (m *Manifold) GetTachyonAt(
	position Position,
) *Tachyon {
	for _, tachyon := range m.Tachyons {
		if tachyon.Position == position {
			return tachyon
		}
	}

	return nil
}
