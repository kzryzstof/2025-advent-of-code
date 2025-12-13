# Day 7 – Multiverse Timeline Splitter

This solution simulates **tachyon beams** traveling downward through a 2D grid (the *manifold*), splitting into multiple timelines when they encounter splitter obstacles.

## Problem Overview

### Input Format

The input file (`input.txt`) is a rectangular grid with:

- `S` – Starting point where the initial tachyon beam originates
- `.` – Empty space (passable)
- `^` – Splitter (causes the beam to split into two timelines)
- `|` – Beam path marker (written during simulation)

Example input:

```text
......S......
.............
.....^.^.....
.............
....^...^....
.............
```

### Simulation Rules

The simulation starts with a single tachyon at position `S`, moving **downward** (row + 1):

1. **Empty cell (`.`)**: The tachyon moves forward and marks the cell as `|`
2. **Splitter (`^`)**: The beam splits into two new beams that diverge horizontally:
   - One beam moves **left** (col - 1)
   - One beam moves **right** (col + 1)
   - Both continue moving downward from their new positions
   - A new **timeline** is created (à la Marvel multiverse)
3. **Another beam (`|`)**: The tachyons **merge** — we track that multiple beams follow the same path rather than simulating them separately (optimization)
4. **Out of bounds or obstacle**: The tachyon stops

The simulation continues until **all tachyons have stopped moving**.

### Timeline Counting

Each time a beam hits a splitter, a new timeline branches off. The program tracks:

- How many beams hit each splitter position
- The total number of timelines created across the entire simulation
- The initial "sacred timeline" (always counts as 1)

**Final output**: Total timelines = 1 (original) + sum of all beams that hit splitters.

## Code Structure

### Core Abstractions (`internal/abstractions/`)

- **`manifold.go`** – The grid and simulation state:
  - `Locations [][]string` – 2D grid of cells
  - `Tachyons []*Tachyon` – All active/stopped tachyons
  - `timelines [][]uint64` – Tracks how many beams hit each splitter position
  - Methods: `SetBeamAt`, `SplitBeamAt`, `Merge`, `CountTimelines`
  
- **`tachyon.go`** – Represents a single beam:
  - `Position` – Current row/column
  - `isMoving bool` – Whether still active
  - `mergedBeams uint64` – Tracks how many beams merged into this one
  - Methods: `Move`, `Stop`, `Split`, `MergeTo`

- **`position.go`** – Simple (row, col) struct with `MoveTo(direction)` helper

- **`direction.go`** – Delta struct for movement (RowDelta, ColDelta)

### Application Logic (`internal/app/`)

- **`beam_simulator.go`** – Main simulation loop:
  - `Simulate(manifold, drawProgress bool)` – Steps all tachyons until none are moving
  - Optionally clears and redraws the grid each iteration for animation

### I/O (`internal/io/`)

- **`manifold_reader.go`** – Parses `input.txt`:
  - Reads grid into 2D string array
  - Finds all `S` positions and creates initial tachyons
  - Returns a `Manifold` ready for simulation

### Entry Point (`cmd/`)

- **`main.go`** – CLI that:
  1. Reads the manifold from the input file
  2. Runs the simulation with live drawing enabled
  3. Prints the total number of timelines created

## Running the Solution

```bash
cd day_7
go run ./cmd ./input.txt
```

### Output

The program will:
1. Animate the simulation (clearing and redrawing the grid as tachyons move)
2. Print the final grid state with all beam paths marked as `|`
3. Display:
   ```
   The beam has created X timelines
   ```
   where `X` is the total timeline count (1 original + all splits)

## Key Implementation Details

### Beam Merging Optimization

When multiple beams converge on the same path, instead of simulating each one:
- The beams **merge** into a single tachyon
- The `mergedBeams` counter tracks how many actual beams this tachyon represents
- When this merged tachyon hits a splitter, all its constituent beams contribute to the timeline count

### Timeline Tracking

The `timelines` 2D array stores beam counts at each splitter position:
```go
m.timelines[row][col] += tachyon.GetMergedBeams()
```
This ensures that if 5 merged beams hit a single splitter, it creates 5 new timelines (not just 1).

### Animation

When `drawProgress` is true:
- Each iteration clears the terminal: `fmt.Print("\033[2J\033[H")`
- Redraws the entire grid with current beam positions
- Creates a visual animation of the beam propagation

