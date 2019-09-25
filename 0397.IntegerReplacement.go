func integerReplacement(n int) int {
    cnt := 0
    for n > 1 {
        if n&1 == 0 {
            n >>= 1
        } else if n == 3 || (n>>1)&1 == 0 {
            n--
        } else {
            n++
        }
        cnt++
    }
    return cnt
}