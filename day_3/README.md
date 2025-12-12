# Day 3 – Battery Banks

The Day 3 solution is implemented in Go in the `day_3` folder. It reads **battery banks** from an input file and, for each bank, computes a 12‑digit "combined voltage" value that is then summed across all banks.

---

## Problem model

- The input file contains one **bank** per line.
- Each character on a line is a single battery's **voltage rating**, a digit `0`–`9`.
  - For example, the line `3845` represents a bank with four batteries whose voltages are `[3, 8, 4, 5]`.

For each bank, the program selects 12 batteries (digits) in a greedy way to form the largest possible 12‑digit number while preserving the left‑to‑right order of batteries.

The final answer is the **sum** of these 12‑digit values for all banks.

---

## Combined voltage – 12‑digit greedy algorithm

For each bank, we construct a **12‑digit combined voltage** by greedily choosing batteries from left to right:

- At each step we choose the next digit to be as large as possible.
- We must always leave **enough remaining batteries** to be able to complete all 12 digits.
- Once a battery is chosen for a digit, all batteries **up to and including** that position are no longer available for subsequent digits.

For a bank with `N` batteries (digits) and `expectedDigits = 12`:

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

In code, this logic is implemented in `Bank.GetHighestVoltage()` (in `internal/abstractions/bank.go`), which uses a helper `getHighestVoltage(fromIndex, toIndex)` to search sub‑slices of `Batteries` and build the final multi‑digit value.

---

## High-level flow

The executable entrypoint is `day_3/cmd/main.go` and the flow is:

1. **Read the input file path** from the command‑line arguments.
2. **Initialize a banks reader** (`BanksReader`) with that file path.
3. Call `reader.Read()` to synchronously read **all banks** from the file into a slice.
4. For each bank in that slice, call `bank.GetHighestVoltage()` to compute its combined voltage.
5. Accumulate these values into `totalVoltage` and finally print:
   - `Sum of all the highest voltage from the <banksCount> banks: <total>`

---

## Packages and responsibilities

- `cmd/main.go`
  - Wires everything together.
  - Reads the input path from `os.Args`.
  - Creates the `BanksReader` via `io.NewReader`.
  - Calls `reader.Read()` to obtain a `[]Bank`.
  - Iterates over all banks, calls `bank.GetHighestVoltage()` on each, sums the results, and prints the final total.

- `internal/io`
  - `BanksReader` encapsulates reading banks from the input file.
  - `NewReader(filePath string)` opens the file and returns a reader.
  - `Read() ([]abstractions.Bank, error)`:
    - Uses a `bufio.Scanner` to read the file line by line.
    - For each non‑empty line:
      - Iterates over each rune (digit) in the line.
      - Converts it to an integer using `strconv.Atoi`.
      - Wraps each digit as a `Battery` with a `Voltage` of type `VoltageRating`.
      - Groups all batteries on the line into a `Bank{Batteries: [...]}`.
    - Appends each bank to a slice and returns the full slice at the end.

- `internal/abstractions`
  - `Battery` – represents a single battery in a bank.
    - `Voltage` is a `VoltageRating` (an alias of `uint`).
  - `Bank` – a collection of batteries for one line of input.
    - `GetHighestVoltage()` implements the 12‑digit greedy selection logic described above, using:
      - A loop over `expectedDigits` (12).
      - A moving `currentBatteryIndex` to ensure order and availability constraints.
      - The helper `getHighestVoltage(fromIndex, toIndex)` to scan a window and return the maximum voltage and its absolute index.

---

## Running Day 3

From the `day_3` directory you can run the solution with:

```bash
cd day_3
make run ARGS="input.txt"
```

or directly with Go:

```bash
cd day_3
go run ./cmd input.txt
```
