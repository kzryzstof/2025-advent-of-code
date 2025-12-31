package abstractions

type Present struct {
	index PresentIndex
	shape Shape
}

func NewPresent(
	index PresentIndex,
	shape Shape,
) *Present {
	return &Present{index, shape}
}

func (p *Present) GetShape() [][]int8 {
	return p.shape.GetCopy()
}

func (p *Present) GetIndex() PresentIndex {
	return p.index
}
