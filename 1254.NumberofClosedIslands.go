func closedIsland(grid [][]int) int {
    if len(grid) <= 2 || len(grid[0]) <= 2 {
        return 0
    }
    cnt := 0
    for i := 1; i < len(grid)-1; i++ {
        for j := 1; j < len(grid[0])-1; j++ {
            if grid[i][j] == 0 && dfs(grid, i, j) {
                cnt++
            }
        }
    }
    return cnt
}

func dfs(grid [][]int, i, j int) bool {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return false
    }
    if grid[i][j] != 0 {
        return true
    }
    grid[i][j] = 2
    // wrong answer:
    //return dfs(grid, i-1, j) && dfs(grid, i+1, j) && dfs(grid, i, j-1) && dfs(grid, i, j+1)
    left := dfs(grid, i-1, j)
    right := dfs(grid, i+1, j)
    top := dfs(grid, i, j-1)
    bottom := dfs(grid, i, j+1)
    return left && right && top && bottom
}