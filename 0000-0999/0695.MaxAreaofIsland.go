func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    area := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            area = max(area, dfs(grid, i, j))
        }
    }
    return area
}

func dfs(grid [][]int, i, j int) int {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
        return 0
    }
    grid[i][j] = -1
    return 1 + dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}