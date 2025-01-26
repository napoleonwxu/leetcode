func countGoodNodes(edges [][]int) int {
    n := len(edges) + 1
    graph := make([][]int, n)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    ans := 0
    dfs(graph, -1, 0, &ans)
    return ans
}

func dfs(graph [][]int, pre, cur int, ans *int) int {
    cnts := make([]int, 0, len(graph[cur]))
    for _, nxt := range graph[cur] {
        if nxt != pre {
            cnts = append(cnts, dfs(graph, cur, nxt, ans))
        }
    }
    sum := 0
    mi, ma := len(graph), -1
    for _, cnt := range cnts {
        sum += cnt
        mi = min(mi, cnt)
        ma = max(ma, cnt)
    }
    if len(cnts) <= 1 || mi == ma {
        *ans++
    }
    return sum + 1
}