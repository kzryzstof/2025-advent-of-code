package abstractions

import "fmt"

func Draw(
	tiles []*Tile,
	point *Tile,
) {
	maxX, maxY := FindFloorSize(tiles)

	for y := uint(0); y <= maxY; y++ {
		for x := uint(0); x <= maxX; x++ {
			if point != nil && point.X == x && point.Y == y {
				fmt.Print(Colorize(point.Color))
				continue
			}

			color := getTileColor(tiles, x, y)
			fmt.Print(Colorize(color))
		}
		fmt.Println()
	}
}

func FindFloorSize(
	tiles []*Tile,
) (uint, uint) {
	maxX := uint(0)
	maxY := uint(0)

	for _, tile := range tiles {
		if tile.X > maxX {
			maxX = tile.X
		}
		if tile.Y > maxY {
			maxY = tile.Y
		}
	}

	/* Adds a buffer for presentation purposes */
	return maxX + 2, maxY + 1
}

func Colorize(color string) string {
	const reset = "\033[0m"

	switch color {
	case Red:
		return "\033[31m" + color + reset // Red
	case Green:
		return "\033[32m" + color + reset // Green
	default:
		return "\033[90m" + color + reset // Gray for empty
	}
}

func getTileColor(
	tiles []*Tile,
	x uint,
	y uint,
) string {
	for _, tile := range tiles {
		if tile.X == x && tile.Y == y {
			return tile.Color
		}
	}
	return Other
}
