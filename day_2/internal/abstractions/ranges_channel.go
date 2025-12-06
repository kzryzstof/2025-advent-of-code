package abstractions

type RangesChannel interface {
	Ranges() <-chan Range
}
