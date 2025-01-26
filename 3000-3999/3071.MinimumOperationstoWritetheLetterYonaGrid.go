func minimumOperationsToWriteY(grid [][]int) int {
    cnt := operations(grid, 0, 1)
    cnt = min(cnt, operations(grid, 0, 2))
    cnt = min(cnt, operations(grid, 1, 0))
    cnt = min(cnt, operations(grid, 1, 2))
    cnt = min(cnt, operations(grid, 2, 0))
    cnt = min(cnt, operations(grid, 2, 1))
    return cnt
}

func operations(grid [][]int, a, b int) int {
    n, cnt := len(grid), 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if (i <= n/2 && i == j) || (i <= n/2 && i+j == n-1) || (i >= n/2 && j == n/2) {
                if grid[i][j] != a {
                    cnt++
                }
            } else {
                if grid[i][j] != b {
                    cnt++
                }
            }
        }
    }
    return cnt
}