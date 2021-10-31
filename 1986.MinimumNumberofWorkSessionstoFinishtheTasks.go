import "math"

var target int

func minSessions(tasks []int, sessionTime int) int {
	n, sum := len(tasks), 0
	for _, task := range tasks {
		sum += task
	}
	target = 1<<n - 1
	dp := make([][]int, target+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return help(tasks, sessionTime, 0, 0, dp)
}

func help(tasks []int, sessionTime, mask, sum int, dp [][]int) int {
	if mask == target {
		return 1
	}
	if dp[mask][sum] != -1 {
		return dp[mask][sum]
	}
	ans, newMask := math.MaxInt32, 0
	for i, task := range tasks {
		if mask&(1<<i) != 0 {
			continue
		}
		newMask = mask | (1 << i)
		if sum+task > sessionTime {
			ans = min(ans, help(tasks, sessionTime, newMask, task, dp)+1)
		} else {
			ans = min(ans, help(tasks, sessionTime, newMask, sum+task, dp))
		}
	}
	dp[mask][sum] = ans
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}