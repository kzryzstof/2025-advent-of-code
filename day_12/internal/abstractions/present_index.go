package abstractions

import "math"

type PresentIndex uint

func NoPresentIndex() PresentIndex {
	return PresentIndex(math.MaxUint)
}

func (pi PresentIndex) AsUint() uint {
	return uint(pi)
}
