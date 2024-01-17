func circularGameLosers(n int, k int) []int {
    rcv := make([]bool, n)
    for idx, i := 0, 1; !rcv[idx]; i++ {
        rcv[idx] = true
        idx = (idx + i*k) % n
    }
    ans := make([]int, 0, n)
    for i, r := range rcv {
        if !r {
            ans = append(ans, i+1)
        }
    }
    return ans
}