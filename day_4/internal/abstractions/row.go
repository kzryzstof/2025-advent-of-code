package abstractions

type Spot string

const (
	Empty Spot = "."
	Roll  Spot = "@"
	Fork  Spot = "x"
)

type Row struct {
	Spots []Spot
}
