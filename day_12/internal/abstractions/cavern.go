package abstractions

type Cavern struct {
	presents       map[uint]*Present
	christmasTrees []*ChristmasTree
}

func NewCavern(
	presents map[uint]*Present,
	christmasTrees []*ChristmasTree,
) *Cavern {
	return &Cavern{
		presents,
		christmasTrees,
	}
}

func (c *Cavern) GetPresentsCount() uint {
	return uint(len(c.presents))
}

func (c *Cavern) GetChristmasTreesCount() uint {
	return uint(len(c.christmasTrees))
}
