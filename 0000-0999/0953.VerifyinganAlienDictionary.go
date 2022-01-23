func compare(w1 string, w2 string, index []int) bool {
    n1, n2 := len(w1), len(w2)
    i := 0
    for i < n1 && i < n2 {
        if index[w1[i]-'a'] != index[w2[i]-'a'] {
            return index[w1[i]-'a'] < index[w2[i]-'a']
        }
        i++
    }
    return n1 <= n2
}

func isAlienSorted(words []string, order string) bool {
    index := make([]int, 26)
    for i, c := range order {
        index[c-'a'] = i
    }
    for i := 1; i < len(words); i++ {
        if !compare(words[i-1], words[i], index) {
            return false
        }
    }
    return true
}