# Day 3 – Battery Banks (Part 1 & Part 2)

The Day 3 solutions are implemented in Go in the `day_3` folder. They read **battery banks** from an input file and, for each bank, compute a "combined voltage" value that is then summed across all banks.

- **Part 1** picks two batteries per bank and forms a two‑digit voltage.
- **Part 2** generalizes this to pick **12 batteries** per bank and form a 12‑digit voltage using a greedy algorithm.

---

## Problem model

- The input file contains one **bank** per line.
- Each character on a line is a single battery's **voltage rating**, a digit `0`–`9`.
  - For example, the line `3845` represents a bank with four batteries whose voltages are `[3, 8, 4, 5]`.

---

## Part 1 – Two‑digit combined voltage

For each bank, Part 1:

1. Finds the **highest‑voltage battery** in the bank, but **excludes the last battery** from consideration for this first pick.
2. Then, starting **after** that first chosen battery, finds the **next highest‑voltage battery** among the remaining batteries (this time including the last battery).
3. Combines these two voltages into a two‑digit number: `first * 10 + second`.

Conceptually, for each line of digits we pick two batteries in order:

- The first is the maximum voltage in the prefix that excludes the last battery.
- The second is the maximum voltage in the suffix that starts immediately after the first battery and goes to the end.

The answer for Part 1 is the **sum** of this two‑digit value for every bank in the input.

---

## Part 2 – 12‑digit greedy combined voltage

For Part 2, the rule is extended: instead of picking only two batteries, we now pick **exactly 12 batteries** from each bank, in order, to form a 12‑digit number.

- At each step we choose the next digit to be as large as possible.
- We must always leave **enough remaining batteries** to be able to complete all 12 digits.
- Once a battery is chosen for a digit, all batteries **up to and including** that position are no longer available for subsequent digits.

More formally, for a bank with `N` batteries (digits) and `expectedDigits = 12`:

1. Keep a pointer `currentBatteryIndex` that marks the first index we are allowed to choose from.
2. For each digit position `d` from `0` to `expectedDigits-1`:
   - Compute how many digits must remain after choosing this digit:
     - `minimumRemainingDigits = expectedDigits - d - 1`.
   - This means we can only choose from the range of indices:
     - `fromIndex = currentBatteryIndex`.
     - `toIndex = batteryCount - minimumRemainingDigits`.
     - In other words, we look for the maximum voltage in `Batteries[fromIndex:toIndex]`.
   - Find the **highest‑voltage battery** in that window, recording both its value and **absolute index**.
   - Append that value as the next digit of the growing 12‑digit number, i.e.:
     - `totalVoltage += digit * 10^(expectedDigits-d-1)`.
   - Move `currentBatteryIndex` to **one past** the chosen battery so that subsequent digits can only come from batteries to its right.

This procedure yields a 12‑digit integer that is the lexicographically largest possible number you can build by picking 12 digits in order from the original sequence while preserving relative order.

The overall Part 2 answer is the **sum** of these 12‑digit values for all banks.

In code, this logic is implemented in `Bank.GetHighestVoltage()` (in `internal/abstractions/bank.go`), which uses a helper `getHighestVoltage(fromIndex, toIndex)` to search sub‑slices of `Batteries` and build the final multi‑digit value.

---

## High-level flow (shared by Part 1 & Part 2)

The executable entrypoint is `day_3/cmd/main.go` and the flow is:

1. **Read the input file path** from the command‑line arguments.
2. **Initialize a banks parser** to stream `Bank` values from the file.
3. **Initialize a banks processor** that consumes those banks and accumulates a total voltage value.
4. Run the parser and processor concurrently using a `sync.WaitGroup`.
5. After all banks have been processed, print:
   - `Sum of all the highest voltage from the <banksCount> banks: <total>`

The only difference between Part 1 and Part 2 is the internal logic of `Bank.GetHighestVoltage()`; the wiring in `cmd`, the parser, and the processor remain the same.

---

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the `BanksParser` and `BanksProcessor`.
  - Starts both concurrently, waits for them to finish, and prints the final total.

- `internal/parser`
  - `BanksParser` opens the input file and scans it line by line.
  - For each non‑empty line, it:
    - Iterates over each rune (digit) in the line.
    - Converts it to an integer using `strconv.Atoi`.
    - Wraps each digit as a `Battery` with a `Voltage` of type `VoltageRating`.
    - Groups all batteries on the line into a `Bank{Batteries: [...]}`.
  - Sends each `Bank` on an internal channel, tracks how many banks were parsed, and finally closes the channel.
  - Exposes the channel via `Banks()` and the count via `GetBanksCount()`.

- `internal/processor`
  - `BanksProcessor` takes something that implements `BanksChannel` and a `*sync.WaitGroup`.
  - Runs a goroutine that ranges over the stream of `Bank` values.
  - For each bank, it calls `bank.GetHighestVoltage()` and adds the result to `totalVoltage`.
  - Exposes the final sum via `GetTotalVoltage()`.

- `internal/abstractions`
  - `Battery` – represents a single battery in a bank.
    - `Voltage` is a `VoltageRating` (an alias of `uint`).
  - `Bank` – a collection of batteries for one line of input.
    - `GetHighestVoltage()` implements the core Part 2 greedy selection logic described above.
  - `BanksChannel` – small interface that exposes `Banks() <-chan Bank`, enabling the processor to depend only on the interface instead of the parser implementation.

---

## Running Day 3

From the `day_3` directory you can run the solution with:

```bash
cd day_3
make run ARGS="input.txt"
```

or directly with Go:

```bash
go run ./cmd input.txt
```
