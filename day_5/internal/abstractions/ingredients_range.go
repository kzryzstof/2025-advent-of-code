package abstractions

type IngredientId uint64

type IngredientRange struct {
	From IngredientId
	To   IngredientId
}

func (r IngredientRange) IsIncluded(id IngredientId) bool {
	return id >= r.From && id <= r.To
}

type AvailableIngredients struct {
	Ids []IngredientId
}

type FreshIngredients struct {
	Ranges []IngredientRange
}

func (f FreshIngredients) IsFresh(id IngredientId) bool {
	for freshRange := range f.Ranges {
		if f.Ranges[freshRange].IsIncluded(id) {
			return true
		}
	}
	return false
}
