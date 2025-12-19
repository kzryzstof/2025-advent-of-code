package abstractions

import "fmt"

// VariableNumber /* Defines a type for variable number used in equations. 1st-based indexed */
type VariableNumber uint8

type Variable struct {
	Number VariableNumber
	Value  float64
}

type Variables struct {
	variables []*Variable
}

func NewVariables(
	count uint,
) *Variables {
	variables := make([]*Variable, count)

	for i := uint(0); i < count; i++ {
		variables[i] = &Variable{
			Number: VariableNumber(i + 1),
			Value:  0.0,
		}
	}

	return &Variables{variables}
}

func FromVariableNumbers(
	variableNumbers []VariableNumber,
) *Variables {

	v := make([]*Variable, len(variableNumbers))

	for i := 0; i < len(variableNumbers); i++ {
		v[i] = &Variable{
			Number: variableNumbers[i],
			Value:  0.0,
		}
	}

	return &Variables{v}
}

func (v *Variables) Get() []*Variable {
	return v.variables
}

func (v *Variables) SetVariable(variable *Variable) {
	v.Set(variable.Number, variable.Value)
}

func (v *Variables) Count() uint {
	return uint(len(v.variables))
}

func (v *Variables) IsLast(
	number VariableNumber,
) bool {
	return number == VariableNumber(len(v.variables))
}

func (v *Variables) Set(
	number VariableNumber,
	value float64,
) {
	v.variables[number-1].Value = value
}

func (v *Variables) Contains(
	number int,
) bool {
	for _, variable := range v.variables {
		if int(variable.Number) == number {
			return true
		}
	}

	return false
}

func (v *Variables) Print() {
	fmt.Print("[")
	for _, variable := range v.variables {
		fmt.Printf(" %d: %.2f ", variable.Number, variable.Value)
	}
	fmt.Print(" ]\n")
}
