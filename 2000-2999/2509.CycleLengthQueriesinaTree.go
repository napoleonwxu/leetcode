func cycleLengthQueries(n int, queries [][]int) []int {
    ans := make([]int, len(queries))
    for i, q := range queries {
        ans[i] = 1
        for q[0] != q[1] {
            if q[0] > q[1] {
                q[0] /= 2
            } else {
                q[1] /= 2
            }
            ans[i]++
        }
    }
    return ans
}