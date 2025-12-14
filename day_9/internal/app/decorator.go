package app

import (
	"day_9/internal/abstractions"
	"fmt"
)

func ArrangeTiles(
	movieTheater *abstractions.MovieTheater,
) *abstractions.Rectangle {

	biggestRectangle := abstractions.NewRectangle(
		&abstractions.Tile{X: 0, Y: 0},
		&abstractions.Tile{X: 0, Y: 0},
	)

	redTiles := movieTheater.GetRedTiles()

	for tileIndexA := 0; tileIndexA < len(redTiles); tileIndexA++ {
		tileA := redTiles[tileIndexA]

		for tileIndexB := tileIndexA + 1; tileIndexB < len(redTiles); tileIndexB++ {
			tileB := redTiles[tileIndexB]

			fmt.Printf("Testing red tile [%05d,%05d] (%d/%d) with other red tile [%05d,%05d] (%d/%d) \r", tileA.X, tileA.Y, tileIndexA, len(redTiles), tileB.X, tileB.Y, tileIndexB, len(redTiles))

			rectangle := abstractions.NewRectangle(tileA, tileB)

			if rectangle.GetArea() > biggestRectangle.GetArea() {

				if !rectangle.IsInside(movieTheater) {
					continue
				}
				biggestRectangle = rectangle
			}
		}
	}

	return biggestRectangle
}
