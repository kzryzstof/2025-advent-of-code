package abstractions

type Position struct {
	Row int
	Col int
}

func OriginPosition() Position {
	return Position{Row: 0, Col: 0}
}

func (p Position) AddPosition(
	other Position,
) Position {
	return Position{
		Row: p.Row + other.Row,
		Col: p.Col + other.Col,
	}
}

func (p Position) Add(
	d Vector,
) Position {
	return Position{
		Row: p.Row + d.Row,
		Col: p.Col + d.Col,
	}
}
