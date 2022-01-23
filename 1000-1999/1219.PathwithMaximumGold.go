func getMaximumGold(grid [][]int) int {
    ans := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] > 0 {
                dfs(grid, i, j, 0, &ans)
            }
        }
    }
    return ans
}

func dfs(grid [][]int, i, j, sum int, ans *int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
        *ans = max(*ans, sum)
        return
    }
    tmp := grid[i][j]
    grid[i][j] = 0
    dfs(grid, i-1, j, sum+tmp, ans)
    dfs(grid, i+1, j, sum+tmp, ans)
    dfs(grid, i, j-1, sum+tmp, ans)
    dfs(grid, i, j+1, sum+tmp, ans)
    grid[i][j] = tmp
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}