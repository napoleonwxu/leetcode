func numSmallerByFrequency(queries []string, words []string) []int {
    f_words := make([]int, len(words))
    for i, w := range words {
        f_words[i] = f(w)
    }
    sort.Ints(f_words)
    ans := make([]int, len(queries))
    for i, query := range queries {
        f_query := f(query)
        idx := sort.Search(len(f_words), func(i int) bool {
            return f_words[i] > f_query
        })
        ans[i] = len(f_words) - idx
    }
    return ans
}

func f(word string) int {
    cnt := make([]int, 26)
    for _, ch := range word {
        cnt[ch-'a']++
    }
    for _, c := range cnt {
        if c > 0 {
            return c
        }
    }
    return 0
}