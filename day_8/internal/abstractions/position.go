package abstractions

type Position struct {
	RowIndex int
	ColIndex int
}

func (p Position) MoveTo(
	direction Direction,
) Position {
	return Position{
		RowIndex: p.RowIndex + direction.RowDelta,
		ColIndex: p.ColIndex + direction.ColDelta,
	}
}
