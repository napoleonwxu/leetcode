func remainingMethods(n int, k int, invocations [][]int) []int {
    grid := make([][]int, n)
    for _, inv := range invocations {
        grid[inv[0]] = append(grid[inv[0]], inv[1])
    }
    visited := make([]bool, n)
    dfs(grid, visited, k)
    invoked := false
    for i := 0; i < n; i++ {
        if !visited[i] {
            for _, j := range grid[i] {
                if visited[j] {
                    invoked = true
                    break
                }
            }
        }
    }
    ans := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if visited[i] {
            if invoked {
                ans = append(ans, i)
            }
        } else {
            ans = append(ans, i)
        }
    }
    return ans
}

func dfs(grid [][]int, visited []bool, cur int) {
    if visited[cur] {
        return
    }
    visited[cur] = true
    for _, nxt := range grid[cur] {
        dfs(grid, visited, nxt)
    }
}