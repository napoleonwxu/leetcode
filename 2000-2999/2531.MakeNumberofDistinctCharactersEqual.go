func isItPossible(word1 string, word2 string) bool {
    cnt1, cnt2 := make(map[byte]int, 26), make(map[byte]int, 26)
    for _, ch := range []byte(word1) {
        cnt1[ch]++
    }
    for _, ch := range []byte(word2) {
        cnt2[ch]++
    }
    for ch1 := byte('a'); ch1 <= 'z'; ch1++ {
        for ch2 := byte('a'); ch2 <= 'z'; ch2++ {
            if cnt1[ch1] == 0 || cnt2[ch2] == 0 {
                continue
            }
            removeAndInsert(cnt1, ch1, ch2)
            removeAndInsert(cnt2, ch2, ch1)
            if len(cnt1) == len(cnt2) {
                return true
            }
            removeAndInsert(cnt1, ch2, ch1)
            removeAndInsert(cnt2, ch1, ch2)
        }
    }
    return false
}

func removeAndInsert(m map[byte]int, remove, insert byte) {
    m[remove]--
    if m[remove] == 0 {
        delete(m, remove)
    }
    m[insert]++
}