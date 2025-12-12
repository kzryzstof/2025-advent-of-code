package abstractions

type Position struct {
	RowIndex int
	ColIndex int
}

func ChangePosition(
	position Position,
	direction Direction,
) Position {
	return Position{
		RowIndex: position.RowIndex + direction.RowDelta,
		ColIndex: position.ColIndex + direction.ColDelta,
	}
}
