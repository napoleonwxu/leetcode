func countServers(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    row, col := make([]int, m), make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                row[i]++
                col[j]++
            }
        }
    }
    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (row[i]>1 || col[j]>1) {
                cnt++
            }
        }
    }
    return cnt
}