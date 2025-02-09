package src

import (
	"testing"
)

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestOriginPriceContainer(t *testing.T) {
	c := NewOriginPriceContainer(2)

	// same company, same date, new price -> update price
	c.InsertPrice(PriceBody{Company: 1, Price: 100, Origin: "origin", Date: "2020-01-02"})
	c.InsertPrice(PriceBody{Company: 1, Price: 21, Origin: "origin", Date: "2020-01-02"})
	if c.meanPrice != 20 && c.topNPrices.Len() != 1 && c.otherPrices.Len() != 0 {
		t.Fatalf(`meanPrice = %d, want 20`, c.meanPrice)
	}
	// same company, new date is older -> do not update price
	c.InsertPrice(PriceBody{Company: 1, Price: 100, Origin: "origin", Date: "2020-01-01"})
	if c.meanPrice != 21 {
		t.Fatalf(`meanPrice = %d, want 21`, c.meanPrice)
	}

	// add company 2, mean price is different
	c.InsertPrice(PriceBody{Company: 2, Price: 20, Origin: "origin", Date: "2020-01-01"})
	if c.meanPrice != 20 || c.topNPrices.Len() != 2 || c.otherPrices.Len() != 0 {
		t.Fatalf(`meanPrice = %d, want 20`, c.meanPrice)
	}

	// add company 3, which is cheapest, now company 2 is in otherPrices
	c.InsertPrice(PriceBody{Company: 3, Price: 10, Origin: "origin", Date: "2020-01-01"})
	if c.meanPrice != 15 || c.topNPrices.Len() != 2 || c.otherPrices.Len() != 1 {
		t.Fatalf(`meanPrice = %d, want 21; sizeTop = %d; size other = %d`, c.meanPrice, c.topNPrices.Len(), c.otherPrices.Len())
	}
	t.Logf("Pass %v %v %v", c.topNPrices.MinHeap[0], c.topNPrices.MinHeap[1], (*c.otherPrices)[0])

	// move company 2 back to topNPrices
	c.InsertPrice(PriceBody{Company: 2, Price: 11, Origin: "origin", Date: "2020-01-01"})
	if c.meanPrice != 10 || c.topNPrices.Len() != 2 || c.otherPrices.Len() != 1 {
		t.Fatalf(`meanPrice = %d, want 10`, c.meanPrice)
	}
}
