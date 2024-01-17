func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
    graph := make([][]int, len(values))
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }
    sum := int64(0)
    for _, value := range values {
        sum += int64(value)
    }
    return sum - dfs(graph, values, 0, -1)
}

func dfs(graph [][]int, values []int, node, parent int) int64 {
    if node != 0 && len(graph[node]) == 1 {
        return int64(values[node])
    }
    sum := int64(0)
    for _, child := range graph[node] {
        if child == parent {
            continue
        }
        sum += dfs(graph, values, child, node)
    }
    return min(sum, int64(values[node]))
}

func min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}