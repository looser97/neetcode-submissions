type IntHeap []int

func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	item := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return item
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Top() int {
	return h[0]
}

type KthLargest struct {
	h IntHeap
	k int
}

func Constructor(kt int, nums []int) KthLargest {
	ht := IntHeap{}
	heap.Init(&ht)
	for _, item := range nums {
		heap.Push(&ht, item)
		if ht.Len() > kt {
			heap.Pop(&ht)
		}
	}

	return KthLargest{
		h: ht,
		k: kt,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(&this.h)
	}
	return this.h.Top()
}
