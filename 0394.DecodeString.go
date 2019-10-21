func decodeString(s string) string {
    words, cnts := []string{}, []int{}
    word, cnt := "", 0
    for _, ch := range s {
        if ch >= '0' && ch <= '9' {
            cnt = 10*cnt + int(ch-'0')
        } else if ch == '[' {
            words = append(words, word)
            cnts = append(cnts, cnt)
            word, cnt = "", 0
        } else if ch == ']' {
            n := len(words)
            word = words[n-1] + strings.Repeat(word, cnts[n-1])
            words = words[:n-1]
            cnts = cnts[:n-1]
        } else {
            word += string(ch)
        }
    }
    return word
}