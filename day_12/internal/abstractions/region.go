package abstractions

type Region struct {
	wide  uint
	long  uint
	ratio float64
	space [][]int8
}

func NewRegion(
	wide uint,
	long uint,
) *Region {
	return &Region{
		wide,
		long,
		float64(wide) / float64(long),
		NewSlice(int(long), int(wide)),
	}
}

func (r *Region) GetSpace() [][]int8 {
	return r.space
}
