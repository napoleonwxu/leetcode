func reorganizeString(S string) string {
    n := len(S)
    cnt := make([]int, 26)
    for _, ch := range S {
        cnt[ch-'a'] += 100
    }
    for i := 0; i < 26; i++ {
        cnt[i] += i
    }
    sort.Ints(cnt)
    
    ans := make([]string, n)
    i := 1
    for _, code := range cnt {
        c := code/100
        ch := string('a'+code%100)
        if c > (n+1)>>1 {
            return ""
        }
        for ; c > 0; c-- {
            if i >= n {
                i = 0
            }
            ans[i] = ch
            i += 2
        }
    }
    return strings.Join(ans, "")
}