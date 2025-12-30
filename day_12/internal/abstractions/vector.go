package abstractions

type Vector struct {
	Row int
	Col int
}

func (d Vector) Mul(value int) Offset {
	return Offset{d.Row * value, d.Col * value}
}

func (d Vector) Opposite() Vector {
	return Vector{-d.Row, -d.Col}
}
