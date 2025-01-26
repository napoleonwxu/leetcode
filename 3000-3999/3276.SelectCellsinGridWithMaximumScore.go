func maxScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    nums := make([][]int, 0, m*n)
    for i, g := range grid {
        for j, num := range g {
            nums = append(nums, []int{num, i, j})
        }
    }
    sort.Slice(nums, func(i, j int) bool {
        return nums[i][0] > nums[j][0]
    })
    dp := make(map[int]int, m*n)
    return dfs(nums, 0, 0, dp)
}

func dfs(nums [][]int, i, mask int, dp map[int]int) int {
    n := len(nums)
    if i >= n {
        return 0
    }
    if ans, ok := dp[mask*100+i]; ok {
        return ans
    }
    ans, row := 0, nums[i][1]
    if ((1<<row) & mask) > 0 {
        ans += dfs(nums, i+1, mask, dp)
    } else {
        j := i
        for ; j < n && nums[i][0] == nums[j][0]; j++ {}
        ans1 := nums[i][0] + dfs(nums, j, mask | (1<<row), dp)
        ans2 := dfs(nums, i+1, mask, dp)
        ans = max(ans1, ans2)
    }
    dp[mask*100+i] = ans
    return ans
}