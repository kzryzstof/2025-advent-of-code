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

func (m *Manifold) Draw() {
	for _, row := range m.Locations {
		for _, location := range row {
			print(location)
		}
		println()
	}
}
