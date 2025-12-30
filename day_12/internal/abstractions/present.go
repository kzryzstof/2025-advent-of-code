package abstractions

type Present struct {
	index uint
	shape Shape
}

func NewPresent(
	index uint,
	shape Shape,
) *Present {
	return &Present{index, shape}
}

func (p *Present) GetShape() [][]int8 {
	return p.shape.GetCopy()
}

func (p *Present) GetIndex() uint {
	return p.index
}

func (p *Present) GetArea() uint {
	return p.shape.Dimension.GetArea()
}
