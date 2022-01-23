func matrixBlockSum(mat [][]int, K int) [][]int {
    // O(mn)
    m, n := len(mat), len(mat[0])
    sum := make([][]int, m+1)
    for i := range sum {
        sum[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1] 
        }
    }
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            r1, r2 := max(i-K, 0), min(i+K+1, m)
            c1, c2 := max(j-K, 0), min(j+K+1, n)
            ans[i][j] = sum[r2][c2] - sum[r1][c2] - sum[r2][c1] + sum[r1][c1]
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}