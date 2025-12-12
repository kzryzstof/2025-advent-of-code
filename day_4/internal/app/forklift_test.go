package app

import (
	"day_4/internal/abstractions"
	"testing"
)

// Helper to build a Department from a slice of ASCII rows where
// '.' maps to abstractions.Empty and '@' maps to abstractions.Roll.
func buildDepartment(lines []string) *abstractions.Department {
	rows := make([]abstractions.Row, len(lines))
	for i, line := range lines {
		spots := make([]abstractions.Spot, len(line))
		for j, ch := range line {
			if ch == '.' {
				spots[j] = abstractions.Empty
			} else if ch == '@' {
				spots[j] = abstractions.Roll
			}
		}
		rows[i] = abstractions.Row{
			Number: uint(i + 1),
			Spots:  spots,
		}
	}
	return &abstractions.Department{Rows: rows}
}

func TestForklift_RemoveRolls(t *testing.T) {
	// Documented use case with full pattern:
	// ..@@.@@@@.
	// @@@.@.@.@@
	// @@@@@.@.@@
	// @.@@@@..@.
	// @@.@@@@.@@
	// .@@@@@@@.@
	// .@.@.@.@@@
	// @.@@@.@@@@
	// .@@@@@@@@.
	// @.@.@@@.@.
	lines := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	department := buildDepartment(lines)

	forklift := NewForklift(false)
	forklift.RemoveRolls(department)

	// This expected value should match the total number of accessible rolls
	// removed when applying the forklift rules to the full pattern. If the
	// rules change, update this constant accordingly.
	const expectedAccessibleRolls uint = 43

	actualAccessibleRolls := forklift.GetAccessedRollsCount()
	if actualAccessibleRolls != expectedAccessibleRolls {
		t.Fatalf("expected %d accessible rolls, got %d", expectedAccessibleRolls, actualAccessibleRolls)
	}
}
