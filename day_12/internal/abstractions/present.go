package abstractions

type Present struct {
	shape [][]byte
	wide  uint
	long  uint
}

func NewPresent(
	shape [][]byte,
	wide uint,
	long uint,
) *Present {
	return &Present{shape, wide, long}
}
