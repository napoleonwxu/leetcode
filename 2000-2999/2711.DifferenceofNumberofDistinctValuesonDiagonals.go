func differenceOfDistinctValues(grid [][]int) [][]int {
    // O(mn)
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        calculate(grid, ans, i, 0)
    }
    for j := 1; j < n; j++ {
        calculate(grid, ans, 0, j)
    }
    return ans
}

func calculate(grid, ans [][]int, i, j int) {
    m, n := len(grid), len(grid[0])
    left, right := make(map[int]struct{}), make(map[int]struct{})
    for ; i < m && j < n; i, j = i+1, j+1 {
        ans[i][j] = len(left)
        left[grid[i][j]] = struct{}{}
    }
    for i, j = i-1, j-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        ans[i][j] = abs(ans[i][j] - len(right))
        right[grid[i][j]] = struct{}{}
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}