package abstractions

type Position struct {
	Row int
	Col int
}

func OriginPosition() Position {
	return Position{Row: 0, Col: 0}
}

var PositionNotFound = Position{-1, -1}

func (p Position) Add(
	d Direction,
) Position {
	return Position{
		Row: p.Row + d.Row,
		Col: p.Col + d.Col,
	}
}

func (p Position) Sub(
	d Direction,
) Position {
	return Position{
		Row: p.Row - d.Row,
		Col: p.Col - d.Col,
	}
}

func (p Position) Offset(
	offset int,
	d Direction,
) Position {
	return Position{
		Row: p.Row + (offset-1)*d.Row,
		Col: p.Col + (offset-1)*d.Col,
	}
}

func (p Position) Equals(
	otherPosition Position,
) bool {
	return p.Row == otherPosition.Row && p.Col == otherPosition.Col
}
