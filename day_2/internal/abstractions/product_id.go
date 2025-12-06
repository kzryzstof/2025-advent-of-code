package abstractions

type ProductId string

func (p ProductId) IsValid() bool {

	if len(p)%2 != 0 {
		return true
	}

	return false
}
