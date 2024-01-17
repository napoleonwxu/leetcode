func findChampion(n int, edges [][]int) int {
    indegree := make([]int, n)
    for _, e := range edges {
        indegree[e[1]]++
    }
    ans := -1
    for i, cnt := range indegree {
        if cnt == 0 {
            if ans != -1 {
                return -1
            }
            ans = i
        }
    }
    return ans
    /*
    grid := make([][]int, n)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    for _, e := range edges {
        grid[e[0]][e[1]] = 1
    }
    m := make(map[int]map[int]bool, n)
    for i := 0; i < n; i++ {
        dfs(grid, i, m)
        if len(m[i]) >= n-1 {
            return i
        }
    }
    return -1
    */
}

func dfs(grid [][]int, i int, m map[int]map[int]bool) {
    if _, ok := m[i]; ok {
        return
    }
    m[i] = make(map[int]bool)
    for j, num := range grid[i] {
        if num == 1 {
            m[i][j] = true
            dfs(grid, j, m)
            for k := range m[j] {
                m[i][k] = true
            }
        }
    }
}