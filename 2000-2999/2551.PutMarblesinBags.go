import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	tmp := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		tmp[i] = weights[i] + weights[i+1]
	}
	sort.Ints(tmp)
	mi, ma := int64(0), int64(0)
	for i := 0; i < k-1; i++ {
		mi += int64(tmp[i])
		ma += int64(tmp[n-2-i])
	}
	return ma - mi
}