import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	h := MaxHeap(nums)
	heap.Init(&h)
	score := int64(0)
	for ; k > 0; k-- {
		score += int64(h[0])
		h[0] = (h[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return score
}
