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
	space := make([][]int8, long)

	for row := uint(0); row < long; row++ {
		space[row] = make([]int8, wide)
		for col := uint(0); col < wide; col++ {
			space[row][col] = int8(0)
		}
	}
	return &Region{
		wide,
		long,
		float64(wide) / float64(long),
		space,
	}
}

func (r *Region) GetArea() uint {
	return r.wide * r.long
}

func (r *Region) GetRatio() float64 {
	return r.ratio
}

func (r *Region) GetSpace() [][]int8 {
	return r.space
}
