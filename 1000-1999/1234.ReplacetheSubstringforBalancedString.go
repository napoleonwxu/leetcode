func balancedString(s string) int {
    n := len(s)
    cnt := make(map[byte]int, 4)
    for i := 0; i < n; i++ {
        cnt[s[i]]++
    }
    i, ans := 0, n
    for j := 0; j < n; j++ {
        cnt[s[j]]--
        for i < n && cnt['Q'] <= n/4 && cnt['W'] <= n/4 && cnt['E'] <= n/4 && cnt['R'] <= n/4 {
            ans = min(ans, j-i+1)
            cnt[s[i]]++
            i++
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}