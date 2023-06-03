import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	pairs := make([][]int, len(nums1))
	for i := range nums1 {
		pairs[i] = []int{nums1[i], nums2[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})
	h := &MinHeap{}
	heap.Init(h)
	score, sum := int64(0), int64(0)
	for _, pair := range pairs {
		sum += int64(pair[0])
		heap.Push(h, pair[0])
		if h.Len() > k {
			sum -= int64(heap.Pop(h).(int))
		}
		if h.Len() == k {
			score = max(score, sum*int64(pair[1]))
		}
	}
	return score
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}