package abstractions

import "testing"

func TestTranspose3x3(t *testing.T) {
	// Only 3x3 matrices, table-driven.
	tests := []struct {
		name string
		in   [][]int8
		want [][]int8
	}{
		{
			name: "row-major identity -> transposed",
			in: [][]int8{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			want: [][]int8{
				{'A', 'D', 'G'},
				{'B', 'E', 'H'},
				{'C', 'F', 'I'},
			},
		},
		{
			name: "transposed -> back to row-major",
			in: [][]int8{
				{'A', 'D', 'G'},
				{'B', 'E', 'H'},
				{'C', 'F', 'I'},
			},
			want: [][]int8{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
		},
	}

	for _, tt := range tests {
		// copy input so test cases don't share backing arrays
		got := make([][]int8, len(tt.in))
		for i := range tt.in {
			got[i] = append([]int8(nil), tt.in[i]...)
		}

		Transpose(got)

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if got[r][c] != tt.want[r][c] {
					t.Fatalf("%s: mismatch at (%d,%d): got %q, want %q", tt.name, r, c, got[r][c], tt.want[r][c])
				}
			}
		}
	}
}

func TestHorizontalFlip3x3(t *testing.T) {
	// renamed from Reverse; horizontal flip of each row.
	tests := []struct {
		name string
		in   [][]int8
		want [][]int8
	}{
		{
			name: "horizontal flip of each row",
			in: [][]int8{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			want: [][]int8{
				{'C', 'B', 'A'},
				{'F', 'E', 'D'},
				{'I', 'H', 'G'},
			},
		},
	}

	for _, tt := range tests {
		got := make([][]int8, len(tt.in))
		for i := range tt.in {
			got[i] = append([]int8(nil), tt.in[i]...)
		}

		HorizontalFlip(got)

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if got[r][c] != tt.want[r][c] {
					t.Fatalf("%s: mismatch at (%d,%d): got %q, want %q", tt.name, r, c, got[r][c], tt.want[r][c])
				}
			}
		}
	}
}

func TestVerticalFlip3x3(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int8
		want [][]int8
	}{
		{
			name: "vertical flip of rows",
			in: [][]int8{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			want: [][]int8{
				{'G', 'H', 'I'},
				{'D', 'E', 'F'},
				{'A', 'B', 'C'},
			},
		},
	}

	for _, tt := range tests {
		got := make([][]int8, len(tt.in))
		for i := range tt.in {
			got[i] = append([]int8(nil), tt.in[i]...)
		}

		VerticalFlip(got)

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if got[r][c] != tt.want[r][c] {
					t.Fatalf("%s: mismatch at (%d,%d): got %q, want %q", tt.name, r, c, got[r][c], tt.want[r][c])
				}
			}
		}
	}
}

func TestRotateClockwise3x3(t *testing.T) {
	// Only 3x3 matrices, table-driven.
	tests := []struct {
		name string
		in   [][]int8
		want [][]int8
	}{
		{
			name: "rotate 90 degrees clockwise",
			in: [][]int8{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			want: [][]int8{
				{'G', 'D', 'A'},
				{'H', 'E', 'B'},
				{'I', 'F', 'C'},
			},
		},
	}

	for _, tt := range tests {
		got := make([][]int8, len(tt.in))
		for i := range tt.in {
			got[i] = append([]int8(nil), tt.in[i]...)
		}

		RotateClockwise(got)

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if got[r][c] != tt.want[r][c] {
					t.Fatalf("%s: mismatch at (%d,%d): got %q, want %q", tt.name, r, c, got[r][c], tt.want[r][c])
				}
			}
		}
	}
}
