# Day 5 – Fresh Ingredients (Part 1)

The Day 5 Part 1 solution is implemented in Go in the `day_5` folder. It reads a list of **freshness ranges** for ingredient IDs and a list of **available ingredient IDs**, then counts how many available ingredients are considered **fresh**.

## Problem model (Part 1)

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

An available ingredient ID is considered **fresh** if it falls inside **at least one** of the fresh ranges defined in the top section.

- Example:
  - Fresh ranges:
    - `1-10`
    - `50-75`
  - Available IDs:
    - `5` → fresh (in `1-10`)
    - `12` → not fresh
    - `60` → fresh (in `50-75`)

The program looks at every available ingredient ID and counts how many are fresh. That count is then printed as the Part 1 answer.

---

## High-level flow

The executable entrypoint is `day_5/cmd/main.go`. The high-level steps are:

1. **Read input file path** from the command‑line arguments.
2. **Initialize the ingredients parser** with that file path via `NewParser`.
3. The parser reads the file and builds:
   - `FreshIngredients` – the list of fresh ID ranges.
   - `AvailableIngredients` – the list of available IDs.
4. In `main`, iterate over all available IDs and use `FreshIngredients.IsFresh(id)` to check each one.
5. Count how many available IDs are fresh.
6. Print the final result:

   ```text
   Number of fresh ingredients: <count>
   ```

---

## Packages and responsibilities

### `cmd/main.go`

- Reads the input path from `os.Args`.
- Calls `initializeParser` to build an `IngredientsParser` from the input file.
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

  - `IsFresh(id IngredientId) bool` returns `true` if the given `id` is included in **any** of the stored ranges:

    ```go
    func (f FreshIngredients) IsFresh(id IngredientId) bool {
        for i := range f.Ranges {
            if f.Ranges[i].IsIncluded(id) {
                return true
            }
        }
        return false
    }
    ```

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

## Running Day 5 Part 1

From the `day_5` directory you can run the solution with:

```bash
cd day_5
make run ARGS="input.txt"
```

or directly with Go:

```bash
go run ./cmd input.txt
```

