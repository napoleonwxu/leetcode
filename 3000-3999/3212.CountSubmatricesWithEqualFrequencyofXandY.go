func numberOfSubmatrices(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    x := make([][]int, m+1)
    y := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        x[i] = make([]int, n+1)
        y[i] = make([]int, n+1)
    }
    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            x[i+1][j+1] = x[i][j+1] + x[i+1][j] - x[i][j]
            if grid[i][j] == 'X' {
                x[i+1][j+1]++
            }
            y[i+1][j+1] = y[i][j+1] + y[i+1][j] - y[i][j]
            if grid[i][j] == 'Y' {
                y[i+1][j+1]++
            }
            if x[i+1][j+1] > 0 && x[i+1][j+1] == y[i+1][j+1] {
                cnt++
            }
        }
    }
    return cnt
}
