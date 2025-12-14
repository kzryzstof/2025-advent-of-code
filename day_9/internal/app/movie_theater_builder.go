package app

import (
	"day_9/internal/abstractions"
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	// ValidTilesComputationConcurrency on macOS M1 Pro, 10 seems to be a nice sweet spot
	ValidTilesComputationConcurrency = 10
)

type movieTheater struct {
	redTiles []*abstractions.Tile
	/* Contains all the valid tiles (green and red) */
	validTiles [][]bool
}

func NewMovieTheater(
	redTiles []*abstractions.Tile,
) abstractions.MovieTheater {
	return &movieTheater{
		redTiles:   redTiles,
		validTiles: ComputeValidTiles(redTiles),
	}
}

func (mt *movieTheater) GetRedTiles() []*abstractions.Tile {
	return mt.redTiles
}

func (mt *movieTheater) IsValidTile(
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

func ComputeValidTiles(
	redTiles []*abstractions.Tile,
) [][]bool {

	maxX, maxY := abstractions.FindFloorSize(redTiles)

	fmt.Printf("\tFloor size is %d x %d\n", maxX, maxY)

	/* Initializes the valid tiles map: this is a big one so let's allocate everything upfront */
	validTiles := make([][]bool, maxY)

	for y := uint(0); y < maxY; y++ {
		validTiles[y] = make([]bool, maxX)
	}

	fmt.Print("\tMemory allocated\n")

	/* The red tiles polygon must be ordered for the ray tracing to work */
	abstractions.OrderPolygonVertices(redTiles)

	fmt.Print("\tVertices ordered\n")

	fmt.Print("\tPre-computing valid tiles\n")

	var (
		processedRows uint
		mu            sync.Mutex
		wgRows        sync.WaitGroup
		wgProgress    sync.WaitGroup
	)

	sem := make(chan struct{}, ValidTilesComputationConcurrency)

	maxRowsPerSecond := 0

	/* Progress reporter */
	/* Prints the number of processed rows every second until all row goroutines have finished */
	wgProgress.Add(1)
	go func(totalRows uint) {
		defer wgProgress.Done()

		lastDone := uint(0)

		for {
			mu.Lock()
			done := processedRows
			mu.Unlock()

			rowsPerSecond := done - lastDone
			maxRowsPerSecond = int(math.Max(float64(maxRowsPerSecond), float64(rowsPerSecond)))
			fmt.Printf("\t\t%05d/%05d processed rows (%05d rows per second)\r", done, totalRows, rowsPerSecond)

			if done >= totalRows {
				// All rows processed; exit the progress loop
				return
			}

			lastDone = done

			time.Sleep(time.Second)
		}
	}(maxY)

	for y := uint(0); y < maxY; y++ {

		y := y // capture loop variable
		wgRows.Add(1)

		go func() {
			defer wgRows.Done()

			sem <- struct{}{}
			defer func() { <-sem }() // release slot

			for x := uint(0); x < maxX; x++ {

				isRedTile := isRedTile(redTiles, x, y)

				if isRedTile {

					/* This is a red tile, so let's add it to your new tiles */
					validTiles[y][x] = true

				} else {

					/* This is NOT a red tile, so let's see if it is inside the red tiles polygon */
					validTiles[y][x] = abstractions.IsPointInPolygon(
						redTiles,
						&abstractions.Tile{X: x, Y: y},
					)
				}
			}

			mu.Lock()
			processedRows++
			mu.Unlock()
		}()
	}

	// Wait for all rows to finish, then for the progress reporter to exit.
	wgRows.Wait()
	wgProgress.Wait()

	fmt.Printf("\t%05d processed rows (max of %05d rows per second)\n", maxY, maxRowsPerSecond)

	return validTiles
}

func isRedTile(
	tiles []*abstractions.Tile,
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
