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

	return !p.checkHasPattern()
}

func (p Product) checkHasPattern() bool {

	/* Starts the longest pattern first and then decreases down to a minimum length 2 (since 1 should already have been checked) */
	for _, patternLength := range p.getDivisors(len(p.Id)) {
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
	return p.Id[0:length]
}

func (p Product) hasPattern(
	pattern string,
) bool {
	patternLength := len(pattern)
	expectedPatternCount := len(p.Id) / patternLength
	/* Checks if the pattern is repeated within the product ID */
	for i := 0; i < expectedPatternCount; i++ {
		if p.Id[i*patternLength:i*patternLength+patternLength] != pattern {
			return false
		}
	}

	fmt.Printf("\t%q | Pattern %q found in product ID\n", p.Id, pattern)

	return true
}

func (p Product) getDivisors(
	length int,
) []int {

	var divisors = map[int][]int{
		1:  {1},
		2:  {1},
		3:  {1},
		4:  {1, 2},
		5:  {1},
		6:  {1, 2, 3},
		7:  {1},
		8:  {1, 2, 4},
		9:  {1, 3},
		10: {1, 2, 5},
		11: {1},
		12: {1, 2, 3, 4, 6},
		13: {1},
		14: {1, 2, 7},
		15: {1, 3, 5},
		16: {1, 2, 4, 8},
		17: {1},
		18: {1, 2, 3, 6, 9},
		19: {1},
		20: {1, 2, 4, 5, 10},
		21: {1, 3, 7},
		22: {1, 2, 11},
		23: {1},
		24: {1, 2, 3, 4, 6, 8, 12},
		25: {1, 5},
		26: {1, 2, 13},
		27: {1, 3, 9},
		28: {1, 2, 4, 7, 14},
		29: {1},
		30: {1, 2, 3, 5, 6, 10, 15},
	}

	return divisors[length]
}
