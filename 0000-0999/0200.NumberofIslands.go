func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    cnt := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                cnt++
                dfs(grid, i, j)
            }
        }
    }
    return cnt
}

func dfs(grid [][]byte, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
        return
    }
    grid[i][j] = '#'
    dfs(grid, i-1, j)
    dfs(grid, i+1, j)
    dfs(grid, i, j-1)
    dfs(grid, i, j+1)
}