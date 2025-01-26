type Pair struct {
    num int
    cnt int64
}

type MaxHeap []Pair

func mostFrequentIDs(nums []int, freq []int) []int64 {
    n := len(nums)
    ans := make([]int64, n)
    cnts := make(map[int]int64)
    pq := make(MaxHeap, 0)
    heap.Init(&pq)
    for i := 0; i < n; i++ {
        cnts[nums[i]] += int64(freq[i])
        heap.Push(&pq, Pair{nums[i], cnts[nums[i]]})
        for len(pq) > 0 && pq[0].cnt != cnts[pq[0].num] {
            heap.Pop(&pq)
        }
        ans[i] = pq[0].cnt
    }
    return ans
}

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i].cnt > h[j].cnt
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
    tmp := *h
    n := len(tmp)
    res := tmp[n-1]
    *h = tmp[:n-1]
    return res
}