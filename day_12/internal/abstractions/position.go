package abstractions

type Position struct {
	Row int
	Col int
}

func (p Position) Add(
	d Vector,
) Position {
	return Position{
		Row: p.Row + d.Row,
		Col: p.Col + d.Col,
	}
}
