func shiftGrid(grid [][]int, k int) [][]int {
    n, m := len(grid), len(grid[0])
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            idx := (i*m + j + k) % (n*m)
            ans[idx/m][idx%m] = grid[i][j]
        }
    }
    return ans
}