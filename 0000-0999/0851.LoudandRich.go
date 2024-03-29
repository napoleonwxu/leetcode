func loudAndRich(richer [][]int, quiet []int) []int {
    n := len(quiet)
    graph := make([][]int, n)
    for _, r := range richer {
        graph[r[1]] = append(graph[r[1]], r[0])
    }
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    for i := 0; i < n; i++ {
        dfs(graph, quiet, ans, i)
    }
    return ans
}

func dfs(graph [][]int, quiet, ans []int, i int) int {
    if ans[i] != -1 {
        return ans[i]
    }
    ans[i] = i
    for _, j := range graph[i] {
        cand := dfs(graph, quiet, ans, j)
        if quiet[cand] < quiet[ans[i]] {
            ans[i] = cand
        }
    }
    return ans[i]
}