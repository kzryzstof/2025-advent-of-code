# Day 8 – Junction Box Circuits

The Day 8 solution is implemented in Go in the `day_8` folder. It reads a playground of junction boxes in 3D space, connects them using a limited number of cables by pairing the closest boxes first, and identifies the three biggest resulting circuits.

## Problem model

- The input file contains **junction boxes** represented as 3D coordinates.
  - Each line has the format: `x,y,z` (e.g., `75262,24842,97390`).
- Each junction box starts as an independent **circuit** (containing only itself).
- The elves have a **limited number of cables** (1000 in the main program) to connect junction boxes.
- When a pair of junction boxes is processed:
  - If they belong to **different circuits**, those circuits are **merged** into one larger circuit and a cable is consumed.
  - If they already belong to the **same circuit**, the pair is logged but no merge happens — however, the available cable counter is still decremented by one.
- The connection priority is **greedy by distance**:
  - All possible pairs of junction boxes are evaluated and sorted by their Euclidean distance in 3D space.
  - Pairs are processed in order from shortest to longest distance, until the available cable counter reaches zero or all pairs are exhausted.
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
   - Iterate through the ordered pairs while cables are still available:
     - Look up which circuits currently contain the two boxes.
     - Depending on their circuits, either merge circuits or do nothing.
     - Decrement the remaining cable counter for each processed pair.
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
    - `NewCircuit(junctionBox)`: Creates a circuit containing a single junction box.
    - `Add(junctionBox)`: Adds a junction box to the circuit.
    - `Contains(junctionBox)`: Checks if a junction box is in this circuit.
    - `IsDisconnected()`: Returns true if the circuit contains only one box.
    - `Count()`: Returns the number of junction boxes in the circuit.
    - `Get()`: Returns the slice of junction boxes in the circuit.
  - **`Circuits`**: Manages all circuits in the playground.
    - `NewCircuits()`: Constructs the collection of circuits.
    - `Add(circuit)`: Adds a new circuit to the collection.
    - `Remove(circuit)`: Removes a circuit (used after merging).
    - `Get(junctionBox)`: Finds which circuit contains a given junction box.
    - `Merge(fromCircuit, toCircuit)`: Merges two circuits together, moving all junction boxes from `fromCircuit` into `toCircuit` and removing `fromCircuit`.
    - `Connect(fromCircuit, toCircuit)`: Alias for `Merge`.
    - `GetBiggestCircuits(count)`: Returns the N biggest circuits sorted by size (largest first).
    - `GetAll()`: Returns all circuits.
    - `Count()`: Returns the total number of circuits.

- `internal/app`
  - **`pair_junction_boxes.go`**:
    - `Pair(junctionBoxes)`: Generates all possible pairs of junction boxes.
    - For each pair, computes and stores the Euclidean distance.
    - Returns a slice of unordered `JunctionBoxPair` values.
  - **`connect_junction_boxes.go`**:
    - `ConnectJunctionBoxes(playground, availableConnectionsCount, verbose)`: Main algorithm.
    - Creates initial circuits (one per junction box).
    - Generates and sorts all pairs by distance using `abstractions.Order`.
    - Iterates through the ordered pairs while `availableConnectionsCount > 0`:
      - Retrieves `circuitA` and `circuitB` using `Circuits.Get`.
      - Handles three main scenarios:
        1. **Both boxes are already in the same circuit** → log and do not merge, but still decrement the cable counter.
        2. **One box is in a single-box circuit** (`IsDisconnected()`) → connect that circuit into the other circuit using `Connect`.
        3. **Both boxes are in different multi-box circuits** → merge the two circuits using `Merge`.
      - After handling a pair, decrement `availableConnectionsCount`.
    - Returns the final `Circuits` collection.

## Algorithm details

### Pairing strategy
The `Pair` function creates all possible combinations of junction boxes (n choose 2), computing the distance for each pair. This results in `n * (n - 1) / 2` pairs.

### Connection strategy
The algorithm is **greedy by shortest distance**, but it is not a perfect minimum spanning tree implementation:
1. Sort all pairs by distance (ascending).
2. Process pairs in order from shortest to longest.
3. For each pair, determine the current circuits of both boxes and either merge circuits or skip merging.
4. Decrement the available cable counter for every processed pair, even if no merge occurs.
5. Stop when the cable counter reaches zero or all pairs have been processed.

Because cables are considered "used" for every processed pair, even when the two boxes are already in the same circuit, some cables are effectively wasted and the final number of circuits can be higher than the theoretical minimum.

### Circuit merging
When connecting two junction boxes that belong to different circuits:
- All junction boxes from the source circuit are appended to the target circuit.
- The source circuit is removed from the `Circuits` collection.
- Over time, this reduces the number of circuits as more connections are made.

## Example

For `N` junction boxes and `K` cables:
- Start: `N` circuits (one per box).
- Each **successful merge** reduces the circuit count by 1.
- Each **processed pair** (whether it triggers a merge or not) consumes 1 cable.
- End: You will typically have between `N - K` and `N` circuits, depending on how many merges were actually performed before running out of connections.

The final product of the three biggest circuit sizes gives the puzzle answer.

## Running the solution

```bash
cd day_8
go run ./cmd ./input.txt
```

This compiles and executes the program with `input.txt` as the input file.
