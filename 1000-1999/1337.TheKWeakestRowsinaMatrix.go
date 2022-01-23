func kWeakestRows(mat [][]int, k int) []int {
    m, n := len(mat), len(mat[0])
    order := make([][]int, m)
    for i := 0; i < m; i++ {
        order[i] = make([]int, 2)
        order[i][0] = i
        for j := 0; j < n; j++ {
            order[i][1] += mat[i][j]
        }
    }
    sort.Slice(order, func(i, j int) bool {
        if order[i][1] == order[j][1] {
            return order[i][0] < order[j][0]
        }
        return order[i][1] < order[j][1]
    })
    ans := make([]int, k)
    for i := 0; i < k; i++ {
        ans[i] = order[i][0]
    }
    return ans
}