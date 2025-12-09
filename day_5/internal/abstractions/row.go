package abstractions

type Spot int

const (
	None  Spot = 0
	Empty Spot = '.'
	Roll  Spot = '@'
)

type Row struct {
	Number uint
	Spots  []Spot
}
