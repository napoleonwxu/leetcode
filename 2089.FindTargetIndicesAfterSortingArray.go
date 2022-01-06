func targetIndices(nums []int, target int) []int {
	smaller, same := 0, 0
	for _, num := range nums {
		if num < target {
			smaller++
		} else if num == target {
			same++
		}
	}
	ans := make([]int, same)
	for i := range ans {
		ans[i] = smaller + i
	}
	return ans
}