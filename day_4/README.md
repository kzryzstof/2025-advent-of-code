# Day 4 – Accessible Rolls

The Day 4 solution is implemented in Go in the `day_4` folder. It reads a 2D map of rolls arranged in rows (a department), repeatedly removes rolls that are **accessible enough** by the forklift (i.e., not too crowded by neighboring rolls), and counts how many rolls were removed in total.

---

## Problem model

- The input file contains a **grid of spots**, one **row per line**.
- Each character in a line represents a **spot**:
  - `.` – an empty spot.
  - `@` – a roll.
- All rows are read into memory as a `Department`, which is simply a collection of `Row` values, each with a sequence of `Spot`s.

A roll is considered **accessible** if, in its immediate 3×3 neighborhood (itself plus the 8 surrounding spots), the total number of rolls is **strictly less than 4**.

More precisely, for each `@` in a given row at index `rowIndex` and column `spotIndex`:

1. Look at the row **above**, the **current row**, and the row **below** (when they exist).
2. For each of these rows, look at the **same column**, plus the immediate left and right columns (i.e., offsets `-1`, `0`, `+1`).
3. Count how many of those spots contain `@`.
4. Subtract `1` for the roll itself (so it isn’t double-counted in the current row).
5. If the resulting count of **surrounding rolls** is `< 4`, the roll is counted as **accessible** and is removed (its spot becomes empty).

The forklift keeps scanning the department and removing accessible rolls **in multiple passes** until an entire pass finds **no more accessible rolls**. The final answer is the total number of rolls that were removed across all passes.

---

## High-level flow

The executable entrypoint is `day_4/cmd/main.go`. The overall flow is:

1. **Read the input file path** from the command‑line arguments.
2. Create a `DepartmentParser` (reader) that opens the input file.
3. Call `reader.Read()` to read all rows into a `Department` structure.
4. Create a `Forklift` with an optional `verbose` flag.
5. Call `forklift.RemoveRolls(department)` to:
   - Loop over the rows repeatedly.
   - On each pass, find and remove accessible rolls.
   - Stop once a full pass removes no rolls.
6. After processing, print the total number of accessible rolls the forklift removed:
   - `Number of accessible rolls in the <rowsCount> rows of the department: <count>`

---

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the `DepartmentParser` via `io.NewReader`.
  - Calls `reader.Read()` to obtain a `*Department` with all rows.
  - Creates a `Forklift` and calls `RemoveRolls(department)`.
  - Prints the final total of accessible rolls removed.

- `internal/io`
  - `DepartmentParser` encapsulates reading the department from the input file.
  - `NewReader(filePath string)` opens the file and returns a parser.
  - `Read() (*abstractions.Department, error)`:
    - Uses a `bufio.Scanner` to read the file line by line.
    - For each line, builds a `Row` with:
      - `Number` – the 1-based row index.
      - `Spots` – a slice of `Spot` values created from the characters (`.` or `@`).
    - Appends each `Row` to the `Department.Rows` slice and returns the `Department` at the end.

- `internal/app`
  - `Forklift` – core domain type that removes accessible rolls:

    ```go
    type Forklift struct {
        verbose       bool
        rollsAccessed uint
    }
    ```

    - `NewForklift(verbose bool)` creates a new forklift.
    - `GetAccessedRollsCount() uint` returns how many rolls have been removed.
    - `RemoveRolls(department *abstractions.Department)`:
      - Repeatedly scans all rows.
      - For each row, calls `countAccessibleRolls` to find and remove accessible rolls.
      - Accumulates `rollsAccessed` and stops when a scan finds no accessible rolls.
    - `countAccessibleRolls(department *Department, rowIndex uint) uint`:
      - For each spot in the given row:
        - Skips non-roll spots.
        - Counts surrounding rolls using `countRolls` on the row above (if any), current row, and row below (if any).
        - If surrounding rolls `< minAccessibleRolls` (4), increments the accessible-roll count and sets that spot to `Empty`.
    - `countRolls(row *Row, spotIndex int) uint`:
      - Looks at indices `spotIndex-1`, `spotIndex`, and `spotIndex+1` within bounds.
      - Counts how many of those spots contain a roll.

- `internal/abstractions`
  - `Spot` – an enum-like integer type representing a kind of spot:
    - `Empty` – `'.'`.
    - `Roll` – `'@'`.
  - `Row` – a single row of spots:
    - `Number` – the 1-based index of the row in the original file.
    - `Spots` – the sequence of `Spot` values for that row.
  - `Department` – the full collection of rows:
    - `Rows []Row` – all rows of the grid.

---

## Running Day 4

From the `day_4` directory you can run the solution with:

```bash
cd day_4
make run ARGS="input.txt"
```

or directly with Go:

```bash
cd day_4
go run ./cmd input.txt
```
