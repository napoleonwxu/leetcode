func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*len(nums))
	for i := 0; i < n; i++ {
		ans[i], ans[i+n] = nums[i], nums[i]
	}
	return ans
}