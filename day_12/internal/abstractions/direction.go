package abstractions

type Direction struct {
	Row int
	Col int
}

func (d Direction) Mul(value int) Offset {
	return Offset{d.Row * value, d.Col * value}
}

func (d Direction) Opposite() Direction {
	return Direction{-d.Row, -d.Col}
}
