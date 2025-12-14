package abstractions

import (
	"fmt"
)

type MovieTheater struct {
	redTiles []*Tile
	/* Contains all the valid tiles (green and red) */
	validTiles [][]bool
}

func NewMovieTheater(
	redTiles []*Tile,
) *MovieTheater {
	return &MovieTheater{
		redTiles:   redTiles,
		validTiles: preComputeValidTiles(redTiles),
	}
}

func preComputeValidTiles(
	redTiles []*Tile,
) [][]bool {

	maxX, maxY := FindFloorSize(redTiles)

	fmt.Printf("Floor size is %d x %d\n", maxX, maxY)

	/* Initializes the valid tiles map: this is a big one so let's allocate everything upfront */
	validTiles := make([][]bool, maxY)

	for y := uint(0); y < maxY; y++ {
		validTiles[y] = make([]bool, maxX)
	}

	fmt.Print("Memory allocated\n")

	/* The red tiles polygon must be ordered for the ray tracing to work */
	OrderPolygonVertices(redTiles)

	fmt.Print("Vertices ordered\n")

	fmt.Print("Building pre-validated tiles\n")

	processedRows := 0

	for y := uint(0); y < maxY; y++ {

		fmt.Printf("\t%05d/%05d processed rows\r", processedRows, maxY)

		for x := uint(0); x < maxX; x++ {

			isRedTile := isRedTile(redTiles, x, y)

			if isRedTile {

				/* This is a red tile, so let's add it to your new tiles */
				validTiles[y][x] = true

			} else {

				/* This is NOT a red tile, so let's see if it is inside the red tiles polygon */
				validTiles[y][x] = IsPointInPolygon(
					redTiles,
					&Tile{X: x, Y: y},
				)
			}
		}

		processedRows++
	}

	return validTiles
}

func (mt *MovieTheater) GetRedTiles() []*Tile {
	return mt.redTiles
}

func isRedTile(
	tiles []*Tile,
	x uint,
	y uint,
) bool {
	for _, tile := range tiles {
		if tile.X == x && tile.Y == y {
			return true
		}
	}
	return false
}

func (mt *MovieTheater) IsValidTile(
	x uint,
	y uint,
) bool {
	if y >= uint(len(mt.validTiles)) {
		return false
	}
	if x >= uint(len(mt.validTiles[y])) {
		return false
	}
	return mt.validTiles[y][x]
}
