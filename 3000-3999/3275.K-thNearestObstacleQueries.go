type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func resultsArray(queries [][]int, k int) []int {
    ans := make([]int, 0, len(queries))
    h := &maxHeap{}
    heap.Init(h)
    for _, q := range queries {
        heap.Push(h, abs(q[0])+abs(q[1]))
        if h.Len() > k {
            heap.Pop(h)
        }
        if h.Len() < k {
            ans = append(ans, -1)
        } else {
            ans = append(ans, (*h)[0])
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}