func islandPerimeter(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    cnt, nei := 0, 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                cnt++
                if i+1 < m && grid[i+1][j] == 1 {
                    nei++
                }
                if j+1 < n && grid[i][j+1] == 1 {
                    nei++
                }
            }
        }
    }
    return 4*cnt - 2*nei
}