import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	n := len(special)
	ans := max(special[0]-bottom, top-special[n-1])
	for i := 1; i < n; i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}