# Day 7 – Tachyon Beam Splitter

This day simulates a **tachyon beam** travelling through a 2D grid (the *manifold*).

## Problem overview

The input (`input.txt`) is a rectangular grid of characters:

- `S` – starting point of the beam
- `.` – empty space
- `^` – splitter
- `|` – beam path (only appears in the output, not in the input)

Example (shortened):

```text
....S....
.........
....^....
.........
```

From the single `S`, a **tachyon** starts moving **downwards** one cell at a time:

- If the next cell is **empty** (`.`), the beam moves into it and marks it as `|`.
- If the next cell is a **splitter** (`^`), the beam **splits** into two beams that diverge horizontally (left and right), continuing from the splitter’s position.
- If the next cell is **outside the manifold** or already contains another beam, that tachyon **stops**.

The simulation runs until **all tachyons have stopped**. At the end, the program:

1. Draws the final manifold (with all paths marked as `|`).
2. Prints how many times the beam has been **split** in total.

## Code structure

- `internal/io/manifold_reader.go` – reads the grid from `input.txt` into a `Manifold`.
  - Finds the starting point `S` and creates an initial moving tachyon.
- `internal/abstractions/manifold.go` – represents the grid and the beam operations:
  - Queries and updates locations
  - Checks boundaries
  - Splits beams and tracks all tachyons
- `internal/abstractions/tachyon.go` – represents a single moving tachyon (position + moving/stopped state).
- `internal/abstractions/direction.go` – small helper for row/column deltas.
- `internal/app/beam_simulator.go` – core simulation loop (`Simulate`):
  - Keeps stepping all moving tachyons until none are moving.
- `cmd/main.go` – CLI entrypoint that wires everything together.

## Running the solution

From the `day_7` directory:

```bash
cd day_7
# Run with the provided input file
go run ./cmd ./input.txt
```

The program will print the final manifold and a line like:

```text
The beam has been splitted X times
```

where `X` is the total number of splits that occurred during the simulation.
