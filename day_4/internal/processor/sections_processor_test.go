package processor

import (
	"day_4/internal/abstractions"
	"testing"
)

func TestSectionsProcessor(t *testing.T) {
	tests := []struct {
		name                    string
		rows                    []abstractions.Row
		rowIndex                int
		expectedAccessibleRolls uint
	}{
		/* Documented use case with full pattern:
		   ..@@.@@@@.
		   @@@.@.@.@@
		   @@@@@.@.@@
		   @.@@@@..@.
		   @@.@@@@.@@
		   .@@@@@@@.@
		   .@.@.@.@@@
		   @.@@@.@@@@
		   .@@@@@@@@.
		   @.@.@@@.@.
		*/
		{
			name: "Row 0",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
			},
			rowIndex:                0,
			expectedAccessibleRolls: 5,
		},
		{
			name: "Row 1",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
			},
			rowIndex:                0,
			expectedAccessibleRolls: 5,
		},
		{
			name: "Row 2",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 1,
		},
		{
			name: "Row 3",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Empty, abstractions.Roll, abstractions.Empty}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 1,
		},
		{
			name: "Row 4",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Empty, abstractions.Roll, abstractions.Empty}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 0,
		},
		{
			name: "Row 5",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Empty, abstractions.Roll, abstractions.Empty}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 2,
		},
		{
			name: "Row 6",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 0,
		},
		{
			name: "Row 7",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 0,
		},
		{
			name: "Row 8",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 1,
		},
		{
			name: "Row 9",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty}},
			},
			rowIndex:                1,
			expectedAccessibleRolls: 0,
		},
		{
			name: "Row 10",
			rows: []abstractions.Row{
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll}},
				{Spots: []abstractions.Spot{abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty}},
				{Spots: []abstractions.Spot{abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Roll, abstractions.Roll, abstractions.Empty, abstractions.Roll, abstractions.Empty}},
			},
			rowIndex:                2,
			expectedAccessibleRolls: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			section := abstractions.Section{
				Rows:     tt.rows,
				RowIndex: tt.rowIndex,
			}

			processor := NewProcessor()
			processor.Analyze(&section)

			actualAccessibleRolls := processor.GetTotalAccessibleRolls()
			if actualAccessibleRolls != tt.expectedAccessibleRolls {
				t.Fatalf("expected %d accessible rolls, got %d", tt.expectedAccessibleRolls, actualAccessibleRolls)
			}
		})
	}
}
