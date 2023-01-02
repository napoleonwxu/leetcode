func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			dp[i+1] = dp[i-1]
			if i >= 2 && nums[i] == nums[i-2] {
				dp[i+1] = dp[i+1] || dp[i-2]
			}
		} else if i >= 2 && nums[i] == nums[i-1]+1 && nums[i] == nums[i-2]+2 {
			dp[i+1] = dp[i-2]
		}
	}
	return dp[n]
}