func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    ans := 0
    m, n := len(matrix), len(matrix[0])
    depth := make([][]int, m)
    for i := range depth {
        depth[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            ans = max(ans, dfs(matrix, depth, i, j))
        }
    }
    return ans
}

func dfs(matrix, depth [][]int, i, j int) int {
    m, n := len(matrix), len(matrix[0])
    if depth[i][j] != 0 {
        return depth[i][j]
    }
    dir := [][2]int{[2]int{-1, 0}, [2]int{0, -1}, [2]int{0, 1}, [2]int{1, 0}}
    for _, d := range dir {
        x, y := i+d[0], j+d[1]
        if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[i][j] {
            depth[i][j] = max(depth[i][j], dfs(matrix, depth, x, y))
        }
    }
    depth[i][j]++
    return depth[i][j]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}