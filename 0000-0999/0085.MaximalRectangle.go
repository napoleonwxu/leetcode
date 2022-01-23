func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    m, n := len(matrix), len(matrix[0])
    left, right, height := make([]int, n), make([]int, n), make([]int, n)
    for i := range right {
        right[i] = n
    }
    ans := 0
    for i := 0; i < m; i++ {
        idx_l, idx_r := 0, n
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                height[j]++
                left[j] = max(left[j], idx_l)
            } else {
                height[j] = 0
                left[j], idx_l = 0, j+1
            }
        }
        for j := n-1; j >= 0; j-- {
            if matrix[i][j] == '1' {
                right[j] = min(right[j], idx_r)
            } else {
                right[j], idx_r = n, j
            }
            ans = max(ans, (right[j]-left[j])*height[j])
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