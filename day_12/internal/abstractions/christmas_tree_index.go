package abstractions

type ChristmasTreeIndex uint

func (pi ChristmasTreeIndex) AsUint() uint {
	return uint(pi)
}
