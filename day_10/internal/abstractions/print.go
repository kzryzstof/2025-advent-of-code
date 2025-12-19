package abstractions

import "fmt"

func Print(
	m *Matrix,
) {
	for row := 0; row < m.Rows(); row++ {
		fmt.Printf("[ ")
		for col := 0; col < m.Cols(); col++ {
			fmt.Printf("%.2f ", m.Get(row, col))
		}
		fmt.Printf("] ")
	}
	fmt.Println()
}

func PrintSlice(
	vector []float64,
) {

	fmt.Printf("[")

	for row := 0; row < len(vector); row++ {
		fmt.Printf(" %.2f ", vector[row])
	}
	fmt.Println("]")
}
