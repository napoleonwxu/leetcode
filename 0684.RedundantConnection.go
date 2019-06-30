func findRedundantConnection(edges [][]int) []int {
    // Union Find, O(n)
    n := len(edges)
    parent := make([]int, n+1)
    for i := range parent {
        parent[i] = i
    }
    for _, edge := range edges {
        pf := find(parent, edge[0])
        pt := find(parent, edge[1])
        if pf == pt {
            return edge
        }
        parent[pf] = pt
    }
    /* DFS, O(n*n)
    graph := make([][]int, 1001)
    for i := range graph {
        graph[i] = []int{}
    }
    for _, edge := range edges {
        Map := make(map[int]int)
        u, v := edge[0], edge[1]
        if len(graph[u]) > 0 && len(graph[v]) > 0 && dfs(graph, Map, u, v) {
            return edge
        }
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    */
    return []int{}
}

func find(parent []int, f int) int {
    if f != parent[f] {
        parent[f] = find(parent, parent[f])
    }
    return parent[f]
}

func dfs(graph [][]int, Map map[int]int, src, dst int) bool {
    if Map[src] == 1 {
        return false
    }
    Map[src] = 1
    if src == dst {
        return true
    }
    for _, nei := range graph[src] {
        if dfs(graph, Map, nei, dst) {
            return true
        }
    }
    return false
}