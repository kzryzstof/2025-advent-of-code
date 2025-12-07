package abstractions

type Range struct {
	From Product
	To   Product
}

func (r Range) FindInvalidProductIds() []int64 {
	invalidProductIds := make([]int64, 0)

	for product := &r.From; !product.IsLast(&r.To); {

		if !product.IsValid() {
			invalidProductIds = append(invalidProductIds, product.GetNumber())
		}

		product = product.Next()
	}

	return invalidProductIds
}
