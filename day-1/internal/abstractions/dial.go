package abstractions

const (
	MinimalDialPosition int = 0
	MaximalDialPosition int = 99
	PositionsCount      int = MaximalDialPosition - MinimalDialPosition + 1
)

type Dial struct {
	count    int
	Position int
}

func (d *Dial) GetCount() int {
	return d.count
}

func (d *Dial) Rotate(
	rotation Rotation,
) {
	d.turn(rotation.Direction, rotation.Distance)
}

func (d *Dial) turn(
	direction Direction,
	distance int,
) {
	var newPosition int

	/* Handle distances larger than the dial size by capping them */
	cappedDistance := d.getCappedDistance(distance)

	if direction == Left {
		newPosition = d.turnLeft(cappedDistance)
	} else {
		newPosition = d.turnRight(cappedDistance)
	}

	d.Position = newPosition

	d.incrementCountIfNeeded()
}

func (d *Dial) getCappedDistance(
	distance int,
) int {
	return distance % PositionsCount
}

func (d *Dial) turnLeft(
	distance int,
) int {
	// Turning left decreases the position, wrapping in [0, MaximalPosition].
	// Example: 0 left 1 -> 99; 5 left 10 -> 95.
	// Ensure positive before modulo to avoid negative results
	return (d.Position - distance + PositionsCount) % PositionsCount
}

func (d *Dial) turnRight(
	distance int,
) int {
	return (d.Position + distance) % PositionsCount
}

func (d *Dial) incrementCountIfNeeded() {
	if d.Position == 0 {
		d.count++
	}
}
