package abstractions

type Spot int32

const (
	None  Spot = 0
	Empty Spot = '.'
	Roll  Spot = '@'
	Fork  Spot = 'x'
)

type Row struct {
	Number int
	Spots  []Spot
}
