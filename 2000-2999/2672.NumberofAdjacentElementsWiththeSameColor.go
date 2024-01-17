func colorTheArray(n int, queries [][]int) []int {
    color := make([]int, n)
    ans := make([]int, len(queries))
    cur := 0
    for i, q := range queries {
        if q[0] > 0 && color[q[0]] != 0 && color[q[0]] == color[q[0]-1] {
            cur--
        }
        if q[0] < n-1 && color[q[0]] != 0 && color[q[0]] == color[q[0]+1] {
            cur--
        }
        color[q[0]] = q[1]
        if q[0] > 0 && color[q[0]] == color[q[0]-1] {
            cur++
        }
        if q[0] < n-1 && color[q[0]] == color[q[0]+1] {
            cur++
        }
        ans[i] = cur
    }
    return ans
}