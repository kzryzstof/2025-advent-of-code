package abstractions

type Cavern struct {
	presents map[uint]*Present
}

func NewCavern(
	presents map[uint]*Present,
) *Cavern {
	return &Cavern{
		presents,
	}
}

func (c *Cavern) GetPresentsCount() uint {
	return uint(len(c.presents))
}
