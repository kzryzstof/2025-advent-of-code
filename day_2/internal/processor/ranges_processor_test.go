package processor

import (
	"day_2/internal/abstractions"
	"sync"
	"testing"
)

// fakeRangesChannel implements abstractions.RangesChannel for testing.
type fakeRangesChannel struct {
	ch chan abstractions.Range
}

func (f *fakeRangesChannel) Ranges() <-chan abstractions.Range {
	return f.ch
}

func TestRangesProcessor_GetTotalProductId(t *testing.T) {
	// Prepare a channel and fake RangesChannel implementation.
	ch := make(chan abstractions.Range)
	fake := &fakeRangesChannel{ch: ch}

	var wg sync.WaitGroup
	processor := NewProcessor(fake, &wg)

	// Start the processor which will read from ch.
	processor.Start()

	// Send all documented ranges in a separate goroutine and then close the channel.
	go func() {
		defer close(ch)

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

		ch <- mkRange("11", "22")
		ch <- mkRange("95", "115")
		ch <- mkRange("998", "1012")
		ch <- mkRange("1188511880", "1188511890")
		ch <- mkRange("222220", "222224")
		ch <- mkRange("1698522", "1698528")
		ch <- mkRange("446443", "446449")
		ch <- mkRange("38593856", "38593862")
		ch <- mkRange("565653", "565659")
		ch <- mkRange("824824821", "824824827")
		ch <- mkRange("2121212118", "2121212124")
	}()

	// Wait for the processor to finish consuming the channel.
	wg.Wait()

	// Expected total is the sum of all invalid IDs in the example:
	// 11 + 22 + 99 + 111 + 999 + 1010 + 1188511885 + 222222 +
	// 446446 + 38593859 + 565656 + 824824824 + 2121212121 = 4174379265
	const expectedTotal int64 = 4174379265
	actualTotal := processor.GetTotalProductId()

	if actualTotal != expectedTotal {
		t.Fatalf("expected total product id %d, got %d", expectedTotal, actualTotal)
	}
}
