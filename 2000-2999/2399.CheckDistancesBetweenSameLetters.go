func checkDistances(s string, distance []int) bool {
    idx := make(map[byte]int, 26)
    for i, ch := range []byte(s) {
        if j, ok := idx[ch]; ok {
            if i-j-1 != distance[ch-'a'] {
                return false
            }
        } else {
            idx[ch] = i
        }
    }
    return true
}