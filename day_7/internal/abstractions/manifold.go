package abstractions

type Location string

const (
	Empty             Location = "."
	Splitter                   = '^'
	BeamStartingPoint          = 'S'
)

type Manifold struct {
	Locations [][]Location
}
