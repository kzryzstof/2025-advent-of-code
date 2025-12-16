# Day 10 – Factory Machine Activation

The Day 10 solution is implemented in Go in the `day_10` folder. It models a factory that contains multiple machines. Each machine has:

- A row of lights that can be **on** (`#`) or **off** (`.`).
- Several button groups, where each button toggles one specific light.

The goal is to find, for every machine, the **minimum number of button presses** needed to turn on all the required lights (activate the machine), and then sum these minimal press counts across all machines.

## Problem model

- The input describes one or more machines, one per line.
- Each line encodes:
  - The **lights** state as one or more bracketed groups of `.` and `#`, e.g. `[#..#.]`.
  - The **button groups** as one or more parenthesized lists of comma-separated indices, e.g. `(0,2,4)`.
- A **Light** is modeled by its state (`on` or `off`).
- A **Button** references a light index; pressing it toggles that light.
- A **ButtonGroup** is a collection of buttons that are all pressed together.
- A **Machine** consists of:
  - A slice of lights.
  - A slice of button groups.
  - Logic to press groups and check whether it is activated (all required lights on).
- A **Factory** is a collection of machines.

The objective is to determine the minimal sequence length of button-group presses that activates each machine, then sum all those lengths.

## High-level flow

The executable entrypoint is `day_10/cmd/main.go` and the flow is:

1. **Read input path** from the command-line arguments.
2. **Initialize the reader** with `io.NewReader(path)`.
3. **Parse the factory**:
   - For each line, construct a `Machine` by parsing:
     - Lights between `[` and `]` into a slice of `Light` values.
     - Button groups between `(` and `)` into `ButtonGroup` values referencing light indices.
   - Collect all machines into a `Factory`.
4. **Activate machines** with `app.ActivateMachines(factory)`:
   - For each machine, call `FindShortestCombinations` to search for the smallest number of button-group presses that activates it.
   - The search tries combinations of button-group indices of increasing length.
   - For each candidate combination, it:
     - Resets/closes lights on the machine.
     - Presses the button groups in the candidate sequence.
     - Tests whether the machine is activated.
   - Once a working combination is found, its length (press count) is returned.
5. **Aggregate and print results**:
   - Sum the minimal press counts across all machines.
   - Print the number of machines, the total presses, and the execution time.

## Combination search

The heart of the solution is `FindShortestCombinations` in `internal/app/combinations.go`:

- It receives:
  - `maximumCombinationLength`: the upper bound on combination size to consider.
  - `totalButtonGroupsCount`: how many button groups exist.
  - `testCombination`: a callback that applies a candidate combination on the machine and returns `true` if it activates the machine.
- It uses a recursive helper to build up combinations of button-group indices without duplicates.
- For each target length `k` from `1` up to `maximumCombinationLength`:
  - It recursively generates all unique combinations of size `k`.
  - For each combination, it calls `testCombination`.
  - As soon as one combination succeeds, it returns `k` and `true`.
- If no combination up to the maximum length works, it returns `-1` and `false`.

This effectively performs a breadth-first search over combination lengths, ensuring the first found activation sequence is minimal in terms of total button-group presses.

## Packages and responsibilities

- `cmd/main.go`
  - Wires together input reading, factory construction, and machine activation.
  - Prints summary information and execution duration.

- `internal/io`
  - `FactoryReader` reads the input file and converts each line into an `abstractions.Machine`:
    - `extractLightIndicators` parses bracketed `[#..#.]` segments into a slice of `Light` values.
    - `extractButtons` parses parenthesized `(0,2,4)` segments into `ButtonGroup` values.

- `internal/abstractions`
  - Core domain types:
    - `Light` with on/off state and toggle behavior.
    - `Button` referencing a light index.
    - `ButtonGroup` containing multiple buttons.
    - `Machine` wrapping lights and button groups with methods like `PressGroup` and `IsActivated`.
    - `Factory` as a container for machines.
  - Utility functions like `Contains` and `Clear` for working with slices.

- `internal/app`
  - `ActivateMachines` orchestrates activation across all machines in the factory, summing up the minimal press counts.
  - `FindShortestCombinations` implements the core combinatorial search used to find the shortest activating sequence of button groups for a single machine.

