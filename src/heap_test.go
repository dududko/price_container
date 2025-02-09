package src

import (
	"container/heap"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := &MaxHeap{}
	heap.Push(h, &CompanyPrice{price: 1})
	heap.Push(h, &CompanyPrice{price: 4})
	heap.Push(h, &CompanyPrice{price: 6})
	if h.MinHeap[0].price != 6 {
		t.Fatalf(`max = %d, want 6`, h.MinHeap[0].price)
	}

	heap.Push(h, &CompanyPrice{price: 10})
	if h.MinHeap[0].price != 10 {
		t.Fatalf(`max = %d, want 10`, h.MinHeap[0].price)
	}
	cp := heap.Pop(h).(*CompanyPrice)
	if cp.price != 10 {
		t.Fatalf(`max = %d, want 10`, cp.price)
	}
}

func TestMinHeap(t *testing.T) {
	h := &MinHeap{}
	heap.Push(h, &CompanyPrice{price: 1})
	heap.Push(h, &CompanyPrice{price: 4})
	heap.Push(h, &CompanyPrice{price: 6})

	cp := heap.Pop(h).(*CompanyPrice)
	if cp.price != 1 {
		t.Fatalf(`max = %d, want 1`, cp.price)
	}
}
