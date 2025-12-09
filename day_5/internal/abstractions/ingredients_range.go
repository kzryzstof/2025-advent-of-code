package abstractions

type IngredientId uint

type IngredientRange struct {
	From IngredientId
	To   IngredientId
}

type AvailableIngredients struct {
	Ids []IngredientId
}

type FreshIngredients struct {
	Ranges []IngredientRange
}
