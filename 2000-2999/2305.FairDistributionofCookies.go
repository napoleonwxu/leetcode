import "math"

var ans int

func distributeCookies(cookies []int, k int) int {
	// O(k^n)
	ans = math.MaxInt32
	dis := make([]int, k)
	dfs(cookies, dis, k, 0)
	return ans
}

func dfs(cookies, dis []int, k, idx int) {
	if idx == len(cookies) {
		tmp := 0
		for _, d := range dis {
			tmp = max(tmp, d)
		}
		ans = min(ans, tmp)
		return
	}
	for i := range dis {
		dis[i] += cookies[idx]
		dfs(cookies, dis, k, idx+1)
		dis[i] -= cookies[idx]
	}
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