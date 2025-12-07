package abstractions

import (
	"fmt"
	"strconv"
)

var divisors = map[int][]int{
	1:  {1},
	2:  {1},
	3:  {1},
	4:  {2, 1},
	5:  {1},
	6:  {3, 2, 1},
	7:  {1},
	8:  {4, 2, 1},
	9:  {3, 1},
	10: {5, 2, 1},
	11: {1},
	12: {6, 4, 3, 2, 1},
	13: {1},
	14: {7, 2, 1},
	15: {5, 3, 1},
	16: {8, 4, 2, 1},
	17: {1},
	18: {9, 6, 3, 2, 1},
	19: {1},
	20: {10, 5, 4, 2, 1},
	21: {7, 3, 1},
	22: {11, 2, 1},
	23: {1},
	24: {12, 8, 6, 4, 3, 2, 1},
	25: {5, 1},
	26: {13, 2, 1},
	27: {9, 3, 1},
	28: {14, 7, 4, 2, 1},
	29: {1},
	30: {15, 10, 6, 5, 3, 2, 1},
}

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

	return !p.checkHasPattern()
}

func (p Product) checkHasPattern() bool {

	idLength := len(p.Id)

	if idLength == 1 {
		return false
	}

	for _, patternLength := range p.getDivisors(idLength) {
		if patternLength <= 0 || idLength%patternLength != 0 {
			continue
		}
		pattern := p.getPattern(patternLength)
		if p.hasPattern(pattern) {
			return true
		}
	}

	return false
}

func (p Product) getPattern(
	length int,
) string {
	return p.Id[:length]
}

func (p Product) hasPattern(
	pattern string,
) bool {
	patternLength := len(pattern)
	expectedPatternCount := len(p.Id) / patternLength

	/* Checks if the pattern is repeated within the product ID */
	for i := 0; i < expectedPatternCount*patternLength; i += patternLength {
		if p.Id[i:i+patternLength] != pattern {
			return false
		}
	}

	fmt.Printf("\t%q | Pattern %q found in product ID\n", p.Id, pattern)

	return true
}

func (p Product) getDivisors(
	length int,
) []int {
	return divisors[length]
}
