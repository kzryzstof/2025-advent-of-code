package abstractions

import (
	"fmt"
)

const (
	DefaultGreenTilesCapacity = 1000
)

type MovieTheater struct {
	tiles []*Tile
}

func NewMovieTheater(
	redTiles []*Tile,
) *MovieTheater {
	return &MovieTheater{
		tiles: injectGreenTiles(redTiles),
	}
}

func injectGreenTiles(
	redTiles []*Tile,
) []*Tile {

	tiles := make([]*Tile, 0, DefaultGreenTilesCapacity)

	/* The red tiles polygon must be ordered for the ray tracing to work */
	OrderPolygonVertices(redTiles)

	maxX, maxY := findFloorSize(redTiles)

	for y := uint(0); y <= maxY; y++ {
		for x := uint(0); x <= maxX; x++ {

			isRedTile := hasTileColor(redTiles, x, y, Red)

			if isRedTile {
				/* This is a red tile, so let's add it to your new tiles */
				tiles = append(tiles, &Tile{X: x, Y: y, Color: Red})
			} else {
				/* This is NOT a red tile, so let's see if it
				is inside the red tiles polygon
				*/

				isTileInside := IsPointInPolygon(
					redTiles,
					&Tile{X: x, Y: y},
				)

				if isTileInside {
					tiles = append(tiles, &Tile{X: x, Y: y, Color: Green})
				}
			}

		}
	}

	return tiles
}

func (mt *MovieTheater) GetTiles() []*Tile {
	return mt.tiles
}

func findFloorSize(
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

func (mt *MovieTheater) getTileColor(
	x uint,
	y uint,
) string {
	for _, tile := range mt.tiles {
		if tile.X == x && tile.Y == y {
			return tile.Color
		}
	}
	return Other
}

func hasTileColor(
	tiles []*Tile,
	x uint,
	y uint,
	color string,
) bool {
	for _, tile := range tiles {
		if tile.X == x && tile.Y == y && tile.Color == color {
			return true
		}
	}
	return false
}

func (mt *MovieTheater) Draw() {
	maxX, maxY := findFloorSize(mt.tiles)

	for y := uint(0); y <= maxY; y++ {
		for x := uint(0); x <= maxX; x++ {
			color := mt.getTileColor(x, y)
			fmt.Print(colorize(color))
		}
		fmt.Println()
	}
}

func colorize(color string) string {
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
