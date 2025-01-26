func shortestSubstrings(arr []string) []string {
    m := make(map[string]int)
    for _, s := range arr {
        for i := 0; i < len(s); i++ {
            for j := i+1; j <= len(s); j++ {
                m[s[i:j]]++
            }
        }
    }
    ans := make([]string, len(arr))
    for k, s := range arr {
        tmp := make(map[string]int)
        for i := 0; i < len(s); i++ {
            for j := i+1; j <= len(s); j++ {
                tmp[s[i:j]]++
                if m[s[i:j]]-tmp[s[i:j]] <= 0 && (ans[k] == "" || j-i < len(ans[k]) || (j-i == len(ans[k]) && s[i:j] < ans[k])) {
                    ans[k] = s[i:j]
                }
            }
        }
    }
    return ans
}