# Day 9 – Movie Theater Seating

The Day 9 solution is implemented in Go in the `day_9` folder. It models a movie theater floor delimited by a polygon of red tiles and computes the largest axis-aligned rectangle that fits entirely inside this area.

## Problem model

- The floor is a 2D grid of tiles.
- The input file lists **red tiles**, each with integer coordinates `(X, Y)`.
- These red tiles are the vertices of a polygon that defines the valid seating area of the theater.
- We are interested in finding the **largest possible rectangle** such that:
  - Its two opposite corners are located on red tiles.
  - Every tile inside that rectangle lies inside the polygon (i.e. is a valid tile).

At the end of the run, the program prints the area of this largest rectangle.

## High-level flow

The executable entrypoint is `day_9/cmd/main.go` and the flow is:

1. **Read input path** from the command-line arguments.
2. **Parse red tiles** from the input file using a dedicated reader.
3. **Build the movie theater model** via `NewMovieTheater`:
   - Compute the floor size (bounding box) from the red tiles.
   - Allocate a 2D `[][]bool` grid representing all tiles on the floor.
   - Order the red tiles into a proper polygon.
   - For each tile on the floor, determine if it lies inside the polygon using a point-in-polygon (ray casting) algorithm and mark it as valid.
   - This precomputation is done concurrently across multiple rows and reports progress.
4. **Search for the biggest rectangle** with `ArrangeTiles`:
   - Iterate over all pairs of red tiles and construct a candidate rectangle.
   - Quickly discard rectangles that are not larger than the best one found so far.
   - Use the precomputed `validTiles` grid to check if all tiles inside the rectangle are valid.
   - Run these checks concurrently with a bounded number of goroutines and report progress.
5. **Output**:
   - Number of red tiles in the movie theater.
   - Area of the largest valid rectangle.

## Packages and responsibilities

- `cmd/main.go`
  - Wires together input reading, theater construction, and rectangle search.
  - Prints diagnostic information and the final result.

- `internal/io`
  - `RedTilesReader` reads the input file and converts each line into an `abstractions.Tile`.

- `internal/abstractions`
  - Core domain types: `Tile`, `Rectangle`, and `MovieTheater` (interface).
  - Geometric helpers:
    - `FindFloorSize` to compute the bounding box.
    - `OrderPolygonVertices` to order vertices into a polygon.
    - `IsPointInPolygon` to test whether a tile lies inside the polygon.

- `internal/app`
  - `NewMovieTheater` (`movie_theater_builder.go`):
    - Implements the concrete `movieTheater` type that stores red tiles and the `validTiles` grid.
    - Precomputes the valid tiles concurrently and exposes the `MovieTheater` interface.
  - `ArrangeTiles` (`decorator.go`):
    - Searches all candidate rectangles defined by pairs of red tiles.
    - Uses the `MovieTheater` interface to check whether a rectangle is fully inside the valid area.
    - Runs the validation in parallel with a configurable level of concurrency and displays progress statistics.

