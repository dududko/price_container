package src

type MinHeap []*CompanyPrice

type MaxHeap struct {
	MinHeap
}

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].price < h[j].price }

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].id = i
	h[j].id = j
}

func (h *MinHeap) Push(x any) {
	cp := x.(*CompanyPrice)
	cp.id = len(*h)
	*h = append(*h, cp)
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i].price > h.MinHeap[j].price }
