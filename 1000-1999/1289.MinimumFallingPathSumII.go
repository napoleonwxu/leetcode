func minFallingPathSum(arr [][]int) int {
    n := len(arr)
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            pre := math.MaxInt32
            for k := 0; k < n; k++ {
                if k == j {
                    continue
                }
                pre = min(pre, arr[i-1][k])
            }
            arr[i][j] += pre
        }
    }
    ans := arr[n-1][0]
    for j := 1; j < n; j++ {
        ans = min(ans, arr[n-1][j])
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}