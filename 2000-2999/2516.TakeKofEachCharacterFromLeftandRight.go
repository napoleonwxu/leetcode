func takeCharacters(s string, k int) int {
    limits := make(map[byte]int, 3)
    for _, ch := range []byte(s) {
        limits[ch]++
    }
    for _, ch := range []byte{'a', 'b', 'c'} {
        limits[ch] -= k
        if limits[ch] < 0 {
            return -1
        }
    }
    cnts := make(map[byte]int, 3)
    ans, l := 0, 0
    for r, ch := range []byte(s) {
        cnts[ch]++
        for cnts[ch] > limits[ch] {
            cnts[s[l]]--
            l++
        }
        ans = max(ans, r-l+1)
    }
    return len(s) - ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
