package maths

type Position struct {
	Row int
	Col int
}

func OriginPosition() Position {
	return Position{0, 0}
}

func (p Position) Add(
	d Vector,
) Position {
	return Position{
		Row: p.Row + d.Row,
		Col: p.Col + d.Col,
	}
}
