package abstractions

type ChristmasTree struct {
	wide uint
	long uint
	/* Contains number of presents of a certain index */
	presentIndices map[uint]uint
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
	}
}

func (ct *ChristmasTree) TryFit(
	presents *Presents,
) bool {
	return false
}
