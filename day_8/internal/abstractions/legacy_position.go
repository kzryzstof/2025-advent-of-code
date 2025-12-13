package abstractions

type LegacyPosition struct {
	RowIndex int
	ColIndex int
}

func (p LegacyPosition) MoveTo(
	direction Direction,
) LegacyPosition {
	return LegacyPosition{
		RowIndex: p.RowIndex + direction.RowDelta,
		ColIndex: p.ColIndex + direction.ColDelta,
	}
}
