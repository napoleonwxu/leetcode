func heightChecker(heights []int) int {
    n := len(heights)
    inc := make([]int, n)
    copy(inc, heights)
    sort.Ints(inc)
    cnt := 0
    for i := 0; i < n; i++ {
        if inc[i] != heights[i] {
            cnt++
        }
    }
    return cnt
}