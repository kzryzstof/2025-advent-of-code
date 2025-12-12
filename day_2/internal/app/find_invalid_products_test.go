package app

import (
	"day_2/internal/abstractions"
	"testing"
)

func TestFindInvalidProductIds(t *testing.T) {
	// Helper to create a Range from string bounds.
	mkRange := func(from, to string) abstractions.Range {
		fromProd, err := abstractions.NewProduct(from)
		if err != nil {
			t.Fatalf("failed to create from product %q: %v", from, err)
		}
		toProd, err := abstractions.NewProduct(to)
		if err != nil {
			t.Fatalf("failed to create to product %q: %v", to, err)
		}
		return abstractions.Range{From: *fromProd, To: *toProd}
	}

	// All documented ranges from the example.
	ranges := []abstractions.Range{
		mkRange("11", "22"),
		mkRange("95", "115"),
		mkRange("998", "1012"),
		mkRange("1188511880", "1188511890"),
		mkRange("222220", "222224"),
		mkRange("1698522", "1698528"),
		mkRange("446443", "446449"),
		mkRange("38593856", "38593862"),
		mkRange("565653", "565659"),
		mkRange("824824821", "824824827"),
		mkRange("2121212118", "2121212124"),
	}

	// Compute the sum of invalid product IDs using the application helper.
	actualTotal := FindInvalidProductIds(ranges)

	// Expected total is the sum of all invalid IDs in the example:
	// 11 + 22 + 99 + 111 + 999 + 1010 + 1188511885 + 222222 +
	// 446446 + 38593859 + 565656 + 824824824 + 2121212121 = 4174379265
	const expectedTotal uint64 = 4174379265

	if actualTotal != expectedTotal {
		t.Fatalf("expected total product id %d, got %d", expectedTotal, actualTotal)
	}
}
