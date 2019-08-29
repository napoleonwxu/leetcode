func firstUniqChar(s string) int {
    idx := make([]int, 26)
    for _, ch := range s {
        idx[ch-'a']++
    }
    for i, ch := range s {
        if idx[ch-'a'] == 1 {
            return i
        }
    }
    return -1
}