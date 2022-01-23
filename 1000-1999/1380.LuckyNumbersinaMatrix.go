func luckyNumbers (matrix [][]int) []int {
    // O(mn) + O(m+n)
    m, n := len(matrix), len(matrix[0])
    minRow, maxCol := make([]int, m), make([]int, n)
    for i := 0; i < m; i++ {
        minRow[i] = 1e5 + 1
        for j := 0; j < n; j++ {
            minRow[i] = min(minRow[i], matrix[i][j])
            maxCol[j] = max(maxCol[j], matrix[i][j])
        }
    }
    Map := make(map[int]bool)
    for _, m := range minRow {
        Map[m] = true
    }
    ans := []int{}
    for _, m := range maxCol {
        if Map[m] {
            ans = append(ans, m)
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}