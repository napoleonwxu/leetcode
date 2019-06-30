func possibleBipartition(N int, dislikes [][]int) bool {
    graph := make([][]int, N+1)
    for _, dis := range dislikes {
        graph[dis[0]] = append(graph[dis[0]], dis[1])
        graph[dis[1]] = append(graph[dis[1]], dis[0])
    }
    seen := make(map[int]int, N)
    for i := 1; i <= N; i++ {
        _, ok := seen[i]
        if !ok && !dfs(i, 0, graph, seen) {
            return false
        }
    }
    return true
}

func dfs(i int, v int, graph [][]int, seen map[int]int) bool {
    iv, ok := seen[i]
    if ok {
        return iv == v
    }
    seen[i] = v
    for _, dis := range graph[i] {
        if !dfs(dis, v^1, graph, seen) {
            return false
        }
    }
    return true
}
