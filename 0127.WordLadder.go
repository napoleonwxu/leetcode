func ladderLength(beginWord string, endWord string, wordList []string) int {
    Map := make(map[string]bool, len(wordList))
    for _, word := range wordList {
        Map[word] = true
    }
    queue := []string{beginWord}
    step := 1
    for len(queue) > 0 {
        nxt := []string{}
        for _, cur := range queue {
            if cur == endWord {
                return step
            }
            for i := 0; i < len(cur); i++ {
                for ch := 'a'; ch <= 'z'; ch++ {
                    tmp := cur[:i] + string(ch) + cur[i+1:]
                    if Map[tmp] {
                        nxt = append(nxt, tmp)
                        delete(Map, tmp)
                    }
                }
            }
        }
        step++
        queue = nxt
    }
    return 0
}