# Day 4 – Accessible Rolls (Part 1 & Part 2)

The Day 4 solutions are implemented in Go in the `day_4` folder. They read a 2D map of rolls arranged in rows and sections, determine which rolls are **accessible enough** (i.e., not too crowded by neighboring rolls), and aggregate those counts across the whole map.

- **Part 1** counts how many individual rolls are accessible.
- **Part 2** is built on the same core accessibility logic, but changes how the final result is interpreted/aggregated (for example, using the same detection logic with a different winning condition).

## Problem model (shared by Part 1 & Part 2)

- The input file contains a **grid of spots**, one **row per line**.
- Each character in a line represents a **spot**:
  - `.` – an empty spot.
  - `@` – a roll.
- Conceptually, the map is processed in **sliding sections** of three consecutive rows:
  - The **middle row** of each 3-row window is the row we are analyzing.
  - The row above and the row below provide context to decide if rolls are accessible.

A roll is considered **accessible** if, in its immediate 3×3 neighborhood (itself plus the 8 surrounding spots), the total number of rolls is **strictly less than 4**.

More precisely, for each `@` in the row being analyzed:

1. Look at the row **above**, the **current row**, and the row **below** (when they exist).
2. For each of these three rows, look at the **same column**, plus the immediate left and right columns (i.e., offsets `-1`, `0`, `+1`).
3. Count how many of those spots contain `@`.
4. Subtract `1` for the roll itself (so it isn’t double-counted in the current row).
5. If the resulting count of **surrounding rolls** is `< 4`, the roll is counted as **accessible**.

---

## Part 1 – Total number of accessible rolls

For Part 1, the goal is to compute the **total number of accessible rolls** across all rows that can be analyzed (i.e., all rows that appear as the middle row in a full 3-row window, plus the appropriate handling of the final rows).

At the end, the program prints that total count.

## Part 2 – Same accessibility logic, different aggregate result

Part 2 reuses the exact same low-level definition of an **accessible** roll and the same 3-row sliding window, but changes what you do with the result. Examples of such Part 2-style changes (depending on the specific puzzle variant) include:

- Focusing on **specific rows** or regions instead of the entire map.
- Applying the accessibility logic under a slightly different threshold or rule while keeping the same code structure.
- Combining the count of accessible rolls with other derived metrics.

In the current code base, the accessibility detection ("is this roll accessible?") and row/section handling are centralized in `SectionsProcessor` and `SectionsParser`, so both Part 1 and Part 2 share the same mechanics and differ only in how the final quantity is interpreted or used.

---

## High-level flow (applies to both parts)

The executable entrypoint is `day_4/cmd/main.go`. The overall flow is:

1. **Read the input file path** from the command‑line arguments.
2. Create a `SectionsParser` that reads the file row by row and builds a sliding window of **three rows**.
3. Create a `SectionsProcessor` that analyzes one row (the middle row of the 3-row section) at a time to count accessible rolls.
4. Repeatedly:
   - Ask the parser for the **next section** via `ReadNextRow()`.
   - Pass that section to `SectionsProcessor.Analyze()`.
5. After all rows have been processed, print the final total of accessible rolls (Part 1) or the adjusted aggregate for Part 2.

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the `SectionsParser` and `SectionsProcessor`.
  - Loops over sections until the parser reports there are no more rows to analyze.
  - Prints the final result at the end.

- `internal/parser`
  - `SectionsParser` opens the input file and holds:
    - A `bufio.Scanner` over the input file.
    - A single reusable `Section` instance with **three** `Row` values.
    - Internal state: current row count, which row index is being filled, and whether analysis has started.
  - On construction:
    - Allocates `section.Rows` as a slice of 3 rows.
    - Initially sets `RowIndex = 1`, meaning the middle row will be the first row analyzed once enough data is available.
  - `ReadNextRow()`:
    - Reads the next line from the file (or recognizes the end of file).
    - On the **first** row:
      - Calls `allocateRows` to size the `Spots` slices for all three rows to the line’s length.
      - Fills the **second** row (index 1), leaving the first row empty to begin with.
      - Does **not** yet return a section for analysis, because a lower row is still missing.
    - On the **second** row:
      - Fills the **third** row (index 2).
      - At this point, rows 0, 1, 2 form the initial 3-row window; analysis can begin.
    - On **subsequent** rows:
      - Shifts the window down by copying row 1 → row 0 and row 2 → row 1.
      - Fills row 2 with the new line.
    - For each filled row, it populates `Row.Spots` by converting each character in the line:
      - `'.'` → `Spot('.')` (empty).
      - `'@'` → `Spot('@')` (roll).
    - Tracks a `Number` on each row (1-based row number in the file).
    - Returns the reusable `Section` pointer plus a boolean indicating whether a section is ready to be analyzed.
  - Internally uses:
    - `allocateRows(line string)` – allocates `Spots` slices for all three rows based on line length.
    - `copyRow(fromRowIndex, toRowIndex)` – copies row metadata and spot contents within the current `Section`.

- `internal/processor`
  - `SectionsProcessor` keeps a running count `rollsAccessed` (the total number of accessible rolls found so far).
  - `Analyze(section *Section)`:
    - Logs which row is being processed.
    - Calls `countAccessibleRolls(section)`.
    - Adds the result to `rollsAccessed`.
  - `countAccessibleRolls(section *Section)` implements the core accessibility logic:
    - The `Section` contains three rows, and `RowIndex` tells which row is currently under analysis (initially 1, later 1 or 2 at the end).
    - For each spot in the analyzed row:
      - Skip if it’s not a roll (`@`).
      - Compute indices for the top row, current row, and bottom row.
      - For each relevant row, call `countRolls(row, spotIndex)` to count rolls in the three columns `[spotIndex-1, spotIndex, spotIndex+1]`, staying within bounds.
      - Subtract `1` from the count for the current row so the current roll isn’t counted twice.
      - If the resulting `surroundingRolls < minAccessibleRolls` (4), the roll is counted as accessible.
  - `countRolls(row *Row, spotIndex int)`:
    - Iterates over offsets `-1` to `+1`.
    - Checks that the computed index stays within `[0, len(row.Spots))`.
    - Increments the counter for each spot that contains a roll.

- `internal/abstractions`
  - `Spot` – an enum-like integer type representing a kind of spot:
    - `Empty` – `'.'`.
    - `Roll` – `'@'`.
  - `Row` – a single row of spots:
    - `Number` – the 1-based index of the row in the original file.
    - `Spots` – the sequence of `Spot` values for that row.
  - `Section` – the sliding window of three rows:
    - `Rows []Row` – always length 3 for this puzzle.
    - `RowIndex int` – which row in `Rows` is currently being analyzed.

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
