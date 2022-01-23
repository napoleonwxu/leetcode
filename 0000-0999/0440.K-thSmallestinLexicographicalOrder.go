func findKthNumber(n int, k int) int {
    cur, k := 1, k-1
    for k > 0 {
        step := count(n, cur, cur+1)
        if step <= k {
            cur++
            k -= step
        } else {
            cur *= 10
            k--
        }
    }
    return cur
}

func count(n, n1, n2 int) int {
    cnt := 0
    for n1 <= n {
        cnt += min(n+1, n2) - n1
        n1, n2 = 10*n1, 10*n2
    }
    return cnt
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
