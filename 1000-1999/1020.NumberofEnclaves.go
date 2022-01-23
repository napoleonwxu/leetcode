func numEnclaves(A [][]int) int {
    if len(A) <= 2 || len(A[0]) <= 2 {
        return 0
    }
    m, n := len(A), len(A[0])
    for i := 0; i < m; i++ {
        dfs(A, i, 0)
        dfs(A, i, n-1)
    }
    for j := 1; j < n-1; j++ {
        dfs(A, 0, j)
        dfs(A, m-1, j)
    }
    cnt := 0
    for i := 1; i < m-1; i++ {
        for j := 1; j < n-1; j++ {
            cnt += A[i][j]
        }
    }
    return cnt
}

func dfs(A [][]int, i, j int) {
    if i < 0 || i >= len(A) || j < 0 || j >= len(A[0]) || A[i][j] != 1 {
        return
    }
    A[i][j] = 0
    dfs(A, i-1, j)
    dfs(A, i+1, j)
    dfs(A, i, j-1)
    dfs(A, i, j+1)
}
