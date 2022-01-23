// Dijkstra + Heap: see python solution.
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    // Bellman-ford, O(n^2*K), at most n^2 edges
    costs := make([]int, n)
    for i := range costs {
        costs[i] = 1e8
    }
    costs[src] = 0
    for k := 0; k < K+1; k++ {
        tmp := make([]int, n)
        copy(tmp, costs)
        for _, f := range flights {
            s, d, p := f[0], f[1], f[2]
            tmp[d] = min(tmp[d], costs[s] + p)
        }
        costs = tmp
    }
    if costs[dst] == 1e8 {
        return -1
    }
    return costs[dst]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
