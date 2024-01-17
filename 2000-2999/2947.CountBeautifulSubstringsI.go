func beautifulSubstrings(s string, k int) int {
    n, cnt := len(s), 0
    for i := 0; i < n; i++ {
        vol, con := 0, 0
        for j := i; j < n; j++ {
            if s[j] == 'a' || s[j] == 'e' || s[j] == 'i' || s[j] == 'o' || s[j] == 'u' {
                vol++
            } else {
                con++
            }
            if vol == con && (vol*con)%k == 0 {
                cnt++
            }
        }
    }
    return cnt
}
