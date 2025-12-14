package abstractions

const (
	Green = "X"
	Red   = "#"
	Other = "."
)

type Tile struct {
	X     uint
	Y     uint
	Color string
}
