# Day 1 – Dial Rotations

The Day 1 solution is implemented in Go in the `day_1` folder. It models a dial that can be rotated left or right based on a list of textual instructions, and counts how many times the dial lands passed by position `0`.

## Problem model

- The dial has a fixed number of positions (100 in this implementation), numbered from `0` to `99`.
- It starts at a given initial position (by default `50`).
- Each instruction in the input describes a **rotation**:
  - The first character is the **direction**: `L` for left, `R` for right.
  - The rest of the line is an integer **distance** to rotate.
- Rotations wrap around the dial:
  - Turning right increases the position.
  - Turning left decreases the position.
  - Both directions wrap within `[0, positions-1]` using modular arithmetic.
- Every time the dial’s position passes by position `0`, an internal counter is incremented.

At the end of the run, the program prints how many times the dial passed by position `0` while processing all instructions.

## High-level flow

The executable entrypoint is `day_1/cmd/main.go` and the flow is:

1. **Read input file path** from the command-line arguments.
2. **Initialize the dial** at the starting position.
3. **Set up a parser** that reads the input file and converts each line into a `Rotation` value (`Direction` + `Distance`).
4. **Set up a processor** that listens on a rotations channel and applies each rotation to the dial.
5. Run the parser and processor concurrently using a `sync.WaitGroup`:
   - The parser streams parsed `Rotation` values into a channel.
   - The processor consumes from that channel and calls `Dial.Rotate(rotation)`.
6. Wait for both goroutines to finish, then print:
   - `Number of the times the dial ended up at position 0: <count>`

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the dial, parser, and processor.
  - Starts the concurrent work and prints the final result.

- `internal/parser`
  - `InstructionsParser` opens the input file and reads it line by line using a scanner.
  - For each non-empty line, it:
    - Parses the direction (`L`/`R`) into a `Direction` enum.
    - Parses the integer distance using `strconv.Atoi`.
    - Emits a `Rotation` on an internal channel.
  - Exposes that channel via the `Rotations()` method, and closes the channel when parsing is done.

- `internal/processor`
  - `InstructionsProcessor` takes something that implements `RotationsChannel` and a `*Dial`.
  - Runs a goroutine that ranges over the rotations channel and, for each rotation, calls `dial.Rotate(rotation)`.

- `internal/abstractions`
  - `Direction` – simple enum-like type for the rotation direction (`Left`, `Right`).
  - `Rotation` – a single rotation instruction: `Direction` + `Distance`.
  - `Dial` – keeps track of:
    - `Position` – current dial position.
    - `count` – how many times the position has been exactly `0`.
  - `Dial.Rotate` delegates to internal helpers:
    - Normalizes distances using modulo with the dial size.
    - `turnRight` moves clockwise.
    - `turnLeft` moves counter-clockwise.
    - After each move, if `Position == 0`, `count` is incremented.
  - `RotationsChannel` – a small interface that exposes a `Rotations() <-chan Rotation` method, so the processor doesn’t depend directly on the parser type.

## Running Day 1

From the `day_1` directory you can run the solution with:

```bash
cd day_1
make run ARGS="input.txt"
```

or directly with Go:

```bash
cd day_1
go run ./cmd input.txt
```

