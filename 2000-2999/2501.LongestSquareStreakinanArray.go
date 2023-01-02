func longestSquareStreak(nums []int) int {
	n := len(nums)
	sq := make(map[int]int, n)
	for _, num := range nums {
		sq[num] = 1
	}
	ans := -1
	visited := make(map[int]bool, n)
	for _, num := range nums {
		ans = max(ans, dfs(sq, visited, num))
	}
	if ans != 1 {
		return ans
	}
	return -1
}

func dfs(sq map[int]int, visited map[int]bool, cur int) int {
	if visited[cur] || sq[cur] == 0 {
		return sq[cur]
	}
	visited[cur] = true
	sq[cur] = dfs(sq, visited, cur*cur) + 1
	return sq[cur]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}