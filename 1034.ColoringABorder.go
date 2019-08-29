func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
    dfs(grid, r0, c0, grid[r0][c0])
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] < 0 {
                grid[i][j] = color
            }
        }
    }
    return grid
}

func dfs(grid [][]int, i, j, old int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != old {
        return
    }
    grid[i][j] = -old
    dfs(grid, i-1, j, old)
    dfs(grid, i+1, j, old)
    dfs(grid, i, j+1, old)
    dfs(grid, i, j-1, old)
    if i > 0 && i < len(grid)-1 && j > 0 && j < len(grid[0])-1 && abs(grid[i-1][j]) == old && abs(grid[i+1][j]) == old && abs(grid[i][j-1]) == old && abs(grid[i][j+1]) == old {
        grid[i][j] = old
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
