var graph [][]int
var visited []bool

func frogPosition(n int, edges [][]int, t int, target int) float64 {
    if n == 1 {
        return 1
    }
    graph = make([][]int, n+1)
    visited = make([]bool, n+1)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    return dfs(1, t, target)
}

func dfs(node, time, target int) float64 {
    if (node != 1 && len(graph[node]) == 1) || time == 0 {
        if node == target {
            return 1
        }
        return 0
    }
    visited[node] = true
    p := 0.0
    for _, next := range graph[node] {
        if !visited[next] {
            p += dfs(next, time-1, target)
        }
    }
    if node != 1 {
        return p/float64(len(graph[node]) - 1)
    }
    return p/float64(len(graph[node]))
}