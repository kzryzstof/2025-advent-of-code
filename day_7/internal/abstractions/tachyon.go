package abstractions

type Tachyon struct {
	Position Position
	isMoving bool
}

func (t *Tachyon) IsMoving() bool {
	return t.isMoving
}

func (t *Tachyon) Start() {
	t.isMoving = true
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
