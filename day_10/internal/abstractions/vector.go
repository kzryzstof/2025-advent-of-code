package abstractions

type Vector struct {
	values []float64
	length int
}

func NewVector(
	length int,
) *Vector {

	values := make([]float64, length)

	return &Vector{
		values,
		length,
	}
}

func (v *Vector) Get(x int) float64 {
	return v.values[x]
}

func (v *Vector) Set(x int, value float64) {
	v.values[x] = value
}

func (v *Vector) Swap(fromIndex, toIndex int) {
	value := v.values[fromIndex]
	v.values[fromIndex] = v.values[toIndex]
	v.values[toIndex] = value
}

func (v *Vector) Scale(index int, factor float64) {
	v.Set(index, v.Get(index)*factor)
}
