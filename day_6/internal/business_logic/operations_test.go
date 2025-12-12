package business_logic

import "testing"

func TestCompute(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		numbers   []uint64
		want      uint64
		wantErr   bool
	}{
		{
			name:      "single number with plus",
			operation: "+",
			numbers:   []uint64{5},
			want:      5,
			wantErr:   false,
		},
		{
			name:      "single number with multiply",
			operation: "*",
			numbers:   []uint64{7},
			want:      7,
			wantErr:   false,
		},
		{
			name:      "simple addition",
			operation: "+",
			numbers:   []uint64{1, 2, 3, 4},
			want:      10,
			wantErr:   false,
		},
		{
			name:      "simple multiplication",
			operation: "*",
			numbers:   []uint64{2, 3, 4},
			want:      24,
			wantErr:   false,
		},
		{
			name:      "includes zero in multiplication",
			operation: "*",
			numbers:   []uint64{5, 0, 10},
			want:      0,
			wantErr:   false,
		},
		{
			name:      "includes zero in addition",
			operation: "+",
			numbers:   []uint64{5, 0, 10},
			want:      15,
			wantErr:   false,
		},
		{
			name:      "invalid operation",
			operation: "-",
			numbers:   []uint64{1, 2, 3},
			want:      0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compute(tt.operation, tt.numbers)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil, result=%d", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("Compute(%q, %v) = %d, want %d", tt.operation, tt.numbers, got, tt.want)
			}
		})
	}
}
