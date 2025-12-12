package abstractions

const (
	Empty         string = "."
	Splitter             = "^"
	Beam                 = "|"
	StartingPoint        = "S"
	Invalid              = " "
)

type Manifold struct {
	Locations [][]string
	Tachyons  []*Tachyon
}

func (m *Manifold) GetNextPosition(
	tachyon *Tachyon,
	direction Direction,
) string {
	newPosition := ChangePosition(tachyon.Position, direction)
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

	newPosition := ChangePosition(position, direction)

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

/* Beam functionalities */

func (m *Manifold) SetBeamAt(
	position Position,
	direction Direction,
) (bool, Position) {

	if !m.isWithinBoundary(position, direction) {
		return false, position
	}

	newPosition := ChangePosition(position, direction)

	m.Locations[newPosition.RowIndex][newPosition.ColIndex] = Beam
	return true, newPosition
}

func (m *Manifold) SplitBeamAt(
	tachyon *Tachyon,
	direction Direction,
) {

	/* Adds a new tachyon and move it to the right if there is no other tachyon */

	newTachyonDirection := Direction{
		RowDelta: direction.RowDelta,
		ColDelta: direction.ColDelta + 1,
	}

	if !m.IsNextLocationOtherTachyon(*tachyon, newTachyonDirection) {

		newTachyon := m.createNewTachyon(tachyon)

		newTachyon.Move(
			m,
			newTachyonDirection,
		)
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
	}
}

func (m *Manifold) createNewTachyon(
	tachyon *Tachyon,
) *Tachyon {

	newTachyon := Tachyon{
		Position: tachyon.Position,
	}

	newTachyon.Start()

	m.Tachyons = append(m.Tachyons, &newTachyon)

	return &newTachyon
}
