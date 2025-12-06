package abstractions

import "math"

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
	var initialPosition int = d.Position
	var newPosition int

	/* Handle distances larger than the dial size by capping them */
	cappedDistance := d.getCappedDistance(distance)

	if direction == Left {
		newPosition = d.turnLeft(cappedDistance)
	} else {
		newPosition = d.turnRight(cappedDistance)
	}

	d.Position = newPosition

	d.incrementCount(
		direction,
		distance,
		initialPosition,
	)
}

func (d *Dial) countFullTurns(
	distance int,
) int {
	return int(math.Floor(float64(distance) / float64(PositionsCount)))
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

func (d *Dial) incrementCount(
	direction Direction,
	distance int,
	initialPosition int,
) {
	//	Counts all the full turns.
	fullTurns := d.countFullTurns(distance)

	//	Counts the remaining potential partial turn i.e., when we cross the initial position without doing a full turn.
	crossedZero := false

	if initialPosition != MinimalDialPosition {
		//	Only count the case when we are not at the initial position 0.
		if direction == Left && d.Position > initialPosition {
			//	Counts the case when we turn left and cross the initial position
			crossedZero = true
		} else if direction == Right && d.Position < initialPosition {
			// 	Counts the case when we turn right and cross the initial position
			crossedZero = true
		} else if d.Position == MinimalDialPosition {
			// 	Count the case when we do end up at the initial position
			crossedZero = true
		}
	}

	if crossedZero {
		fullTurns += 1
	}

	d.count += fullTurns
}
