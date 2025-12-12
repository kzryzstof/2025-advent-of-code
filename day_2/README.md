# Day 2 – Invalid Product ID Ranges

The Day 2 solution is implemented in Go in the `day_2` folder. It reads ranges of product IDs, finds IDs that are composed of a smaller digit sequence repeated at least twice, and sums all such invalid IDs.

## Problem model

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

## High-level flow

The executable entrypoint is `day_2/cmd/main.go` and the flow is:

1. **Read input file path** from the command‑line arguments.
2. **Use the ranges parser** to read the file and turn each `A-B` fragment into a `Range` (`From` and `To` products), returning a slice of ranges.
3. Call the application helper `FindInvalidProductIds(ranges)` to:
   - Walk every range.
   - Enumerate each product ID in that range.
   - Check if the ID is invalid (made of a smaller repeated digit pattern).
   - Accumulate the sum of all invalid IDs.
4. Print:
   - `Sum of all the invalid product IDs found in <rangesCount> ranges: <total>`

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Uses the parser to obtain all ranges from the input file.
  - Calls `app.FindInvalidProductIds` with that slice.
  - Prints the final sum of all invalid IDs.

- `internal/io` (or `internal/parser`, depending on naming)
  - `RangesParser` opens the input file and scans it line by line.
  - Each line is split on `","` into range strings like `"A-B"`.
  - Each range string is split on `"-"` into `from` and `to` product IDs.
  - For each valid `A-B` pair it:
    - Builds `Product` values via `NewProduct`.
    - Wraps them in a `Range` (`From`, `To`).
  - Returns the full slice of parsed ranges.

- `internal/app`
  - `FindInvalidProductIds(ranges []abstractions.Range) uint64` – core application logic used by both `main` and tests:
    - Iterates over each `Range`.
    - Calls `Range.FindInvalidProductIds()` to get all invalid IDs in that range.
    - Sums their numeric values into a single `uint64` result.

- `internal/abstractions`
  - `Product` – wraps a product ID string and its integer value.
    - `IsValid` returns whether the ID **does not** consist solely of a repeated digit pattern.
    - Internally, `checkHasPattern` checks whether the ID’s length has divisors and tries each possible pattern length.
    - For a given pattern length, it takes the prefix as the candidate pattern and verifies the entire string is made by repeating that pattern at least twice.
  - `Range` – represents an inclusive range of products: `From` and `To`.
    - `FindInvalidProductIds` walks from `From` up to `To`, calling `IsValid` on each `Product` and collecting those that are invalid.

## Running Day 2

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
