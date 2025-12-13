package abstractions

type Tachyon struct {
	Position Position
	isMoving bool
	/* Tracks how many beams are actually following the same trajectory */
	mergedBeams uint64
}

func NewTachyon(
	position Position,
) *Tachyon {
	return &Tachyon{
		Position:    position,
		isMoving:    true,
		mergedBeams: 1,
	}
}

func (t *Tachyon) Split() *Tachyon {
	return &Tachyon{
		Position:    t.Position,
		isMoving:    t.isMoving,
		mergedBeams: t.mergedBeams,
	}
}

func (t *Tachyon) GetMergedBeams() uint64 {
	return t.mergedBeams
}

func (t *Tachyon) MergeTo(other *Tachyon) {
	other.mergedBeams += t.mergedBeams
}

func (t *Tachyon) IsMoving() bool {
	return t.isMoving
}

func (t *Tachyon) Stop() {
	t.isMoving = false
}

func (t *Tachyon) Move(
	manifold *Manifold,
	direction Direction,
) {

	result, newPosition := manifold.SetBeamAt(
		t.Position,
		direction,
	)

	if result == true {
		t.Position = newPosition
	}
}
