package abstractions

import (
	"fmt"
	"math"
)

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
		variables[i] = nil
	}

	return &Variables{variables}
}

func CopyVariables(
	variables *Variables,
) *Variables {

	variablesCopy := make([]*Variable, len(variables.variables))

	for i := 0; i < len(variablesCopy); i++ {
		variablesCopy[i] = CopyVariable(variables.variables[i])
	}

	return &Variables{variablesCopy}
}

func CopyVariable(
	variable *Variable,
) *Variable {

	if variable == nil {
		return nil
	}

	return &Variable{
		variable.Number,
		math.Floor(variable.Value),
	}
}

func FromVariableNumbers(
	variableNumbers []VariableNumber,
	defaultValue float64,
) *Variables {

	v := make([]*Variable, len(variableNumbers))

	for i := 0; i < len(variableNumbers); i++ {
		v[i] = &Variable{
			Number: variableNumbers[i],
			Value:  defaultValue,
		}
	}

	return &Variables{v}
}

func (v *Variables) Get() []*Variable {
	return v.variables
}

func (v *Variables) GetNumberByIndex(
	index uint,
) VariableNumber {
	return v.variables[index].Number
}

func (v *Variables) SetVariable(variable *Variable) {
	v.variables[variable.Number-1] = variable
}

func (v *Variables) Count() uint {
	count := uint(0)
	for _, variable := range v.variables {
		if variable != nil {
			count++
		}
	}
	return count
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
	for _, variable := range v.variables {
		if variable.Number == number {
			variable.Value = value
			break
		}
	}
}

func (v *Variables) Contains(
	number VariableNumber,
) bool {
	for _, variable := range v.variables {
		if variable == nil {
			continue
		}
		if variable.Number == number {
			return true
		}
	}

	return false
}

func (v *Variables) GetValue(
	number VariableNumber,
) float64 {
	for _, variable := range v.variables {
		if variable == nil {
			continue
		}
		if variable.Number == number {
			return variable.Value
		}
	}

	panic(fmt.Errorf("no variable %d found", number))
}

func (v *Variables) GetValues() []float64 {

	// BUG WITH NIL
	values := make([]float64, v.Count())

	for i, variable := range v.variables {
		values[i] = variable.Value
	}

	return values
}

func ContainsNumber(
	variableNumbers []VariableNumber,
	number int,
) bool {
	for _, variable := range variableNumbers {
		if int(variable) == number {
			return true
		}
	}

	return false
}

func (v *Variables) Print() {
	fmt.Print("[")
	for _, variable := range v.variables {
		if variable == nil {
			fmt.Printf(" XX: XXX ")
			continue
		}
		fmt.Printf(" %d: %.2f ", variable.Number, variable.Value)
	}
	fmt.Print(" ]\n")
}
