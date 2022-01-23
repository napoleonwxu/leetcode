func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
    // BFS
    graph := make([][][]int, 2)
    graph[0] = make([][]int, n)
    graph[1] = make([][]int, n)
    for _, edge := range red_edges {
        graph[0][edge[0]] = append(graph[0][edge[0]], edge[1])
    }
    for _, edge := range blue_edges {
        graph[1][edge[0]] = append(graph[1][edge[0]], edge[1])
    }
    step := make([][]int, 2)
    step[0] = make([]int, n)
    step[1] = make([]int, n)
    for i := 1; i < n; i++ {
        step[0][i], step[1][i] = 2*n, 2*n
    }
    queue := [][]int{[]int{0, 0}, []int{0, 1}}
    for len(queue) > 0 {
        node, color := queue[0][0], queue[0][1]
        queue = queue[1:]
        for _, nei := range graph[color][node] {
            if step[color^1][nei] == 2*n {
                step[color^1][nei] = step[color][node] + 1
                queue = append(queue, []int{nei, color^1})
            } 
        }
    }
    ans := make([]int, n)
    for i := range ans {
        ans[i] = min(step[0][i], step[1][i])
        if ans[i] == 2*n {
            ans[i] = -1
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
