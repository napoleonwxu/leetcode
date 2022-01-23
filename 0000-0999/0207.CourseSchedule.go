func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[0]] = append(graph[p[0]], p[1])
    }
    visit := make([]int, numCourses)
    for i := 0; i < numCourses; i++ {
        if !dfs(graph, visit, i) {
            return false
        }
    }
    return true
}

func dfs(graph [][]int, visit []int, i int) bool {
    if visit[i] == -1 {
        return false
    }
    if visit[i] == 1 {
        return true
    }
    visit[i] = -1
    for _, j := range graph[i] {
        if !dfs(graph, visit, j) {
            return false
        }
    }
    visit[i] = 1
    return true
}