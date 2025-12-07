# 2025 Advent of Code

## Day 1 – Dial Rotations

The Day 1 solution is implemented in Go in the `day_1` folder. It models a dial that can be rotated left or right based on a list of textual instructions, and counts how many times the dial lands exactly on position `0`.

### Problem model

- The dial has positions from `0` to `99` (100 positions in total).
- It starts at position `50`.
- Each instruction in the input describes a **rotation**:
  - The first character is the **direction**: `L` for left, `R` for right.
  - The rest of the line is an integer **distance** to rotate.
- Rotations wrap around the dial:
  - Turning right increases the position.
  - Turning left decreases the position.
  - Both directions wrap within `[0, 99]` using modular arithmetic.
- Every time the dial’s position becomes `0` or crosses it, an internal counter is incremented.

At the end of the run, the program prints how many times the dial ended up at position `0` while processing all instructions.

### High-level flow

The executable entrypoint is `day_1/cmd/main.go` and the flow is:

1. **Read input file path** from the command-line arguments.
2. **Initialize the dial** at position `50`.
3. **Set up a parser** that reads the input file and converts each line into a `Rotation` value (`Direction` + `Distance`).
4. **Set up a processor** that listens on a rotations channel and applies each rotation to the dial.
5. Run the parser and processor concurrently using a `sync.WaitGroup`:
   - The parser streams parsed `Rotation` values into a channel.
   - The processor consumes from that channel and calls `Dial.Rotate`.
6. Wait for both goroutines to finish, then print:
   - `Number of the times the dial ended up at position 0: <count>`

### Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
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
  - `Direction` – simple enum-like type: `Left`, `Right`.
  - `Rotation` – a single rotation instruction: `Direction` + `Distance`.
  - `Dial` – keeps track of:
    - `Position` – current dial position in `[0, 99]`.
    - `count` – how many times the position has been exactly `0`.
  - `Dial.Rotate` delegates to internal helpers:
    - Distances are first reduced modulo the number of positions.
    - `turnRight` moves clockwise: `(Position + distance) % 100`.
    - `turnLeft` moves counter‑clockwise: `(Position - distance + 100) % 100`.
    - After each move, if `Position == 0`, `count` is incremented.
  - `RotationsChannel` – a small interface that exposes a `Rotations() <-chan Rotation` method, so the processor doesn’t depend directly on the parser type.

### Running Day 1

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

## Day 2 – Invalid Product ID Ranges

The Day 2 solution is implemented in Go in the `day_2` folder. It reads ranges of product IDs, finds IDs that are composed of a smaller digit sequence repeated at least twice, and sums all such invalid IDs.

### Problem model

- The input file contains **ranges** of product IDs in the form `A-B`, separated by commas and/or newlines.
  - Example: `11-22,95-115`.
- A **product ID** is a positive integer written as a string of digits, e.g. `123123`.
- A product ID is **invalid** if its entire digit string is made only of some shorter digit sequence repeated **at least twice**. Examples:
  - `12341234` → `"1234"` repeated 2 times → invalid.
  - `123123123` → `"123"` repeated 3 times → invalid.
  - `1212121212` → `"12"` repeated 5 times → invalid.
  - `1111111` → `"1"` repeated 7 times → invalid.
- For each range `A-B`, all IDs from `A` up to and including `B` are considered.
- The program finds every invalid ID across all ranges and sums their numeric values.

At the end of the run, the program prints the total sum of all invalid product IDs, along with how many ranges were processed.

### High-level flow

The executable entrypoint is `day_2/cmd/main.go` and the flow is:

1. **Read input file path** from the command‑line arguments.
2. **Set up a ranges parser** that reads the file and turns each `A-B` fragment into a `Range` (`From` and `To` products).
3. **Set up a ranges processor** that listens on a ranges channel and, for each range, enumerates the product IDs and checks each for validity.
4. Run the parser and processor concurrently using a `sync.WaitGroup`:
   - The parser streams parsed `Range` values into a channel.
   - The processor consumes from that channel and collects invalid product IDs.
5. Wait for all goroutines to finish, then print:
   - `Sum of all the invalid product IDs found in <rangesCount> ranges: <total>`

### Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the ranges parser and processor.
  - Starts them concurrently and waits for completion.
  - Prints the final sum of all invalid IDs.

- `internal/parser`
  - `RangesParser` opens the input file and scans it line by line.
  - Each line is split on `","` into range strings like `"A-B"`.
  - Each range string is split on `"-"` into `from` and `to` product IDs.
  - For each valid `A-B` pair it:
    - Builds `Product` values via `NewProduct`.
    - Wraps them in a `Range` (`From`, `To`).
    - Sends the `Range` into an internal channel.
  - Maintains a `rangesCount` of how many ranges were parsed.
  - Exposes the channel via `Ranges()` and closes it when parsing is done.

- `internal/processor`
  - `RangesProcessor` takes something that implements `RangesChannel` and a `*sync.WaitGroup`.
  - Runs a goroutine that ranges over the emitted `Range` values.
  - For each range it calls `Range.FindInvalidProductIds()`.
  - Accumulates all invalid product IDs into an internal slice and maintains their **sum** in `totalProductId`.
  - Exposes the final sum via `GetTotalProductId()`.

- `internal/abstractions`
  - `Product` – wraps a product ID string and its integer value.
    - `IsValid` returns whether the ID **does not** consist solely of a repeated digit pattern.
    - Internally, `checkHasPattern` checks whether the ID’s length has divisors and tries each possible pattern length.
    - For a given pattern length, it takes the prefix as the candidate pattern and verifies the entire string is made by repeating that pattern at least twice.
  - `Range` – represents a half‑open range of products: `From` and `To`.
    - `FindInvalidProductIds` walks from `From` up to `To`, calling `IsValid` on each `Product` and collecting those that are invalid.
  - `RangesChannel` – a small interface that exposes a `Ranges() <-chan Range` method so that the processor only depends on the interface, not the concrete parser.

### Running Day 2

From the `day_2` directory you can run the solution with:

```bash
cd day_2
make run ARGS="input.txt"
```

or directly with Go:

```bash
cd day_2
go run ./cmd input.txt
```
