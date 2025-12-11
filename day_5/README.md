# Day 5 – Fresh Ingredients (Part 1 & Part 2)

The Day 5 solutions are implemented in Go in the `day_5` folder. They read a list of **freshness ranges** for ingredient IDs and a list of **available ingredient IDs**, then determine how many available ingredients are considered **fresh**.

- **Part 1**: Count how many available ingredient IDs fall inside at least one fresh range.
- **Part 2**: First **compact/merge** overlapping or touching fresh ranges into a minimal set of ranges, then count how many available ingredient IDs fall into these compacted ranges.

## Problem model (shared by Part 1 & Part 2)

The input file is split into two sections, separated by a blank line:

1. **Fresh ingredient ranges** (top section)
2. **Available ingredient IDs** (bottom section)

### Fresh ingredient ranges

- Each non-empty line in the first section represents a **range of fresh ingredient IDs** in the form:

  ```text
  FROM-TO
  ```

  where:

  - `FROM` and `TO` are positive integers.
  - The range is **inclusive**, so any ID `id` such that `FROM <= id <= TO` is considered fresh.

- Example lines:

  ```text
  1-10
  50-75
  492359630059500-492359630059600
  ```

### Available ingredient IDs

- After a blank line, the second section lists **one ingredient ID per line**:

  ```text
  5
  12
  60
  492359630059532
  ```

- These are the ingredient IDs that are currently available.

### Fresh vs non-fresh

An available ingredient ID is considered **fresh** if it falls inside **at least one** of the fresh ranges defined in the top section (after compaction for Part 2).

- Example:
  - Fresh ranges:
    - `1-10`
    - `50-75`
  - Available IDs:
    - `5` → fresh (in `1-10`)
    - `12` → not fresh
    - `60` → fresh (in `50-75`)

For Part 1, these ranges are used **as-is**. For Part 2, they may be compacted/merged first (e.g. `1-5` and `4-10` become a single range `1-10`).

The program looks at every available ingredient ID and counts how many are fresh. That count is then printed as the answer.

---

## Part 1 – Basic fresh count

For Part 1, the logic is straightforward:

1. Parse all fresh ranges and available IDs from the input.
2. For each available ID, check whether it falls into **any** of the fresh ranges.
3. Count all available IDs that are fresh.
4. Print:

   ```text
   Number of fresh ingredients: <count>
   ```

This corresponds to calling `FreshIngredients.IsFresh(id)` directly on the raw set of ranges.

---

## Part 2 – Compacting fresh ranges before counting

Part 2 builds on Part 1 by first **compacting** the list of fresh ranges before evaluating which available IDs are fresh.

Multiple fresh ranges can overlap or touch each other. For example:

- `1-5` and `4-10` overlap → union is `1-10`.
- `20-25` and `26-30` are touching (no gap) → they may be treated as a single continuous range `20-30`.

The compaction step merges such ranges into a minimal set of disjoint (and possibly larger) ranges. This has two benefits:

- It simplifies subsequent membership checks.
- It ensures that the set of fresh IDs is represented without redundancy.

The high-level algorithm for compaction (as implemented in `FreshIngredients.Compact`) is:

1. Take the current list of ranges.
2. Repeatedly look for pairs of ranges that:
   - Overlap, or
   - Touch exactly at a boundary (e.g. `r1.To + 1 == r2.From`).
3. Replace those two ranges with a new range that covers their union.
4. Continue until no more merges are possible.

Once compacted, Part 2 then proceeds like:

- Count and print the total number of fresh available IDs.

Depending on how the puzzle is phrased, Part 2 may either:

- Produce the **same** numerical answer as Part 1 but with more efficient/cleaner internal representation, or
- Change the semantics slightly (e.g. by treating touching ranges as continuous), which can alter which IDs are considered fresh.

---

## High-level flow

The executable entrypoint is `day_5/cmd/main.go`. The high-level steps are:

1. **Read input file path** from the command‑line arguments.
2. **Initialize the ingredients parser** with that file path via `NewParser`.
3. The parser reads the file and builds:
   - `FreshIngredients` – the list of fresh ID ranges.
   - `AvailableIngredients` – the list of available IDs.
4. For Part 2, call the compaction step on `FreshIngredients` (e.g. `Compact`) before checking freshness.
5. In `main`, iterate over all available IDs and use `FreshIngredients.IsFresh(id)` to check each one.
6. Count how many available IDs are fresh.
7. Print the final result:

   ```text
   Number of fresh ingredients: <count>
   ```

---

## Packages and responsibilities

### `cmd/main.go`

- Reads the input path from `os.Args`.
- Calls `initializeParser` to build an `IngredientsParser` from the input file.
- For Part 2, can invoke the compaction step on `ingredientsParser.Fresh` before counting.
- Iterates over `ingredientsParser.Available.Ids`:
  - For each `ingredientId`, calls `ingredientsParser.Fresh.IsFresh(ingredientId)`.
  - Increments `freshIngredientsCount` for each ID that is fresh.
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

- `AvailableIngredients` – holds all available IDs:

  ```go
  type AvailableIngredients struct {
      Ids []IngredientId
  }
  ```

- `FreshIngredients` – holds all fresh ranges:

  ```go
  type FreshIngredients struct {
      Ranges []IngredientRange
  }
  ```

  - `IsFresh(id IngredientId) bool` returns `true` if the given `id` is included in **any** of the stored ranges.
  - `Compact()` (Part 2) merges overlapping or touching ranges into a smaller set of ranges that cover the same IDs.

### `internal/parser`

Defined in `ingredients_parser.go`.

- `IngredientsParser` groups together the two main abstractions:

  ```go
  type IngredientsParser struct {
      Fresh     *abstractions.FreshIngredients
      Available *abstractions.AvailableIngredients
  }
  ```

- `NewParser(filePath string)`:
  - Opens the input file.
  - Delegates to `readIngredients(filePath)` to build:
    - `FreshIngredients` with all the ranges from the first section.
    - `AvailableIngredients` with all the IDs from the second section.
  - Returns an `IngredientsParser` that exposes these two structures.

- `readIngredients(filePath string)`:
  - Opens the file using `os.OpenFile` and creates a `bufio.Scanner`.
  - Initializes:
    - `freshIngredients` with an empty `Ranges` slice (preallocated capacity for performance).
    - `availableIngredients` with an empty `Ids` slice.
  - Maintains a `processingSection` state to know whether it’s reading:
    - `freshIngredientsSection` – before the blank line.
    - `availableIngredientsSection` – after the blank line.
  - For each scanned line:
    - Skips empty/whitespace-only lines and switches to the **available ingredients** section when the first blank line is encountered.
    - In the **fresh ranges** section:
      - Splits the line on `"-"` into `FROM` and `TO`.
      - Parses both as `uint64` using `strconv.ParseUint`.
      - Validates that `FROM <= TO`; otherwise returns an error.
      - Converts them to `IngredientId` and appends an `IngredientRange` to `freshIngredients.Ranges`.
    - In the **available IDs** section:
      - Parses the line as a `uint64` ID.
      - Converts it to `IngredientId` and appends it to `availableIngredients.Ids`.
  - Returns pointers to the populated `FreshIngredients` and `AvailableIngredients`.

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
