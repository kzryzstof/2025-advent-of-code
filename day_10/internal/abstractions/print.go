package abstractions

import "fmt"

func Print(
	m *Matrix,
	v *Vector,
) {
	for row := 0; row < m.Rows(); row++ {
		fmt.Printf("[ ")
		for col := 0; col < m.Cols(); col++ {
			fmt.Printf("%.2f ", m.Get(row, col))
		}
		fmt.Printf("] ")

		fmt.Printf("[")
		fmt.Printf(" %.2f ", v.Get(row))
		fmt.Println("] ")
	}
	fmt.Println()
}
