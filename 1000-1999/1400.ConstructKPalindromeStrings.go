func canConstruct(s string, k int) bool {
    if k > len(s) {
        return false
    }
    cnts := make([]int, 26)
    for _, ch := range s {
        cnts[ch-'a']++
    }
    odd := 0
    for _, cnt := range cnts {
        odd += cnt&1
    }
    return odd <= k
}