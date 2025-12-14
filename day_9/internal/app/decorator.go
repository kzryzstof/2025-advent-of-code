package app

import "day_9/internal/abstractions"

func ArrangeTiles(
	movieTheater *abstractions.MovieTheater,
) *abstractions.Rectangle {

	biggestRectangle := abstractions.NewRectangle(
		&abstractions.Tile{X: 0, Y: 0},
		&abstractions.Tile{X: 0, Y: 0},
	)

	for tileIndexA := 0; tileIndexA < len(movieTheater.RedTiles); tileIndexA++ {
		tileA := movieTheater.RedTiles[tileIndexA]

		for tileIndexB := tileIndexA + 1; tileIndexB < len(movieTheater.RedTiles); tileIndexB++ {
			tileB := movieTheater.RedTiles[tileIndexB]

			rectangle := abstractions.NewRectangle(tileA, tileB)

			if rectangle.GetArea() > biggestRectangle.GetArea() {
				biggestRectangle = rectangle
			}
		}
	}

	return biggestRectangle
}
