func findCircleNum(M [][]int) int {
    cnt := 0
    checked := make([]bool, len(M))
    for i := 0; i < len(M); i++ {
        if !checked[i] {
            cnt++
            dfs(M, i, checked)
        }
    }
    return cnt
}

func dfs(M [][]int, i int, checked []bool) {
    if checked[i] {
        return
    }
    checked[i] = true
    for j, r := range M[i] {
        if r == 1 {
            dfs(M, j, checked)
        }
    }
}