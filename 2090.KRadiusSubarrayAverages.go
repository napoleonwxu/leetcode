func getAverages(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	if n <= 2*k {
		return ans
	}
	sum := 0
	for i := 0; i <= 2*k; i++ {
		sum += nums[i]
	}
	for i := k; i < n-k; i++ {
		ans[i] = sum / (2*k + 1)
		if i < n-k-1 {
			sum += nums[i+k+1] - nums[i-k]
		}
	}
	return ans
}