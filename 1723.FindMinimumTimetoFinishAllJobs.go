var ans int

func minimumTimeRequired(jobs []int, k int) int {
	// DFS, O(k^n)
	ans = sum(jobs)
	// sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	times := make([]int, k)
	dfs(jobs, times, k, 0)
	return ans
}

func dfs(jobs, times []int, k, cur int) {
	if cur >= len(jobs) {
		ans = min(ans, max(times))
		return
	}
	seen := make(map[int]bool)
	for i := 0; i < k; i++ {
		if seen[times[i]] {
			continue
		}
		if times[i]+jobs[cur] >= ans {
			continue
		}
		seen[times[i]] = true
		times[i] += jobs[cur]
		dfs(jobs, times, k, cur+1)
		times[i] -= jobs[cur]
	}
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func max(nums []int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
