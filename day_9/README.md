# Day 9 – Movie Theater Red Tiles

Santa needs to decorate a movie theater floor. The floor is covered with red tiles, and he wants to find the **biggest rectangle** whose corners match existing red tiles.

## Problem Description

Given a list of `(x,y)` coordinates representing red tiles, compute the largest axis-aligned rectangle by area. The rectangle's corners must coincide with red tiles from the input.

The area is calculated **inclusively** on both axes:

```
area = (|x2 - x1| + 1) * (|y2 - y1| + 1)
```

## Input Format

Each line contains two comma-separated unsigned integers representing a tile's `(x,y)` position:

```
98149,50096
98149,51320
98283,51320
...
```

## Solution

The code is organized as follows:

| Path | Description |
|------|-------------|
| `cmd/main.go` | CLI entry point; reads input, calls `ArrangeTiles`, prints result. |
| `internal/abstractions/tile.go` | `Tile` struct with `X`, `Y` coordinates. |
| `internal/abstractions/rectangle.go` | `Rectangle` struct; computes inclusive area via `GetArea()`. |
| `internal/abstractions/movie_theater.go` | `MovieTheater` holding a slice of red tiles. |
| `internal/app/decorator.go` | `ArrangeTiles` – finds the biggest rectangle among all tile pairs. |
| `internal/io/red_tiles_reader.go` | Parses the input file into a `MovieTheater`. |

## Running

```bash
make run ARGS="input.txt"
```

or

```bash
go run ./cmd input.txt
```

## Testing

Unit tests use table-driven style:

```bash
go test ./...
```

Key test files:

- `internal/abstractions/rectangle_test.go` – validates `NewRectangle` and `GetArea`.
- `internal/app/decorator_test.go` – validates `ArrangeTiles` with various tile configurations.

