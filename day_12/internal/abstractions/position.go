package abstractions

import "math"

type Position struct {
	Row int
	Col int
}

func OriginPosition() Position {
	return Position{Row: 0, Col: 0}
}

func (p Position) AddPosition(
	other Position,
) Position {
	return Position{
		Row: p.Row + other.Row,
		Col: p.Col + other.Col,
	}
}

func (p Position) SubPosition(
	other Position,
) Position {
	return Position{
		Row: p.Row - other.Row,
		Col: p.Col - other.Col,
	}
}

func (p Position) Add(
	d Vector,
) Position {
	return Position{
		Row: p.Row + d.Row,
		Col: p.Col + d.Col,
	}
}

func (p Position) Sub(
	d Vector,
) Position {
	return Position{
		Row: p.Row - d.Row,
		Col: p.Col - d.Col,
	}
}

func (p Position) Mul(
	d Vector,
) Position {
	return Position{
		Row: p.Row * d.Row,
		Col: p.Col * d.Col,
	}
}

func (p Position) Offset(
	offset int,
	d Vector,
) Position {
	return Position{
		Row: p.Row + (offset-1)*d.Row,
		Col: p.Col + (offset-1)*d.Col,
	}
}

func (p Position) GetDistanceTo(
	other Position,
) float64 {
	rowDistance := float64(p.Row - other.Row)
	columnDistance := float64(p.Col - other.Col)
	return math.Sqrt(rowDistance*rowDistance + columnDistance*columnDistance)
}

func (p Position) Equals(
	otherPosition Position,
) bool {
	return p.Row == otherPosition.Row && p.Col == otherPosition.Col
}

func FindMinCol(positions []Position) int {

	if len(positions) == 0 {
		return -1
	}

	minCol := math.MaxInt32

	for _, position := range positions {
		minCol = int(math.Min(float64(minCol), float64(position.Col)))
	}

	return minCol
}

func FindMinRow(positions []Position) int {

	if len(positions) == 0 {
		return -1
	}

	minRow := math.MaxInt32

	for _, position := range positions {
		minRow = int(math.Min(float64(minRow), float64(position.Row)))
	}

	return minRow
}

func FindMaxCol(positions []Position) int {

	if len(positions) == 0 {
		return -1
	}

	maxCol := math.MinInt32

	for _, position := range positions {
		maxCol = int(math.Max(float64(maxCol), float64(position.Col)))
	}

	return maxCol
}

func FindMaxRow(positions []Position) int {

	if len(positions) == 0 {
		return -1
	}

	maxRow := math.MinInt32

	for _, position := range positions {
		maxRow = int(math.Max(float64(maxRow), float64(position.Row)))
	}

	return maxRow
}
