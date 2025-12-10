package abstractions

import "fmt"

type IngredientId uint64

type IngredientRange struct {
	From IngredientId
	To   IngredientId
}

func (r *IngredientRange) Replace(ingredientRange IngredientRange) {
	r.From = ingredientRange.From
	r.To = ingredientRange.To
}

func (r *IngredientRange) IsIncluded(id IngredientId) bool {
	return id >= r.From && id <= r.To
}

func (r *IngredientRange) Count() uint64 {
	return uint64(r.To - r.From + 1)
}

type FreshIngredients struct {
	Ranges []IngredientRange
}

func (f *FreshIngredients) IsFresh(id IngredientId) bool {
	for freshRange := range f.Ranges {
		if f.Ranges[freshRange].IsIncluded(id) {
			return true
		}
	}
	return false
}

func (f *FreshIngredients) Count() uint64 {
	freshRangesCount := uint64(0)

	for _, freshRange := range f.Ranges {
		freshRangesCount += freshRange.Count()
	}

	return freshRangesCount
}

func (f *FreshIngredients) Compact() *FreshIngredients {

	copyRange := &(f.Ranges)

	compactingDone := false

	for !compactingDone {

		rangesCount := len(*copyRange)

		newRanges := make([]IngredientRange, 0, rangesCount)
		excludedRanges := make([]IngredientRange, 0, rangesCount)

		for sourceIndex := 0; sourceIndex < rangesCount; sourceIndex++ {

			sourceRange := (*copyRange)[sourceIndex]

			if isFound, _ := exists(&excludedRanges, sourceRange); isFound == true {
				continue
			}

			for otherIndex := 0; otherIndex < rangesCount; otherIndex++ {

				if sourceIndex == otherIndex {
					continue
				}

				otherRange := (*copyRange)[otherIndex]

				if sourceRange.From <= otherRange.From && sourceRange.To >= otherRange.To {
					//	Source range overlaps the other range entirely: we just remove the other range.
					excludedRanges = append(excludedRanges, otherRange)
					break
				} else if sourceRange.From < otherRange.From && sourceRange.To > otherRange.From {
					//	Source range overlaps the other range partially from the left: extend the source range
					newRanges = append(newRanges, IngredientRange{From: sourceRange.From, To: otherRange.To})
					excludedRanges = append(excludedRanges, sourceRange)
					excludedRanges = append(excludedRanges, otherRange)
				} else if sourceRange.From < otherRange.To && sourceRange.To > otherRange.To {
					//	Source range overlaps the other range partially from the right: extend the other range
					newRanges = append(newRanges, IngredientRange{From: otherRange.From, To: sourceRange.To})
					excludedRanges = append(excludedRanges, sourceRange)
					excludedRanges = append(excludedRanges, otherRange)
				} else if sourceRange.To+1 == otherRange.From {
					//	Source range touches the other range from the right: extend the other range
					newRanges = append(newRanges, IngredientRange{From: sourceRange.From, To: otherRange.To})
					excludedRanges = append(excludedRanges, sourceRange)
					excludedRanges = append(excludedRanges, otherRange)
				}
			}

			fmt.Printf("Updated ranges %d\n", len(newRanges))
		}

		for _, outdatedRange := range excludedRanges {
			removeAt(copyRange, outdatedRange)
		}

		for _, newRange := range newRanges {
			*copyRange = append(*copyRange, newRange)
		}

		compactingDone = len(newRanges) == 0 && len(excludedRanges) == 0
	}

	return &FreshIngredients{Ranges: *copyRange}
}

func exists(
	existingRanges *[]IngredientRange,
	ingredientRange IngredientRange,
) (bool, int) {
	for index, existingRange := range *existingRanges {
		if existingRange == ingredientRange {
			return true, index
		}
	}

	return false, -1
}

func removeAt(
	ranges *[]IngredientRange,
	removedRange IngredientRange,
) {

	isFound, index := exists(ranges, removedRange)

	if !isFound || index < 0 || index >= len(*ranges) {
		return
	}

	*ranges = append((*ranges)[:index], (*ranges)[index+1:]...)
}
