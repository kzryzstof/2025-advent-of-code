package abstractions

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

	maxX, maxY := FindFloorSize(redTiles)

	isInjectingGreen := false

	for y := uint(0); y <= maxY; y++ {
		for x := uint(0); x <= maxX; x++ {

			isRedTile := hasTileColor(redTiles, x, y, Red)

			if isRedTile {
				/* This is a red tile, so let's add it to your new tiles */
				tiles = append(tiles, &Tile{X: x, Y: y, Color: Red})

				/* Patch to inject green tiles (issue with ray tracing on such small & gross resolution) */
				if !isInjectingGreen {
					isInjectingGreen = true
				} else {
					isInjectingGreen = false
				}

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
				} else if isInjectingGreen {
					/* Patch */
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
