func beautifulSubarrays(nums []int) int64 {
	dp := make(map[int]int)
	dp[0] = 1
	pre, cnt := 0, int64(0)
	for _, num := range nums {
		pre ^= num
		cnt += int64(dp[pre])
		dp[pre]++
	}
	return cnt
}