package abstractions

type RotationsChannel interface {
	Rotations() <-chan Rotation
}
