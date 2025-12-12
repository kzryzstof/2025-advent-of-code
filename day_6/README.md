# Day 6 – Column-wise Arithmetic Problems

The Day 6 solution is implemented in Go in the `day_6` folder. It reads a grid of numbers with associated operations for each column, applies those operations (multiplication or addition) down each column, and sums the column results to produce a final total.

The unique twist in this puzzle is that the numbers in the grid are represented as **fixed-width text cells** that must be **transposed** before arithmetic operations can be applied.

---

## Problem model

The input file contains a **grid of fixed-width text cells** arranged in rows and columns, followed by a final row that specifies an **operation** for each column.

### Input structure

- The file contains several lines of fixed-width text representing numbers.
- Each column has a fixed width, and numbers within a column are aligned.
- The **last line** contains operation symbols (either `+` or `*`) instead of numbers, one symbol positioned within each column's width.

Example input:

```text
10  20  5
2    3  4
5    1  2
+   *   +
```

In this example:

- There are 3 rows of numbers and 3 columns.
- Each column is delimited by the position of the operation symbols on the last line.
- The last row specifies:
  - Column 0: use `+` (addition)
  - Column 1: use `*` (multiplication)
  - Column 2: use `+` (addition)

### Key insight: Column transposition

The numbers in each column are stored as **text cells** that must be read **digit by digit** and **transposed** to form actual numbers before applying operations.

For example, if a column contains the cells:

```text
123
45
6
```

These are transposed by reading digits column-wise (right to left within each cell) to potentially form multiple numbers, which are then combined using the column's operation.

This transposition is implemented in `app.TransposeColumns`, which:

1. Converts each cell to its constituent digits (reversed).
2. Rebuilds numbers by reading digits column-wise across all cells.
3. Returns a slice of `uint64` numbers ready for arithmetic operations.

### Computation rules

For each column:

1. Extract the text cells from that column (one per row).
2. **Transpose** the cells using `TransposeColumns` to get a slice of actual numbers.
3. Apply the column's operation (`+` or `*`) to combine all numbers in the slice.
4. The result is the **column total**.

After computing the total for each column, **sum all column totals** to get the final answer.

---

## High-level flow

The executable entrypoint is `day_6/cmd/main.go`. The high-level steps are:

1. **Read input file path** from the command‑line arguments.
2. **Initialize the problems reader** with that file path via `io.NewReader`.
3. Call `reader.Read()` to parse the file and build a `Problems` structure containing:
   - `Numbers` – a 2D slice of `string` representing the text cells.
   - `Operations` – a slice of operation strings (`"+"` or `"*"`) for each column.
4. Call `problems.ComputeTotal()` to:
   - For each column, read the cells, transpose them to numbers, apply the operation, and accumulate the column total.
   - Sum all column totals.
5. Print the final result:

   ```text
   Total = <total>
   ```

---

## Packages and responsibilities

### `cmd/main.go`

- Reads the input path from `os.Args`.
- Calls `initializeReader` to build a `ProblemsReader` from the input file.
- Calls `reader.Read()` to obtain a `*Problems` value.
- Calls `problems.ComputeTotal()` to get the final total.
- Prints the result:

  ```go
  fmt.Printf("Total = %d\n", total)
  ```

### `internal/abstractions`

Defined in `problems.go`:

- `Problems` – represents the entire grid and operations:

  ```go
  type Problems struct {
      Numbers    [][]string
      Operations []string
  }
  ```

  - `Numbers` is a 2D slice where `Numbers[rowIndex][columnIndex]` is the text cell at that position.
  - `Operations[columnIndex]` is the operation string for that column (`"+"` or `"*"`).

- `ComputeTotal() (uint64, error)` – implements the core computation logic:

  ```go
  func (p *Problems) ComputeTotal() (uint64, error) {
      total := uint64(0)

      for columnIndex, operation := range p.Operations {
          // Read all the text cells from the current column
          cells := p.readNumbers(columnIndex)

          // Transpose the cells to get actual numbers
          numbers := app.TransposeColumns(cells)

          // Perform the operation on the numbers
          columnTotal, err := app.Compute(operation, numbers)
          if err != nil {
              return 0, err
          }

          total += columnTotal
      }

      return total, nil
  }
  ```

  For each column:

  - Calls `readNumbers(columnIndex)` to extract the text cells from that column.
  - Calls `app.TransposeColumns(cells)` to convert cells to a slice of `uint64` numbers.
  - Calls `app.Compute(operation, numbers)` to apply the column's operation.
  - Accumulates the result in `total`.

- `readNumbers(columnIndex int) []string` – helper that extracts all cells from a given column into a slice.

### `internal/io`

Defined in `problems_reader.go`.

- `ProblemsReader` encapsulates reading problems from the input file:

  ```go
  type ProblemsReader struct {
      inputFile *os.File
  }
  ```

- `NewReader(filePath string)`:
  - Opens the input file read-only.
  - Returns a `ProblemsReader` bound to that file.

- `Read() (*abstractions.Problems, error)`:
  - Uses a `bufio.Scanner` to read all lines into memory first (needed to determine column widths from the last line).
  - Identifies the last line as the operations line.
  - Parses the operations line character by character to determine:
    - Where each column starts and ends (delimited by operation symbols and spaces).
    - The operation symbol for each column.
  - For each identified column:
    - Extracts the corresponding substring from each number row.
    - Stores it as a text cell in `Numbers[rowIndex][columnIndex]`.
  - Returns a `*Problems` with:
    - `Numbers` – 2D slice of text cells.
    - `Operations` – slice of operation strings.

### `internal/app`

Defined in `operations.go` and `transpose_columns.go`:

- `Compute(operation string, numbers []uint64) (uint64, error)`:
  - Takes an operation (`"+"` or `"*"`) and a slice of numbers.
  - Uses `applyOperation` with a corresponding lambda to compute the result:
    - `*` → multiply all numbers together.
    - `+` → sum all numbers together.
  - Returns the final result.

- `applyOperation(numbers []uint64, operation func(uint64, uint64) uint64) uint64`:
  - Starts with the first number.
  - Iterates through remaining numbers, applying the given operation function cumulatively.

- `TransposeColumns(cells []string) []uint64`:
  - Converts each text cell to its digits (reversed).
  - Determines how many numbers to produce (based on the longest cell).
  - Rebuilds numbers digit by digit, reading column-wise across all cells.
  - Returns a slice of `uint64` numbers.

### `internal/extensions`

Utility functions for string/slice manipulation (e.g., `ToDigits`, `Reverse`).

---

## Running Day 6

From the `day_6` directory you can run the solution with:

```bash
cd day_6
make run ARGS="input.txt"
```

or directly with Go:

```bash
cd day_6
go run ./cmd input.txt
```

