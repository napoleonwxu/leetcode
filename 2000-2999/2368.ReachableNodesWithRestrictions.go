func reachableNodes(n int, edges [][]int, restricted []int) int {
    m := make(map[int][]int, n)
    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
        m[edge[1]] = append(m[edge[1]], edge[0])
    }
    visited := make(map[int]bool, n)
    for _, restrict := range restricted {
        visited[restrict] = true
    }
    return dfs(0, m, visited)
}

func dfs(i int, m map[int][]int, visited map[int]bool) int {
    if visited[i] {
        return 0
    }
    visited[i] = true
    cnt := 0
    for _, nxt := range m[i] {
        cnt += dfs(nxt, m, visited)
    }
    return cnt + 1
}