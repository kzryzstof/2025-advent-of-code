package abstractions

import (
	"fmt"
	"sync"
	"time"
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

	var (
		processedRows uint
		mu            sync.Mutex
		wgRows        sync.WaitGroup
		wgProgress    sync.WaitGroup
	)

	/* Progress reporter */
	/* Prints the number of processed rows every second until all row goroutines have finished */
	wgProgress.Add(1)
	go func(totalRows uint) {
		defer wgProgress.Done()

		for {
			mu.Lock()
			done := processedRows
			mu.Unlock()

			fmt.Printf("%05d/%05d processed rows\r", done, totalRows)

			if done >= totalRows {
				// All rows processed; exit the progress loop
				return
			}

			time.Sleep(time.Second)
		}
	}(maxY)

	for y := uint(0); y < maxY; y++ {

		y := y // capture loop variable
		wgRows.Add(1)

		go func() {
			defer wgRows.Done()

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

			mu.Lock()
			processedRows++
			mu.Unlock()
		}()
	}

	// Wait for all rows to finish, then for the progress reporter to exit.
	wgRows.Wait()
	wgProgress.Wait()

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
