# Day 10: Factory - Advent of Code 2025

## Problem Description

### Part 2: Joltage Levels (⭐)
Now we need to configure joltage counters instead of lights. Each machine has:
- **Joltage counters** that start at 0
- **Target values** for each counter (e.g., `{3,5,4,7}`)
- **Buttons** that increment specific counters by 1 (e.g., `(1,3)` increments counters 1 and 3)

The goal is to find the **minimum number of button presses** needed to reach the exact target values for all counters on all machines.

**Example:**
```
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
```
Minimum: 10 presses (one way: press `(3)` once, `(1,3)` 3×, `(2,3)` 3×, `(0,2)` once, `(0,1)` 2×)

**Part 2 Answer:** 19857 button presses

---

## Solution Approach

This problem is a **system of linear equations** problem. For Part 2:

### Mathematical Formulation

For each machine with buttons and target joltage values, we need to solve:
```
Button1_presses × Button1_counters + Button2_presses × Button2_counters + ... = Target_values
```

**Example:** Machine with buttons `(3)`, `(1,3)`, `(2)`, `(2,3)`, `(0,2)`, `(0,1)` and target `{3,5,4,7}`:

```
x₁ · (0,0,0,1) + x₂ · (0,1,0,1) + x₃ · (0,0,1,0) + x₄ · (0,0,1,1) + x₅ · (1,0,1,0) + x₆ · (1,1,0,0) = (3,5,4,7)
```

This translates to:
```
x₅ + x₆ = 3    (counter 0)
x₂ + x₆ = 5    (counter 1)
x₃ + x₄ + x₅ = 4    (counter 2)
x₁ + x₂ + x₄ = 7    (counter 3)
```

### Algorithm: Hermite Normal Form (HNF)

We solve this using the **Hermite Normal Form** algorithm, which consists of several steps:

#### 1. **Row Reduction to HNF**
Transform the augmented matrix into Hermite Normal Form using row operations:
- Find pivot elements (leading non-zero entry in each row)
- Eliminate entries below and above each pivot
- Ensure pivots are positive and increase left-to-right

**Implementation:** `RowReduction.ToHermiteNormalForm()`

#### 2. **Detect Free Variables**
Identify variables (i.e. rows) that don't have a leading 1 in any row. These variables can take any value, leading to infinitely many solutions.

**Implementation:** `HermiteNormalForm.findFreeVariables()`

#### 3. **Find Solutions**

**For unique solutions:**
- Use back-substitution starting from the last row
- Each variable is uniquely determined by the equations

**Implementation:** `HermiteNormalForm.getUniqueSolution()`

**For systems with free variables:**
- Try all combinations of free variable values starting from 0
- For each combination, solve for the remaining variables
- Check if all values are non-negative integers
- Return the first valid solution found (minimal solution)

**Implementation:** `HermiteNormalForm.findMinimalSolution()`

#### 4. **Optimization: Early Termination**
During the search for minimal solutions:
- Use modulo operator to check if divisions result in integers before computing
- Reject combinations that lead to negative values early
- Stop at the first valid solution (since we search from the smallest values)

---

## Code Structure

```
day_10/
├── cmd/
│   └── main.go                    # Entry point, reads input and solves
├── internal/
│   ├── abstractions/
│   │   ├── button.go              # Button and ButtonGroup structures
│   │   ├── counter.go             # Joltage counter abstraction
│   │   ├── factory.go             # Factory containing machines
│   │   ├── machine.go             # Machine with buttons and counters
│   │   ├── matrix.go              # Matrix operations and structures
│   │   ├── augmented_matrix.go    # Augmented matrix creation from machines
│   │   └── variable.go            # Variable and Variables for solutions
│   ├── algorithms/
│   │   ├── row_reduction.go       # Hermite Normal Form reduction
│   │   ├── hermite_normal_form.go # Solution finder for HNF systems
│   │   └── *_test.go              # Comprehensive unit tests
│   └── io/
│       └── factory_reader.go      # Parse input file
└── README.md                       # This file
```

---

## Key Algorithms Implemented

### 1. Row Reduction (Hermite Normal Form)
```go
func (rr *RowReduction) ToHermiteNormalForm() *abstractions.Matrix
```
Transforms a matrix into HNF using:
- Pivot selection and swapping
- Row operations to eliminate elements
- Adjustment for negative remainders (Euclidean division)

### 2. Free Variable Detection
```go
func (hnf *HermiteNormalForm) findFreeVariables() []abstractions.VariableNumber
```
Scans each row to find which columns have leading 1s, then identifies missing variables.

### 3. Unique Solution Solver
```go
func (hnf *HermiteNormalForm) getUniqueSolution() *abstractions.Variables
```
Uses back-substitution to solve systems with no free variables.

### 4. Minimal Solution Finder
```go
func (hnf *HermiteNormalForm) findMinimalSolution(freeVars []abstractions.VariableNumber, ...) *abstractions.Variables
```
Recursively tries combinations of free variable values to find the minimal valid solution.

---

## Example Walkthrough

### Input
```
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
```

### Step 1: Build Augmented Matrix
```
[  0  0  0  0  1  1  |  3 ]
[  0  1  0  0  0  1  |  5 ]
[  0  0  1  1  1  0  |  4 ]
[  1  1  0  1  0  0  |  7 ]
```

### Step 2: Reduce to HNF
```
[  1  0  0  1  0  -1  |  2 ]
[  0  1  0  0  0   1  |  5 ]
[  0  0  1  1  0  -1  |  1 ]
[  0  0  0  0  1   1  |  3 ]
```

### Step 3: Detect Free Variables
Variables with leading 1: x₁, x₂, x₃, x₅
Free variables: **x₄, x₆**

### Step 4: Find Minimal Solution
Try combinations of (x₄, x₆) starting from (0, 0):
- x₆ = 0: leads to negative values ❌
- x₆ = 1, x₄ = 1: x₁=2, x₂=4, x₃=1, x₅=2 ❌ (doesn't satisfy)
- Continue searching...
- Eventually find: x₁=1, x₂=3, x₃=0, x₄=3, x₅=1, x₆=2

**Total presses:** 1+3+0+3+1+2 = **10**

---

## Running the Solution

```bash
# Build and run
make run

# Run tests
go test ./...

# Run specific test
go test -v ./internal/algorithms -run TestHermiteNormalForm
```

---

## Complexity Analysis

- **Time Complexity:** O(m²n) for HNF reduction, where m = rows, n = columns
- **Space Complexity:** O(mn) for matrix storage
- **Solution Finding:** Exponential in the worst case (number of free variables), but optimized with early termination

---

## References

- [Hermite Normal Form](https://en.wikipedia.org/wiki/Hermite_normal_form)
- [Gaussian Elimination](https://en.wikipedia.org/wiki/Gaussian_elimination)
- [Systems of Linear Equations](https://en.wikipedia.org/wiki/System_of_linear_equations)

