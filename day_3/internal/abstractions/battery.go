package abstractions

type JotageRating uint8

type Battery struct {
	Rating JotageRating
	On     bool
}
