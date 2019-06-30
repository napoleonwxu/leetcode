// https://leetcode.com/submissions/detail/77520312/ solution is better
func longestPalindrome(s string) string {
    n := len(s)
    idx, maxlen := 0, 0
    for i := 0; i < n; i++ {
        if n - i <= maxlen>>1 {
            break
        }
        expand(s, i, i, &idx, &maxlen)
        expand(s, i, i+1, &idx, &maxlen)
    }
    return s[idx:idx+maxlen]
}

func expand(s string, i, j int, idx, maxlen *int) {
    for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {}
    if j-i-1 > *maxlen {
        *idx = i+1
        *maxlen = j-i-1
    }
} 