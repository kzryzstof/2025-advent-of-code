# 2025 Advent of Code

This repository contains my solutions for the **Advent of Code 2025** programming puzzles, implemented in Go, a language I am currently learning.

Each day lives in its own folder (`day_1`, `day_2`, `day_3`, `day_4`, …) with its own `go.mod`, a small command-line entrypoint under `cmd/`, and an `internal` package that holds the core logic.

## Structure

- `day_1/` – solution for Day 1 (dial rotations). See `day_1/README.md` for full details.
- `day_2/` – solution for Day 2 (invalid product ID ranges). See `day_2/README.md` for full details.
- `day_3/` – solution for Day 3 (battery banks, 12-digit greedy voltage). See `day_3/README.md` for full details.
- `day_4/` – solution for Day 4 (accessible rolls removal using a forklift). See `day_4/README.md` for full details.
- `day_5/` – solution for Day 5 (fresh ingredient ranges with compaction). See `day_5/README.md` for full details.
- `day_6/` – solution for Day 6 (column-wise arithmetic with cell transposition). See `day_6/README.md` for full details.
- `day_7/` – solution for Day 7 (multiverse timeline splitter with tachyon beams). See `day_7/README.md` for full details.
- `day_8/` – solution for Day 8 (junction box playground wiring). See `day_8/README.md` for full details.
- `day_9/` – solution for Day 9 (movie theater seating and largest rectangle). See `day_9/README.md` for full details.
- `day_10/` – solution for Day 10 (factory machine joltage configuration using Hermite Normal Form). See `day_10/README.md` for full details.
- `day_11/` – solution for Day 11 (reactor device graph with constrained path counting). See `day_11/README.md` for full details.
- `day_12/` – solution for Day 12 (packing weird-shaped presents under Christmas trees). See `day_12/README.md` for full details.

Each day generally includes:

- An `input.txt` file containing the puzzle input.
- A `Makefile` with convenience targets to run the solution.
- Unit tests under `internal/...` validating the core logic.

## Running a day

To run a given day, `cd` into the corresponding folder and either use `make`:

```bash
cd day_1
make run ARGS="input.txt"
```

or call `go run` directly on the command:

```bash
cd day_1
go run ./cmd input.txt
```

Replace `day_1` with `day_2`, `day_3`, `day_4`, … to run other days.
