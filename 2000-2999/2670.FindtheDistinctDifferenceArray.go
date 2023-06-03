func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	pre, suf := make(map[int]int, n), make(map[int]int, n)
	for _, num := range nums {
		suf[num]++
	}
	ans := make([]int, n)
	for i, num := range nums {
		suf[num]--
		if suf[num] == 0 {
			delete(suf, num)
		}
		pre[num]++
		ans[i] = len(pre) - len(suf)
	}
	return ans
}