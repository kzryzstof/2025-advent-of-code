package abstractions

type Variables struct {
	variables []*Variable
}

func NewVariables(
	count uint64,
) *Variables {
	variables := make([]*Variable, count)

	for i := uint64(0); i < count; i++ {
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
