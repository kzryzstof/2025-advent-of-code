package abstractions

import (
	"fmt"
	"strconv"
)

type Product struct {
	Id         string
	internalId int64
}

func NewProduct(id string) (*Product, error) {
	internalId, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return &Product{
		id,
		int64(internalId),
	}, nil
}

func (p Product) GetNumber() int64 {
	return p.internalId
}

func (p Product) IsLast(lastProduct *Product) bool {
	return p.internalId > lastProduct.internalId
}

func (p Product) Next() *Product {
	nextInternalId := p.internalId + 1
	return &Product{fmt.Sprint(nextInternalId), nextInternalId}
}

func (p Product) IsValid() bool {

	if len(p.Id)%2 != 0 {
		return true
	}

	middleSidePosition := len(p.Id) / 2

	for i := 0; i < middleSidePosition; i++ {
		if p.Id[i] != p.Id[middleSidePosition+i] {
			return true
		}
	}

	return false
}
