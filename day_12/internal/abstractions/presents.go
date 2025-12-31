package abstractions

const (
	// MaximumShapeSize /* All the presents occupies a 3x3 region */
	MaximumShapeSize = 3
)

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

func (p *Presents) GetAllPresents() <-chan *Present {
	ch := make(chan *Present)

	go func() {
		defer close(ch)
		for _, pr := range p.presents {
			ch <- pr
		}
	}()

	return ch
}

func (p *Presents) GetPresent(
	index uint,
) *Present {
	/* Returns a copy of the present that can be manipulated */
	return p.presents[index]
}
