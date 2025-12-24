package abstractions

type ChristmasTree struct {
	wide uint
	long uint
	/* Contains number of presents of a certain index */
	presentIndices map[uint]uint
	Region         *Region
}

func NewChristmasTree(
	wide uint,
	long uint,
	presentIndices map[uint]uint,
) *ChristmasTree {
	return &ChristmasTree{
		wide,
		long,
		presentIndices,
		NewRegion(wide, long),
	}
}

func (ct *ChristmasTree) GetPresents() map[uint]uint {
	return ct.presentIndices
}

func (ct *ChristmasTree) GetRegion() [][]byte {
	region := make([][]byte, ct.long)
	for row := uint(0); row < ct.long; row++ {
		region[row] = make([]byte, ct.wide)
		for col := uint(0); col < ct.wide; col++ {
			region[row][col] = byte(0)
		}
	}
	return region
}
