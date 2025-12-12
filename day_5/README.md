# Day 5 – Fresh Ingredients

The Day 5 solution is implemented in Go in the `day_5` folder. It reads a list of **freshness ranges** for ingredient IDs, compacts those ranges by merging overlaps and adjacencies, and then counts how many individual ingredient IDs are fresh.

The core steps are:

1. Parse all fresh ranges from the input file.
2. **Compact** the fresh ranges by merging overlapping or touching ranges into a minimal set of ranges.
3. Count how many distinct ingredient IDs are covered by the compacted ranges.
4. Print that count.

---

## Problem model

The input file contains **fresh ingredient ranges**, one per line, in the form:

```text
FROM-TO
```

where:

- `FROM` and `TO` are positive integers.
- The range is **inclusive**, so any ID `id` such that `FROM <= id <= TO` is considered fresh.

Example lines:

```text
1-10
50-75
492359630059500-492359630059600
```

Each such line describes a contiguous block of fresh ingredient IDs.

---

## Compaction of fresh ranges

Multiple fresh ranges can overlap or touch each other. For example:

- `1-5` and `4-10` overlap → union is `1-10`.
- `20-25` and `26-30` are touching (no gap) → they may be treated as a single continuous range `20-30`.

The **compaction** step merges such ranges into a smaller set of disjoint (and possibly larger) ranges. This has two main goals:

- Simplify reasoning about freshness.
- Represent the set of fresh IDs without redundancy.

At a high level, the compaction logic (as implemented in `FreshIngredients.Compact`) repeatedly:

1. Scans the list of current ranges.
2. Looks for pairs of ranges that:
   - Overlap, or
   - Touch exactly at a boundary (e.g. `r1.To + 1 == r2.From`).
3. Replaces those two ranges with a new range that covers their union.
4. Removes the old ranges from the list.
5. Continues until a full pass finds no more ranges to merge.

After compaction, `FreshIngredients` holds a minimal set of ranges that cover the same IDs as the original input ranges.

The total number of fresh ingredient IDs is then the sum over all compacted ranges of `(To - From + 1)`.

---

## High-level flow

The executable entrypoint is `day_5/cmd/main.go`. The high-level steps are:

1. **Read input file path** from the command‑line arguments.
2. **Initialize the ingredients reader** with that file path via `io.NewReader`.
3. Call `reader.Read()` to parse all fresh ranges into a `FreshIngredients` value.
4. Call `FreshIngredients.Compact()` to merge overlapping/touching ranges.
5. Call `FreshIngredients.Count()` to compute how many distinct ingredient IDs are fresh.
6. Print the final result:

   ```text
   Number of fresh ingredients: <count>
   ```

---

## Packages and responsibilities

### `cmd/main.go`

- Reads the input path from `os.Args`.
- Calls `initializeReader` to build an `IngredientsReader` from the input file.
- Calls `reader.Read()` to obtain a `*FreshIngredients` value.
- Calls `FreshIngredients.Compact()` to compact the fresh ranges.
- Calls `FreshIngredients.Count()` to compute the total number of fresh IDs.
- Prints the final count:

  ```go
  fmt.Printf("Number of fresh ingredients: %d\n", freshIngredientsCount)
  ```

### `internal/abstractions`

Defined in `ingredients_range.go`:

- `IngredientId` – type alias for `uint64` representing an ingredient ID.

- `IngredientRange` – represents a contiguous inclusive range of IDs:

  ```go
  type IngredientRange struct {
      From IngredientId
      To   IngredientId
  }
  ```

  - `IsIncluded(id IngredientId) bool` returns `true` if `From <= id <= To`.
  - `Count() uint64` returns how many IDs are in the range: `To - From + 1`.

- `FreshIngredients` – holds all fresh ranges:

  ```go
  type FreshIngredients struct {
      Ranges []IngredientRange
  }
  ```

  - `IsFresh(id IngredientId) bool` returns `true` if the given `id` is included in **any** of the stored ranges.
  - `Compact() *FreshIngredients` merges overlapping or touching ranges into a smaller set of ranges that cover the same IDs and returns a new `FreshIngredients` instance.
  - `Count() uint64` returns the total number of distinct fresh IDs represented by all ranges (by summing `Count()` over each range).

### `internal/io`

Defined in `ingredients_reader.go`.

- `IngredientsReader` encapsulates reading fresh ranges from the input file:

  ```go
  type IngredientsReader struct {
      inputFile *os.File
  }
  ```

- `NewReader(filePath string)`:
  - Opens the input file read-only.
  - Returns an `IngredientsReader` bound to that file.

- `Read() (*abstractions.FreshIngredients, error)`:
  - Uses a `bufio.Scanner` to read the file line by line.
  - For each non-empty line:
    - Splits the line on `"-"` into `FROM` and `TO`.
    - Parses both as `uint64` using `strconv.ParseUint`.
    - Ensures `FROM <= TO`, otherwise returns an error.
    - Converts them to `IngredientId` and appends an `IngredientRange` to the internal slice of ranges.
  - Stops scanning when it encounters a blank line (if present).
  - Returns a `*FreshIngredients` containing all parsed ranges.

---

## Running Day 5

From the `day_5` directory you can run the solution with:

```bash
cd day_5
make run ARGS="input.txt"
```

or directly with Go:

```bash
go run ./cmd input.txt
```
