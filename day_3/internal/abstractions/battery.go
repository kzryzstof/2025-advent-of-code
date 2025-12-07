package abstractions

type VoltageRating uint

type Battery struct {
	Voltage VoltageRating
	On      bool
}
