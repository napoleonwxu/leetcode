func minIncrements(n int, cost []int) int {
	ans := 0
	for i := n/2 - 1; i >= 0; i-- {
		ans += abs(cost[2*i+1] - cost[2*i+2])
		cost[i] += max(cost[2*i+1], cost[2*i+2])
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}