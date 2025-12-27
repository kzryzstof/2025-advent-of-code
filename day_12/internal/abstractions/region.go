package abstractions

type Region struct {
	wide  uint
	long  uint
	ratio float64
	space [][]byte
}

func NewRegion(
	wide uint,
	long uint,
) *Region {
	space := make([][]byte, long)
	for row := uint(0); row < long; row++ {
		space[row] = make([]byte, wide)
		for col := uint(0); col < wide; col++ {
			space[row][col] = byte(0)
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
