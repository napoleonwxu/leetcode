func similarPairs(words []string) int {
    m := make(map[int]int, len(words))
    cnt := 0
    for _, word := range words {
        mask := 0
        for _, ch := range word {
            mask |= 1 << (ch - 'a')
        }
        cnt += m[mask]
        m[mask]++
    }
    return cnt
}