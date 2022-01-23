func longestStrChain(words []string) int {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    Map := make(map[string]int)
    cnt := 0
    for _, word := range words {
        if Map[word] > 0 {
            continue
        }
        Map[word] = 1
        for i := 0; i < len(word); i++ {
            w := word[:i] + word[i+1:]
            if Map[w] > 0 {
                Map[word] = max(Map[word], Map[w]+1)
            }
        }
        cnt = max(cnt, Map[word])
    }
    return cnt
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}