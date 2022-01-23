func minFallingPathSum(A [][]int) int {
    n := len(A)
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            pre := A[i-1][j]
            if j >= 1 {
                pre = min(pre, A[i-1][j-1])
            }
            if j < n-1 {
                pre = min(pre, A[i-1][j+1])
            }
            A[i][j] += pre
        }
    }
    ans := A[n-1][0]
    for j := 1; j < n; j++ {
        ans = min(ans, A[n-1][j])
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}