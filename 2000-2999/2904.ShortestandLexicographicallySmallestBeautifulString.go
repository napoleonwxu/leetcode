import "sort"

func shortestBeautifulSubstring(s string, k int) string {
    ans := make([]string, 0)
    oneCnt, minLen := 0, len(s)
    for l, r := 0, 0; r < len(s); r++ {
        if s[r] == '1' {
            oneCnt++
        }
        for oneCnt == k {
            if r-l+1 <= minLen {
                if r-l+1 < minLen {
                    ans = make([]string, 0)
                }
                minLen = r - l + 1
                ans = append(ans, s[l:l+minLen])
            }
            if s[l] == '1' {
                oneCnt--
            }
            l++
        }
    }
    if len(ans) == 0 {
        return ""
    }
    sort.Strings(ans)
    return ans[0]
}