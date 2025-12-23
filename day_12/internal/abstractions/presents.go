package abstractions

type Presents struct {
	presents map[uint]*Present
}

func NewPresents(
	presents map[uint]*Present,
) *Presents {
	return &Presents{
		presents,
	}
}

func (p *Presents) Count() uint {
	return uint(len(p.presents))
}
