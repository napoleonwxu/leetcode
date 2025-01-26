func minChanges(n int, k int) int {
    if n&k != k {
        return -1
    }
    cnt := 0
    for n > 0 {
        if n%2 == 1 && k%2 == 0 {
            cnt++
        }
        n, k = n/2, k/2
    }
    return cnt
}