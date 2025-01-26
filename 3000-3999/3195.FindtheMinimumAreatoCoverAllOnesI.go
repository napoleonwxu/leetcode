func minimumArea(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    x1, x2 := m, -1
    y1, y2 := n, -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                x1 = min(x1, i)
                x2 = max(x2, i)
                y1 = min(y1, j)
                y2 = max(y2, j)
            }
        }
    }
    if x2 >= x1 && y2 >= y1 {
        return (x2-x1+1) * (y2-y1+1)
    }
    return 0
}