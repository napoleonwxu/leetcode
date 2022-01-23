func findWords(words []string) []string {
    key := make(map[byte]int)
    lines := []string{"QWERTYUIOP", "ASDFGHJKL", "ZXCVBNM"}
    for i, line := range lines {
        for j := 0; j < len(line); j++ {
            key[line[j]] = i
        }
    }
    ans := []string{}
    for _, word := range words {
        upper := strings.ToUpper(word)
        flag := true
        for i := 1; i < len(upper); i++ {
            if key[upper[i]] != key[upper[0]] {
                flag = false
                break
            }
        }
        if flag {
            ans = append(ans, word)
        }
    }
    return ans
}