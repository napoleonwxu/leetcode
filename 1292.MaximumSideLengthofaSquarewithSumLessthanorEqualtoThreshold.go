func maxSideLength(mat [][]int, threshold int) int {
    // O(m*n*log(min(m,n)))
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
    left, right := 0, min(m, n)
    for left <= right {
        mid := left + (right-left)>>1
        if squareExist(sum, threshold, mid) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right
}

func squareExist(sum [][]int, threshold, l int) bool {
    for i := l; i < len(sum); i++ {
        for j := l; j < len(sum[0]); j++ {
            if sum[i][j] - sum[i-l][j] - sum[i][j-l] + sum[i-l][j-l] <= threshold {
                return true
            }
        }
    }
    return false
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}