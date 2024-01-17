var score int

func minScore(n int, roads [][]int) int {
    score = int(1e4)
    graph := make(map[int]map[int]int)
    for i := 1; i <= n; i++ {
        graph[i] = make(map[int]int)
    }
    for _, r := range roads {
        graph[r[0]][r[1]] = r[2]
        graph[r[1]][r[0]] = r[2]
    }
    visited := make(map[int]bool, n)
    dfs(graph, visited, 1, n)
    return score
}

func dfs(graph map[int]map[int]int, visited map[int]bool, cur, n int) {
    if visited[cur] {
        return
    }
    visited[cur] = true
    for nxt, dis := range graph[cur] {
        score = min(score, dis)
        dfs(graph, visited, nxt, n)
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}