const INF = 200 * int(1e6)

func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	if n <= 1 || k >= n {
		return 0
	}
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(nums, memo, 0, n, k)
}

func dp(nums []int, memo [][]int, i, n, k int) int {
	if i == n {
		return 0
	}
	if k < 0 {
		return INF
	}
	if memo[i][k] != -1 {
		return memo[i][k]
	}
	ans := INF
	maxNum, sum := nums[i], 0
	for j := i; j < n; j++ {
		maxNum = max(maxNum, nums[j])
		sum += nums[j]
		wasted := maxNum*(j-i+1) - sum
		ans = min(ans, dp(nums, memo, j+1, n, k-1)+wasted)
	}
	memo[i][k] = ans
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
