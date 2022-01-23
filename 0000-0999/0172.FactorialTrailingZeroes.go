func trailingZeroes(n int) int {
    cnt := 0
    for n > 0 {
        n /= 5
        cnt += n
    }
    return cnt
}