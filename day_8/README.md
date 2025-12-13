# Day 8 – Junction Box Circuits

The Day 8 solution is implemented in Go in the `day_8` folder. It reads a playground of junction boxes in 3D space, connects them using a limited number of cables by pairing the closest boxes first, and identifies the three biggest resulting circuits.

## Problem model

- The input file contains **junction boxes** represented as 3D coordinates.
  - Each line has the format: `x,y,z` (e.g., `75262,24842,97390`).
- Each junction box starts as an independent **circuit** (containing only itself).
- The elves have a **limited number of cables** (1000 in the main program) to connect junction boxes.
- When two junction boxes are connected:
  - If they belong to different circuits, those circuits are **merged** into one larger circuit.
  - If they already belong to the same circuit, the connection is skipped (no cable used).
- The connection strategy is **greedy by distance**:
  - All possible pairs of junction boxes are evaluated and sorted by their Euclidean distance in 3D space.
  - Pairs are connected in order from shortest to longest distance, until all available cables are used.
- At the end, the program identifies the **three biggest circuits** (by number of junction boxes) and computes the product of their sizes.

## High-level flow

The executable entrypoint is `day_8/cmd/main.go` and the flow is:

1. **Read input file path** from the command-line arguments.
2. **Initialize the playground reader** to parse the input file.
3. **Read all junction boxes** from the file into a `Playground` structure.
4. **Connect junction boxes** using the `ConnectJunctionBoxes` function:
   - Create initial circuits (one per junction box).
   - Generate all possible pairs of junction boxes with their distances.
   - Sort pairs by distance (shortest first).
   - Iterate through pairs, connecting them if:
     - Cables are still available.
     - The two boxes belong to different circuits.
   - Merge circuits when connections are made.
5. **Find the three biggest circuits** by junction box count.
6. Print the results:
   - Total number of circuits created.
   - The sizes of the three biggest circuits.
   - The product of those three sizes.

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the playground reader and parses the input.
  - Calls `app.ConnectJunctionBoxes` with the playground and cable count.
  - Identifies the three biggest circuits and prints their product.

- `internal/io`
  - `PlaygroundReader` opens the input file and scans it line by line.
  - Each line is parsed as three comma-separated integers: `x,y,z`.
  - Creates a `JunctionBox` at each position.
  - Returns a `Playground` containing all junction boxes.

- `internal/abstractions`
  - **`Playground`**: Container for all junction boxes in 3D space.
  - **`JunctionBox`**: Represents a single junction box with a 3D position.
    - `MeasureDistance(other)`: Computes the Euclidean distance to another junction box.
  - **`Position`**: Represents a 3D coordinate (x, y, z).
    - `Distance(other)`: Calculates the Euclidean distance between two positions.
  - **`JunctionBoxPair`**: Represents a pair of junction boxes with their computed distance.
    - Used for sorting and prioritizing connections.
  - **`Circuit`**: Represents a collection of connected junction boxes.
    - `Add(junctionBox)`: Adds a junction box to the circuit.
    - `Contains(junctionBox)`: Checks if a junction box is in this circuit.
    - `IsDisconnected()`: Returns true if the circuit contains only one box.
    - `Count()`: Returns the number of junction boxes in the circuit.
  - **`Circuits`**: Manages all circuits in the playground.
    - `Add(circuit)`: Adds a new circuit to the collection.
    - `Get(junctionBox)`: Finds which circuit contains a given junction box.
    - `Merge(fromCircuit, toCircuit)`: Merges two circuits together, removing the source circuit.
    - `Connect(fromCircuit, toCircuit)`: Alias for `Merge`.
    - `GetBiggestCircuits(count)`: Returns the N biggest circuits sorted by size.
    - `Count()`: Returns the total number of circuits.

- `internal/app`
  - **`pair_junction_boxes.go`**:
    - `Pair(junctionBoxes)`: Generates all possible pairs of junction boxes.
    - For each pair, computes and stores the Euclidean distance.
    - Returns a slice of unordered `JunctionBoxPair` values.
  - **`connect_junction_boxes.go`**:
    - `ConnectJunctionBoxes(playground, availableConnectionsCount, verbose)`: Main algorithm.
    - Creates initial circuits (one per junction box).
    - Generates and sorts all pairs by distance.
    - Iterates through pairs, connecting them when:
      - Cables are available.
      - The boxes belong to different circuits.
    - Handles three connection scenarios:
      1. One box is not yet connected → add it to the other's circuit.
      2. Both boxes are in different circuits → merge the circuits.
      3. Both boxes are already in the same circuit → skip (no cable used).
    - Returns the final `Circuits` collection.

## Algorithm details

### Pairing strategy
The `Pair` function creates all possible combinations of junction boxes (n choose 2), computing the distance for each pair. This results in `n * (n - 1) / 2` pairs.

### Connection strategy
The algorithm uses a **greedy minimum spanning tree** approach:
1. Sort all pairs by distance (ascending).
2. Connect pairs in order, similar to Kruskal's algorithm.
3. Stop when all cables are used or all boxes are in a single circuit.

### Circuit merging
When connecting two junction boxes:
- If they're in different circuits, merge the smaller circuit into the larger one.
- If they're already in the same circuit, skip without using a cable.
- Track which circuit each junction box belongs to for efficient lookups.

## Example

For 1000 junction boxes and 1000 cables:
- Start: 1000 circuits (one per box)
- Each cable used reduces the circuit count by 1 (when merging different circuits)
- End: Potentially 1 circuit (if all boxes are connected), or more if cables run out

The final product of the three biggest circuit sizes gives the puzzle answer.

## Running the solution

```bash
cd day_8
go run ./cmd ./input.txt
```

This compiles and executes the program with `input.txt` as the input file.

