package maths

type Region struct {
	wide  uint
	long  uint
	ratio float64
	space [][]int8
}

func NewRegion(
	wide uint,
	long uint,
	defaultValue int8,
) *Region {
	return &Region{
		wide,
		long,
		float64(wide) / float64(long),
		NewSlice(int(long), int(wide), defaultValue),
	}
}

func (r *Region) GetSpace() [][]int8 {
	return r.space
}
