func longestCycle(edges []int) int {
    ans, n := -1, len(edges)
    visited := make([]bool, n)
    for i := 0; i < n; i++ {
        steps := make(map[int]int)
        node, step := i, 0
        for node != -1 {
            if pre, ok := steps[node]; ok {
                ans = max(ans, step-pre)
                break
            }
            if visited[node] {
                break
            } else {
                visited[node] = true
            }
            steps[node] = step
            step++
            node = edges[node]
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
