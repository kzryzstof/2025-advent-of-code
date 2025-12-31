# Day 12 – Present Packing Under Christmas Trees

This day models a **2D packing** problem: you’re given a fixed set of weird 3×3 present silhouettes (cells are either occupied `#` or empty `.`), and for each Christmas tree you’re given a rectangular region plus a shopping list of how many presents of each shape must fit.

A placement is valid if:

- Presents are aligned on the unit grid.
- Presents can be **rotated** and **flipped**.
- Occupied cells may not overlap.
- Holes (`.`) in a shape do **not** block other presents.

The program counts how many trees’ regions can be fully packed with their required presents.

---

## High-level approach

The implementation uses two main ideas:

1. **Precompute “good” pairwise combinations of shapes** (under rotations / flips / limited vertical shifts) to get higher-density pieces.
2. **Greedily place** the densest combined pieces first, then place leftover single presents.

This isn’t a full exact tiling/ILP solver. It’s a pragmatic approach that performs well for the puzzle input while keeping the code approachable.

---

## Input parsing

Parsing is handled by `internal/io/cavern_reader.go`.

- Each shape is read as 3 lines of `#` / `.`, converted into a 3×3 `[][]int8` grid.
- Empty cells use a sentinel constant `E = -99` (not `0`, because present index `0` is a valid ID).
- Each tree line is parsed as:

  ```text
  <wide>x<long>: c0 c1 c2 c3 c4 c5
  ```

  and becomes a `ChristmasTree` with a `Region` of that size and a map of required counts by present index.

---

## Core abstractions

### Shapes and density

`internal/abstractions/shape.go`

A `Shape` is:

- `Cells [][]int8` – the 2D grid.
- `Dimension` – `Wide` and `Long`.
- `FillRatio` – occupancy ratio computed by `ComputeFillRatio`.

`FillRatio` is the main heuristic measure of “goodness”: denser shapes are usually easier to place without leaving unusable holes.

### Region (the space under a tree)

`internal/abstractions/region.go`

A `Region` owns the backing grid (`space`) and tracks its dimensions. The packing algorithm writes placed shapes into this grid.

### ChristmasTree present configuration

`internal/abstractions/christmas_tree.go`

Each tree stores remaining counts in `PresentConfiguration` objects, sorted descending by count (helpful for iterating and updating).

---

## Precomputing combined shapes (pairwise permutations)

`internal/abstractions/shape_permutations.go`

Before packing any tree, the code precomputes how two present shapes can be “packed together” into a single composite shape.

### Transformations (rotations / flips)

The algorithm applies a sequence of in-place 3×3 operations:

- `RotateClockwise`
- `VerticalFlip`
- `HorizontalFlip`
- `NoOp`

This generates multiple orientations.

### Combining two shapes

`internal/abstractions/shape_packing.go` implements:

- `PackShapes(fixedShapeID, fixedShape, movingShapeID, movingShape, slideOffset, verbose) Shape`

Process:

1. Optionally **slide** the moving shape down by `slideOffset` using `SlideShape`.
2. Compute how much the moving shape can be shifted left without overlap using `computeColOffset`.
3. Create a new canvas sized to hold both shapes.
4. Paste the fixed shape at `(0,0)` and the moving shape at the computed offset using `PasteShape`.
5. Return a new `Shape` with updated `Dimension` and `FillRatio`.

### Choosing the “best” combination

`CombinationCatalog` (`internal/abstractions/combination_catalog.go`) keeps, for each left present index, the combinations with other presents and retains the most “optimal” one based on `FillRatio`.

This gives the packing stage a menu of “dense composite pieces” it can try to place first.

---

## Placing shapes into a region

`internal/abstractions/shape_packing.go` also contains:

- `PackShape(region, shape, verbose) bool`

It attempts to place the given shape into the region grid:

- It searches for an insert position (`findInsertPosition`) where all occupied cells of the shape fit in empty cells of the region.
- On success, it writes the shape cells into the region and returns `true`.
- On failure, it returns `false`.

(There are debug helpers like `PrintShape`/`PrintShapes` to visualize intermediate states when `verbose` is enabled.)

---

## Packing strategy per Christmas tree

`internal/abstractions/cavern.go` orchestrates the full solution.

For each tree:

1. Build the permutation catalog once for all presents:
   - `catalog := ComputePermutations(c.presents, verbose)`
2. For the current tree region, **place combined shapes first**:
   - Iterate present indices in descending fill-ratio order.
   - For each present type, try its best combinations (from the catalog).
   - Decrement both involved present counts as combinations are placed.
3. Then **place remaining individual presents** one by one.
4. If any required present can’t be placed, this tree is counted as **failed**.

At the end, the program prints how many trees could be fully packed.

---

## Entry point and running

The executable is `cmd/main.go`:

- Reads the input file path from `os.Args`.
- Uses `CavernReader` to parse presents + trees.
- Runs `cavern.PackAll(false)`.

### Commands

From the `day_12` directory:

```bash
make test
make run ARGS="input.txt"
```

---

## Code layout

```text
day_12/
├── cmd/
│   └── main.go                     # Entry point
├── internal/
│   ├── io/
│   │   └── cavern_reader.go        # Input parsing
│   └── abstractions/
│       ├── cavern.go               # Orchestrates packing for all trees
│       ├── christmas_tree.go       # Tree + required counts
│       ├── region.go               # 2D region grid
│       ├── shape.go                # Shape + fill ratio heuristic
│       ├── shape_permutations.go   # Precompute pairwise packed shapes
│       ├── shape_packing.go        # PackShapes + PackShape placement
│       ├── combination_catalog.go  # Stores best combinations per shape
│       └── slice_extensions.go     # 3×3 transforms + helpers
└── README.md
```

