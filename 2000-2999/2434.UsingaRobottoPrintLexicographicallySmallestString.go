func robotWithString(s string) string {
    chs := []byte(s)
    cnts := make(map[byte]int, 26)
    for _, ch := range chs {
        cnts[ch]++
    }

    n := len(chs)
    ans := make([]byte, 0, n)
    stack := make([]byte, n)
    idx := -1
    cur := byte('a')
    for _, ch := range chs {
        idx++
        stack[idx] = ch
        cnts[ch]--
        for cur < 'z' && cnts[cur] == 0 {
            cur++
        }
        for idx >= 0 && stack[idx] <= cur {
            ans = append(ans, stack[idx])
            idx--
        }
    }
    return string(ans)
}