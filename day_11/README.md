# Day 11: Reactor - Advent of Code 2025

## Problem Description

### The Challenge
Inside the factory's underground reactor room, there's a new server rack that won't communicate with the toroidal reactor. The issue is caused by data following specific paths through a network of interconnected devices.

Each device has:
- A unique **name** (e.g., `aaa`, `bbb`, `you`, `out`)
- Zero or more **outputs** connecting to other devices

**Example Input:**
```
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
```

This means:
- Device `aaa` has outputs to `you` and `hhh`
- Device `you` has outputs to `bbb` and `ccc`
- And so on...

### The Task
Find **all possible paths** from the device labeled `you` to the device labeled `out`.

**Important:** Data only flows forward through outputs; it cannot flow backwards.

### Example Solution
In the example above, there are **5 different paths** from `you` to `out`:

1. `you → bbb → ddd → ggg → out`
2. `you → bbb → eee → out`
3. `you → ccc → ddd → ggg → out`
4. `you → ccc → eee → out`
5. `you → ccc → fff → out`

---

## Solution Approach

This is a **graph traversal problem** where we need to count all possible paths between two nodes in a directed graph.

### Algorithm: Depth-First Search (DFS)

The solution uses a recursive depth-first search approach to explore all possible paths:

#### 1. **Build the Graph**
Parse the input and construct a directed graph where:
- Each device is a **node**
- Each output connection is a **directed edge**

**Implementation:** `BuildGraph()` in `graph.go`

```go
func BuildGraph(devices []*Device) *Graph {
    graph := &Graph{
        nodes:       []*Node{},
        nodesByName: map[string]*Node{},
    }
    
    for _, device := range devices {
        // Get or create node for this device
        deviceNode := graph.getNodeByName(device.name)
        if deviceNode == nil {
            deviceNode = graph.createNewNode(device.name)
        }
        
        // Connect to all output devices
        for _, outputName := range device.outputs {
            outputNode := graph.getNodeByName(outputName)
            if outputNode == nil {
                outputNode = graph.createNewNode(outputName)
            }
            deviceNode.AddNext(outputNode)
        }
    }
    
    return graph
}
```

#### 2. **Count Paths Using Recursive DFS**
Starting from the `you` node, recursively explore all outgoing edges:
- If we reach the `out` node, we've found a valid path (increment count)
- Otherwise, recursively explore all connected nodes
- Sum up all paths found from child nodes

**Implementation:** `CountPathsTo()` in `node.go`

```go
func (n *Node) CountPathsTo(to string) uint {
    if n.name == to {
        return 1  // Base case: reached destination
    }
    
    currentCount := uint(0)
    return n.testPaths(to, currentCount)
}

func (n *Node) testPaths(to string, currentCount uint) uint {
    if n.name == to {
        return currentCount + 1
    }
    
    // Explore all outgoing edges
    for _, nextNode := range n.next {
        currentCount = nextNode.testPaths(to, currentCount)
    }
    
    return currentCount
}
```

#### 3. **Optimization: Hash Map Lookup**
To efficiently find nodes by name, we maintain a `nodesByName` map:

```go
nodesByName: map[string]*Node{}
```

This provides O(1) lookup time instead of O(n) linear search.

---

## Example Walkthrough

### Input
```
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
```

### Step 1: Build Graph Structure
```
        you
       /   \
     bbb   ccc
     / \   /|\
   ddd eee ddd eee fff
    |   |       |   |
   ggg  out    out out
    |
   out
```

### Step 2: Traverse and Count Paths

Starting from `you`:
1. **Explore `you → bbb`:**
   - Explore `bbb → ddd → ggg → out` ✓ (path 1)
   - Explore `bbb → eee → out` ✓ (path 2)

2. **Explore `you → ccc`:**
   - Explore `ccc → ddd → ggg → out` ✓ (path 3)
   - Explore `ccc → eee → out` ✓ (path 4)
   - Explore `ccc → fff → out` ✓ (path 5)

**Total Paths:** 5

---

## Code Structure

```
day_11/
├── cmd/
│   └── main.go                    # Entry point, reads input and counts paths
├── internal/
│   ├── abstractions/
│   │   ├── device.go              # Device structure (name + outputs)
│   │   ├── node.go                # Graph node with DFS path counting
│   │   ├── node_test.go           # Comprehensive unit tests for Node
│   │   └── graph.go               # Graph builder and path counter
│   └── io/
│       └── devices_reader.go      # Parse input file
└── README.md                       # This file
```

---

## Key Data Structures

### Device
```go
type Device struct {
    name    string
    outputs []string
}
```
Represents a device from the input file.

### Node
```go
type Node struct {
    name string
    next []*Node  // Outgoing edges
}
```
Represents a node in the directed graph.

### Graph
```go
type Graph struct {
    nodes       []*Node
    nodesByName map[string]*Node  // Fast O(1) lookup
}
```
Manages the entire graph structure.

---

## Algorithm Analysis

### Time Complexity
- **Graph Building:** O(V + E) where V = devices, E = connections
- **Path Counting:** O(V + E) in the best case, but can be exponential in worst case if there are many paths
  - In graphs with cycles or many paths, this could be O(2^V)
  - For DAGs (Directed Acyclic Graphs) like this problem, it's more efficient

### Space Complexity
- **Graph Storage:** O(V + E) for nodes and edges
- **Recursion Stack:** O(V) for the depth of DFS traversal
- **Hash Map:** O(V) for node name lookup

### Optimization Opportunities
For large graphs with many paths, we could:
1. **Memoization:** Cache path counts for each node to avoid recomputation
2. **Dynamic Programming:** Bottom-up approach from destination to source
3. **Cycle Detection:** Prevent infinite loops if the graph has cycles

---

## Running the Solution

```bash
# Build and run
make run

# Run with specific input
go run cmd/main.go input.txt

# Run tests
go test ./...

# Run Node tests specifically
go test -v ./internal/abstractions -run TestNode
```

---

## Example Output

```
Reader initialized: &{devices_reader.go file}
The room has 9 devices. Graph has 5 path from 'you' to 'out'.
Execution time: 123.456µs
```

---

## Testing

The solution includes comprehensive table-driven unit tests for the Node struct:
- ✅ Node creation (`NewNode`)
- ✅ Adding child nodes (`AddNext`)
- ✅ Finding nodes by name (`FindNodeByName`)
- ✅ Path counting in various tree structures
- ✅ Edge cases (empty trees, non-existent nodes, deep nesting)

Run tests with: `go test -v ./internal/abstractions/node_test.go`

---

## Key Insights

1. **Graph Representation:** The problem is fundamentally about traversing a directed graph
2. **DFS is Natural:** Recursive DFS naturally explores all paths from source to destination
3. **No Cycle Handling Needed:** The problem description implies a DAG (acyclic), so we don't need cycle detection
4. **Hash Maps for Speed:** Using a map for node lookup provides significant performance improvement over linear search

---

## References

- [Depth-First Search](https://en.wikipedia.org/wiki/Depth-first_search)
- [Directed Acyclic Graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph)
- [Graph Traversal](https://en.wikipedia.org/wiki/Graph_traversal)
- [All-Pairs Shortest Paths](https://en.wikipedia.org/wiki/All-pairs_shortest_path_problem)

