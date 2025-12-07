# Day 3 – Battery Banks (Part 1)

The Day 3 Part 1 solution is implemented in Go in the `day_3` folder. It reads **battery banks** from an input file, finds the two highest‑voltage batteries in each bank, and sums a combined "voltage" value for all banks.

## Problem model (Part 1)

- The input file contains one **bank** per line.
- Each character on a line is a single battery's **voltage rating**, a digit `0`–`9`.
  - For example, the line `3845` represents a bank with four batteries whose voltages are `[3, 8, 4, 5]`.
- For each bank, we:
  1. Find the **highest‑voltage battery** in the bank, but **exclude the last battery** from consideration for this first pick.
  2. Then, starting **after** that first chosen battery, find the **next highest‑voltage battery** among the remaining batteries (this time including the last battery).
  3. Combine these two voltages into a two‑digit number: `first * 10 + second`.
- The answer for Part 1 is the **sum** of this two‑digit value for every bank in the input.

So, conceptually, for each line of digits we pick two batteries in order:

- The first is the maximum voltage in the prefix that excludes the last battery.
- The second is the maximum voltage in the suffix that starts immediately after the first battery and goes to the end.

## High-level flow

The executable entrypoint is `day_3/cmd/main.go` and the flow is:

1. **Read the input file path** from the command‑line arguments.
2. **Initialize a banks parser** to stream `Bank` values from the file.
3. **Initialize a banks processor** that consumes those banks and accumulates a total voltage value.
4. Run the parser and processor concurrently using a `sync.WaitGroup`.
5. After all banks have been processed, print:
   - `Sum of all the highest voltage from the <banksCount> banks: <total>`

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
    - `GetHighestVoltage()` implements the core Part 1 logic:
      - Calls a helper `getHighestVoltage(fromIndex, toIndex)` to search sub‑slices of `Batteries`.
      - First call searches from index `0` up to (but not including) the last battery to find the first highest‑voltage battery and its index.
      - Second call searches from `firstIndex + 1` up to the end to find the second highest‑voltage battery.
      - Returns `firstVoltage*10 + secondVoltage` as an unsigned integer.
  - `BanksChannel` – small interface that exposes `Banks() <-chan Bank`, enabling the processor to depend only on the interface instead of the parser implementation.

## Running Day 3 Part 1

From the `day_3` directory you can run the solution with:

```bash
cd day_3
make run ARGS="input.txt"
```

or directly with Go:

```bash
go run ./cmd input.txt
```

