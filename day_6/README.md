# Day 6 – Column-wise Arithmetic Problems (Part 1)

The Day 6 Part 1 solution is implemented in Go in the `day_6` folder. It reads a grid of numbers with associated operations for each column, applies those operations (multiplication or addition) down each column, and sums the column results to produce a final total.

## Problem model (Part 1)

The input file contains a **grid of numbers** arranged in rows and columns, followed by a final row that specifies an **operation** for each column.

### Input structure

- The file contains several lines of whitespace-separated numbers forming a rectangular grid.
- The **last line** contains operation symbols (either `+` or `*`) instead of numbers, one for each column.

Example input:

```text
10  20  5
2   3   4
5   1   2
+   *   +
```

In this example:

- There are 3 rows of numbers and 3 columns.
- The last row specifies:
  - Column 0: use `+` (addition)
  - Column 1: use `*` (multiplication)
  - Column 2: use `+` (addition)

### Computation rules

For each column:

1. Start with the first number in that column.
2. Apply the column's operation to combine all subsequent numbers in that column, working top to bottom.
3. The result is the **column total**.

After computing the total for each column, **sum all column totals** to get the final answer.

#### Example walkthrough

Using the example input above:

- **Column 0** (operation `+`):
  - Start: `10`
  - `10 + 2 = 12`
  - `12 + 5 = 17`
  - Column total: `17`

- **Column 1** (operation `*`):
  - Start: `20`
  - `20 * 3 = 60`
  - `60 * 1 = 60`
  - Column total: `60`

- **Column 2** (operation `+`):
  - Start: `5`
  - `5 + 4 = 9`
  - `9 + 2 = 11`
  - Column total: `11`

- **Final total**: `17 + 60 + 11 = 88`

The program prints this final total as the Part 1 answer.

---

## High-level flow

The executable entrypoint is `day_6/cmd/main.go`. The high-level steps are:

1. **Read input file path** from the command‑line arguments.
2. **Initialize the problems parser** with that file path via `NewParser`.
3. The parser reads the file and builds a `Problems` structure containing:
   - `Numbers` – a 2D slice of `uint64` representing the grid.
   - `Operations` – a slice of operation strings (`"+"` or `"*"`) for each column.
4. Call `Problems.ComputeTotal()` to compute the answer.
5. Print the final result:

   ```text
   Total = <total>
   ```

---

## Packages and responsibilities

### `cmd/main.go`

- Reads the input path from `os.Args`.
- Calls `initializeParser` to build a `ProblemsParser` from the input file.
- Calls `problemsParser.Problems.ComputeTotal()` to get the final total.
- Prints the result:

  ```go
  fmt.Printf("Total = %d\n", total)
  ```

### `internal/abstractions`

Defined in `problems.go`:

- `Problems` – represents the entire grid and operations:

  ```go
  type Problems struct {
      Numbers    [][]uint64
      Operations []string
  }
  ```

  - `Numbers` is a 2D slice where `Numbers[rowIndex][columnIndex]` is the number at that position.
  - `Operations[columnIndex]` is the operation string for that column (`"+"` or `"*"`).

- `ComputeTotal() uint64` – implements the core computation logic:

  ```go
  func (p *Problems) ComputeTotal() uint64 {
      rowsCount := len(p.Numbers)
      total := uint64(0)

      for columnIndex, operation := range p.Operations {
          columnTotal := uint64(0)

          switch operation {
          case "*":
              for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
                  number := p.Numbers[rowIndex][columnIndex]
                  if rowIndex == 0 {
                      columnTotal = number
                  } else {
                      columnTotal *= number
                  }
              }
          case "+":
              for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
                  number := p.Numbers[rowIndex][columnIndex]
                  if rowIndex == 0 {
                      columnTotal = number
                  } else {
                      columnTotal += number
                  }
              }
          }

          total += columnTotal
      }

      return total
  }
  ```

  For each column:

  - Iterate through all rows in that column.
  - Start with the first number, then apply the column's operation to each subsequent number.
  - Accumulate the result in `columnTotal`.
  - Add `columnTotal` to the running `total`.

  Finally, return the sum of all column totals.

### `internal/parser`

Defined in `problems_parser.go`.

- `ProblemsParser` wraps the parsed `Problems`:

  ```go
  type ProblemsParser struct {
      Problems *abstractions.Problems
  }
  ```

- `NewParser(filePath string)`:
  - Opens the input file.
  - Delegates to `readProblems(filePath)` to build the `Problems` structure.
  - Returns a `ProblemsParser` that exposes the parsed data.

- `readProblems(filePath string)`:
  - Opens the file using `os.OpenFile` and creates a `bufio.Scanner`.
  - Initializes `Problems` with:
    - `Numbers` – a 2D slice preallocated to `DefaultSize` (4) rows.
    - `Operations` – initially empty, allocated when the last line is encountered.
  - Scans the file line by line using `strings.Fields` to split each line into whitespace-separated tokens.
  - For each line:
    - Attempts to parse the first token as a `uint64` using `strconv.ParseUint`.
    - If parsing **succeeds** → the line contains numbers:
      - Allocates a new row slice for that row.
      - Parses each token as a `uint64` and stores it in `Numbers[rowIndex][columnIndex]`.
    - If parsing **fails** → the line contains operations (last line):
      - Allocates `Operations` slice with length equal to the number of tokens.
      - Stores each operation string (e.g. `"+"` or `"*"`) in `Operations[columnIndex]`.
  - Returns the populated `Problems` structure.

---

## Running Day 6 Part 1

From the `day_6` directory you can run the solution with:

```bash
cd day_6
make run ARGS="input.txt"
```

or directly with Go:

```bash
go run ./cmd input.txt
```

