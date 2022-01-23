func countBinarySubstrings(s string) int {
    n := len(s)
    cnt := 0
    // O(n) + O(1)
    pre, cur := 0, 1
    for i := 1; i < n; i++ {
        if s[i] == s[i-1] {
            cur++
        } else {
            cnt += min(pre, cur)
            pre, cur = cur, 1
        }
    }
    return cnt + min(pre, cur)
    // O(n*n) + O(1)
    /*
    for i := 0; i < n; i++ {
        l, r := i-1, i
        for l >= 0 && r < n {
            if s[r] == s[i] && s[l] != s[i] {
                cnt++
                r++
                l--
            } else {
                break
            }
        }
    }
    return cnt
    */
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}