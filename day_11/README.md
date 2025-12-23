# Day 11 – Reactor (Part 2)

This day models a directed graph of devices and connections inside a toroidal reactor. Each device has one or more outputs; data always flows forward along these edges.

For Part 2, we need to:

- Count all paths from a **start device** `svr` to the **output device** `out`.
- Only keep paths that visit **both** special devices `dac` and `fft` (in any order).
- The input can be very large, so a naive "enumerate all paths" approach would be far too slow.

The final answer is the number of valid `svr → … → out` paths that include both `dac` and `fft`.

---

## Graph model

Each input line has the form:

```text
device: childA childB childC
```

This is parsed into:

- A `Device` abstraction holding a name and a list of output names.
- A `Graph` made of `Node` objects connected by `next` pointers.

Important details of the model:

- Edges are **directed**: data only flows from a device to its outputs.
- The graph may contain **convergence** (multiple parents share a child) and significant **fan-out** (a node can have many children).
- We don’t store explicit paths; we only store structure plus some metadata used to prune the search.

`Graph` keeps:

- `nodesByName`: fast lookup by device name.
- `nodes`: a flat list of all created nodes (not used for the core traversal, but useful as a root set).

---

## Building the graph efficiently

The graph is built once from the list of devices in `BuildGraph`:

1. **Node reuse by name**  
   For each device name (and each of its outputs), we check `nodesByName`:

   - If the node already exists, we reuse it.
   - Otherwise we create it once and register it in `nodesByName`.

   This guarantees there is exactly one `Node` for each device name, even if it appears in many places. It also avoids repeatedly allocating and wiring subgraphs.

2. **Tracking required nodes on creation**  
   When a node is created, we check if its name is one of the required devices (`dac`, `fft` for Part 2). If so, we initialize its internal `requiredNodes` list with that single name. Other nodes start with an empty list.

3. **Propagating required-node information backwards**  
   When we connect `parent → child` via `AddNext`, the child’s `parents` slice is updated, and then:

   - If the child already knows that some required nodes are reachable below it, we propagate that information **backwards to the parent** (and recursively to its parents).
   - This "bubble up" of `requiredNodes` means: for any node, `node.requiredNodes` represents the set of required devices that can be reached somewhere in the subgraph starting from that node.

This propagation step is crucial: it gives every node an upper bound on how many required devices are still available if we go down from that node. The traversal uses this information to prune impossible paths early.

---

## Counting constrained paths

The method `Graph.CountPaths(from, to, requiredNodes)` delegates to `Node.CountPathsTo`. The core logic is in `Node.testPaths`, which recursively explores outgoing edges while enforcing the Part 2 constraints.

### Contract

`testPaths` is called as:

```go
func (n *Node) testPaths(
    to string,
    currentCount uint,
    requiredNodes []string,   // e.g. ["dac", "fft"]
    encounteredNodes []string,
    visitedNodes []string,
) uint
```

It returns the updated `currentCount`, i.e. the total number of valid paths found so far that end in `to` and that have visited all required nodes.

---

## Traversal optimizations

The naive DFS would:

- Explore every possible path from `svr` to `out`.
- Keep full path history.
- Recompute subtrees repeatedly.

That would be far too slow and memory-heavy for the actual input. Instead, the implementation adds several layers of optimization.

### 1. Required-nodes availability pruning

Each `Node` holds a `requiredNodes` slice, populated during graph building. For a given global list of required nodes (e.g. `{"dac", "fft"}`) and a runtime list of already `encounteredNodes`, we can compute how many required nodes are still missing on the current partial path.

Before traversing to a child `nextNode`, we check:

```go
if n.requiredNodes == nil ||
   len(nextNode.requiredNodes) < (len(requiredNodes) - len(encounteredNodes)) {
    // not enough required nodes reachable from this child; skip it
    continue
}
```

Interpretation:

- `requiredNodes` is the full set we need to see at least once on the path.
- `encounteredNodes` is the subset already seen on the current partial path.
- So `len(requiredNodes) - len(encounteredNodes)` is the count of required devices still missing.
- `nextNode.requiredNodes` is the set of required devices **reachable below** that child.

If the child’s reachable set is **smaller** than the number of missing required nodes, then no path going through this child can possibly hit all remaining required devices. We can safely skip that edge altogether.

This pruning is very powerful when the graph has a lot of branches that never contain `dac` or `fft`, or that only contain one of them when we still need both.

### 2. Path memoization (caching subtree results)

Each `Node` has a field `newPathsCount` that caches how many **additional valid paths** are contributed by the subtree rooted at that node for the given traversal configuration.

At the top of `testPaths`:

```go
if n.newPathsCount != -1 {
    // We have already visited this node; reuse the cached number of new paths.
    return currentCount + uint(n.newPathsCount)
}
```

At the end of `testPaths`, after exploring all children, we compute:

```go
n.newPathsCount = int64(currentCount) - int64(previousCurrentCount)
return currentCount
```

Here `previousCurrentCount` is the count before exploring this node’s children. So `newPathsCount` is the number of *extra* valid paths that went through this node.

Whenever we revisit the same node again along a different path prefix, we don’t re-traverse its entire subtree; we just add `newPathsCount` to the running total.

This turns what could be exponential behavior into something much closer to linear in the size of the graph (plus some overhead for the pruning logic).

> In particular, this avoids "traversing the same path again": once a node has been fully explored for the given set of required nodes, any future visits can immediately reuse the cached count instead of recursing again.

### 3. Tracking encountered required nodes only

Instead of storing and comparing entire paths, the algorithm only tracks:

- `visitedNodes`: mainly for debugging/printing during development.
- `encounteredNodes`: the subset of required devices we’ve seen so far on the current recursion branch.

When we reach the `to` device (`out`):

```go
if n.name == to {
    if len(requiredNodes) == len(encounteredNodes) {
        return currentCount + 1
    }
    return currentCount
}
```

So a path only counts if, by the time we hit `out`, we’ve seen `dac` and `fft` (order doesn’t matter). Non-qualifying paths are discarded immediately.

---

## Why this scales to the real input

Combining the above optimizations:

- **Graph build:** reuses nodes, tracks which required devices are reachable from each node, and pushes that information backwards so every node knows what’s available downstream.
- **Traversal:** uses that metadata to prune children that cannot possibly lead to a valid path, and memoizes the number of valid paths contributed by each node’s subtree.

This means:

- Large portions of the graph that never contain `dac` or `fft` are never explored.
- Shared subgraphs are explored **once**, even if they’re reachable from many different parents.
- The algorithm can handle very large path counts (like `5,024,474,986,900,123`) without ever constructing those paths explicitly.

---

## Running the solution

From the `day_11` directory:

```bash
make test   # run unit tests for the graph and node logic
make run    # run the actual puzzle solution
```
