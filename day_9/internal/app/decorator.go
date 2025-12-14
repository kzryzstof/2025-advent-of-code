package app

import (
	"day_9/internal/abstractions"
	"fmt"
)

func ArrangeTiles(
	movieTheater *abstractions.MovieTheater,
) *abstractions.Rectangle {

	abstractions.Draw(movieTheater.GetTiles(), nil)

	biggestRectangle := abstractions.NewRectangle(
		&abstractions.Tile{X: 0, Y: 0},
		&abstractions.Tile{X: 0, Y: 0},
	)

	tiles := movieTheater.GetTiles()

	for tileIndexA, tileA := range tiles {
		if tileA.Color != abstractions.Red {
			continue
		}
		for tileIndexB, tileB := range tiles {
			if tileIndexA == tileIndexB {
				continue
			}

			if tileB.Color != abstractions.Red {
				continue
			}

			rectangle := abstractions.NewRectangle(tileA, tileB)

			if rectangle.GetArea() > biggestRectangle.GetArea() {

				if rectangle.GetArea() == 24 {
					fmt.Print()
				}
				if !rectangle.IsInside(tiles) {
					continue
				}
				biggestRectangle = rectangle
			}
		}
	}

	return biggestRectangle
}
