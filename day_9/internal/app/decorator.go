package app

import (
	"day_9/internal/abstractions"
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	// TilesValidationConcurrency on macOS M1 Pro, 10 seems to be a nice sweet spot
	TilesValidationConcurrency = 4
)

func ArrangeTiles(
	movieTheater abstractions.MovieTheater,
) *abstractions.Rectangle {

	biggestRectangle := abstractions.NewRectangle(
		&abstractions.Tile{X: 0, Y: 0},
		&abstractions.Tile{X: 0, Y: 0},
	)

	redTiles := movieTheater.GetRedTiles()

	var (
		processedTiles    uint
		innerTilesCount   uint
		maxTilesPerSecond uint

		processedTilesMutex sync.Mutex
		rectangleMutex      sync.Mutex

		tilesWaitGroup    sync.WaitGroup
		progressWaitGroup sync.WaitGroup
	)

	sem := make(chan struct{}, TilesValidationConcurrency)

	/* Progress reporter */
	/* Prints the number of processed rows every second until all row goroutines have finished */
	progressWaitGroup.Add(1)
	go func(totalTiles uint) {
		defer progressWaitGroup.Done()

		lastDone := uint(0)

		for {
			processedTilesMutex.Lock()
			_processedTiles := processedTiles
			_innerTilesCount := innerTilesCount
			processedTilesMutex.Unlock()

			tilesPerSecond := _processedTiles - lastDone
			maxTilesPerSecond = uint(math.Max(float64(maxTilesPerSecond), float64(tilesPerSecond)))
			fmt.Printf("\t%05d/%05d processed tiles (%02d) (%05d tiles per second)\r", _processedTiles, totalTiles, _innerTilesCount, tilesPerSecond)

			if _processedTiles >= totalTiles {
				// All rows processed; exit the progress loop
				return
			}

			lastDone = _processedTiles

			time.Sleep(time.Second)
		}
	}(uint(len(redTiles)))

	for tileIndexA := 0; tileIndexA < len(redTiles); tileIndexA++ {

		innerTilesCount := 0

		for tileIndexB := tileIndexA + 1; tileIndexB < len(redTiles); tileIndexB++ {

			tilesWaitGroup.Add(1)

			go func(tileA *abstractions.Tile, tileB *abstractions.Tile) {

				defer tilesWaitGroup.Done()

				sem <- struct{}{}
				defer func() { <-sem }() // release slot

				rectangle := abstractions.NewRectangle(tileA, tileB)

				/* Initial check: is the rectangle potentially bigger than the biggest found so far? */
				isRectangleBigger := rectangle.GetArea() > biggestRectangle.GetArea()

				if !isRectangleBigger {
					processedTilesMutex.Lock()
					innerTilesCount++
					processedTilesMutex.Unlock()
					return
				}

				/* Validates that the rectangle is inside the movie theater */
				if !rectangle.IsInside(movieTheater) {
					processedTilesMutex.Lock()
					innerTilesCount++
					processedTilesMutex.Unlock()
					return
				}

				rectangleMutex.Lock()

				if rectangle.GetArea() > biggestRectangle.GetArea() {
					//fmt.Printf("\tBiggest rectangle: %v\r", rectangle)
					biggestRectangle = rectangle
				}

				rectangleMutex.Unlock()

				processedTilesMutex.Lock()
				innerTilesCount++
				processedTilesMutex.Unlock()

			}(redTiles[tileIndexA], redTiles[tileIndexB])
		}

		tilesWaitGroup.Wait()

		processedTilesMutex.Lock()
		processedTiles++
		processedTilesMutex.Unlock()
	}

	// Wait for all rows to finish, then for the progress reporter to exit.
	progressWaitGroup.Wait()

	fmt.Printf("\tBiggest rectangle: %v\n", biggestRectangle)

	return biggestRectangle
}
