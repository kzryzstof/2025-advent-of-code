package abstractions

type Present struct {
	index uint
	shape [][]byte
	wide  uint
	long  uint
}

func NewPresent(
	index uint,
	shape [][]byte,
	wide uint,
	long uint,
) *Present {
	return &Present{index, shape, wide, long}
}

func (p *Present) GetShape() [][]byte {
	return GetCopy(p.shape)
}

func (p *Present) GetIndex() uint {
	return p.index
}

func (p *Present) GetArea() uint {
	return p.wide * p.long
}
