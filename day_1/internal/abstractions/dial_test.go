package abstractions

import "testing"

func TestDial_Turn(t *testing.T) {
	tests := []struct {
		name             string
		initialPosition  int
		direction        Direction
		distance         int
		expectedPosition int
		expectedCount    int
	}{
		{
			name:             "Turn right without wrap",
			initialPosition:  10,
			direction:        Right,
			distance:         5,
			expectedPosition: 15,
			expectedCount:    0,
		},
		{
			name:             "Turn right with wrap",
			initialPosition:  95,
			direction:        Right,
			distance:         10,
			expectedPosition: 5,
			expectedCount:    1,
		},
		{
			name:             "Full turn right but not at position 0",
			initialPosition:  10,
			direction:        Right,
			distance:         100,
			expectedPosition: 10,
			expectedCount:    1,
		},
		{
			name:             "Full turn right and at position 0",
			initialPosition:  0,
			direction:        Right,
			distance:         100,
			expectedPosition: 0,
			expectedCount:    1,
		},
		{
			name:             "2 full turns right and at position 0",
			initialPosition:  0,
			direction:        Right,
			distance:         200,
			expectedPosition: 0,
			expectedCount:    2,
		},
		{
			name:             "Turn right with 2 wraps",
			initialPosition:  95,
			direction:        Right,
			distance:         100 + 10,
			expectedPosition: 5,
			expectedCount:    2,
		},
		{
			name:             "Turn left without wrap",
			initialPosition:  20,
			direction:        Left,
			distance:         5,
			expectedPosition: 15,
			expectedCount:    0,
		},
		{
			name:             "Turn left with wrap",
			initialPosition:  5,
			direction:        Left,
			distance:         10,
			expectedPosition: 95,
			expectedCount:    1,
		},
		{
			name:             "Turn left completely wrapped",
			initialPosition:  5,
			direction:        Left,
			distance:         100,
			expectedPosition: 5,
			expectedCount:    1,
		},
		{
			name:             "Turn left with 2 wrap",
			initialPosition:  5,
			direction:        Left,
			distance:         100 + 10,
			expectedPosition: 95,
			expectedCount:    2,
		},
		/* original use cases (part 1) */
		{
			name:             "Documented use case #1",
			initialPosition:  11,
			direction:        Right,
			distance:         8,
			expectedPosition: 19,
			expectedCount:    0,
		},
		{
			name:             "Documented use case #2",
			initialPosition:  19,
			direction:        Left,
			distance:         19,
			expectedPosition: 0,
			expectedCount:    1,
		},
		{
			name:             "Documented use case #3",
			initialPosition:  0,
			direction:        Left,
			distance:         1,
			expectedPosition: 99,
			expectedCount:    0,
		},
		{
			name:             "Documented use case #4",
			initialPosition:  99,
			direction:        Right,
			distance:         1,
			expectedPosition: 0,
			expectedCount:    1,
		},
		/* new use cases (part 2) */
		{
			name:             "Documented use case #1",
			initialPosition:  50,
			direction:        Left,
			distance:         68,
			expectedPosition: 82,
			expectedCount:    1,
		},
		{
			name:             "Documented use case #2",
			initialPosition:  52,
			direction:        Right,
			distance:         48,
			expectedPosition: 0,
			expectedCount:    1,
		},
		{
			name:             "Documented use case #3",
			initialPosition:  0,
			direction:        Left,
			distance:         5,
			expectedPosition: 95,
			expectedCount:    0,
		},
		{
			name:             "Documented use case #3",
			initialPosition:  50,
			direction:        Right,
			distance:         1000,
			expectedPosition: 50,
			expectedCount:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dial := Dial{Position: tt.initialPosition}
			dial.Rotate(Rotation{Direction: tt.direction, Distance: tt.distance})

			if dial.Position != tt.expectedPosition {
				t.Errorf("Expected position %d, got %d", tt.expectedPosition, dial.Position)
			}

			if dial.GetCount() != tt.expectedCount {
				t.Errorf("Expected count %d, got %d", tt.expectedCount, dial.GetCount())
			}
		})
	}
}
