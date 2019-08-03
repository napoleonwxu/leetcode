func transpose(A [][]int) [][]int {
    m, n := len(A), len(A[0])
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, m)
        for j := range ans[i] {
            ans[i][j] = A[j][i]
        }
    }
    return ans
}