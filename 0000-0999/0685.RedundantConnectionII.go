func findRedundantDirectedConnection(edges [][]int) []int {
    // Union-Find, O(n)
    n := len(edges)
    parent := make([]int, n+1)
    var candA, candB []int
    for _, edge := range edges {
        if parent[edge[1]] == 0 {
            parent[edge[1]] = edge[0]
        } else {
            candA = []int{parent[edge[1]], edge[1]}
            candB = []int{edge[0], edge[1]}
            edge[1] = 0
        }
    }
    for i := range parent {
        parent[i] = i
    }
    for _, edge := range edges {
        if edge[1] == 0 {
            continue
        }
        f, t := edge[0], edge[1]
        pf := find(parent, f)
        if pf == t {
            if len(candA) == 0 {
                return edge
            } else {
                return candA
            }
        }
        parent[t] = pf
    }
    return candB
}

func find(parent []int, f int) int {
    if f != parent[f] {
        parent[f] = find(parent, parent[f])
    }
    return parent[f]
}