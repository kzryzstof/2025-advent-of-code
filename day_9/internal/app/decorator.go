package app

import "day_9/internal/abstractions"

func ArrangeTiles(
	movieTheater *abstractions.MovieTheater,
) *abstractions.Rectangle {

	movieTheater.Draw()

	biggestRectangle := abstractions.NewRectangle(
		&abstractions.Tile{X: 0, Y: 0},
		&abstractions.Tile{X: 0, Y: 0},
	)

	tiles := movieTheater.GetTiles()

	for tileIndexA := 0; tileIndexA < len(tiles); tileIndexA++ {
		tileA := tiles[tileIndexA]

		for tileIndexB := tileIndexA + 1; tileIndexB < len(tiles); tileIndexB++ {
			tileB := tiles[tileIndexB]

			rectangle := abstractions.NewRectangle(tileA, tileB)

			if rectangle.GetArea() > biggestRectangle.GetArea() {
				biggestRectangle = rectangle
			}
		}
	}

	return biggestRectangle
}
