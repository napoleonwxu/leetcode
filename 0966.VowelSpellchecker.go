func spellchecker(wordlist []string, queries []string) []string {
    n := len(wordlist)
    Map1 := make(map[string]bool, n)
    Map2 := make(map[string]int, n)
    Map3 := make(map[string]int, n)
    for i := n-1; i >= 0; i-- {
        Map1[wordlist[i]] = true
        lower := strings.ToLower(wordlist[i])
        Map2[lower] = i
        vowel := dealVowel(lower)
        Map3[vowel] = i
    }
    ans := make([]string, len(queries))
    for i, query := range queries {
        if Map1[query] {
            ans[i] = query
            continue
        }
        lower := strings.ToLower(query)
        if idx, ok := Map2[lower]; ok {
            ans[i] = wordlist[idx]
        } else if idx, ok := Map3[dealVowel(lower)]; ok {
            ans[i] = wordlist[idx]
        } else {
            ans[i] = ""
        }
    }
    return ans
}

func dealVowel(lower string) string {
    ans := []rune(lower)
    for i, ch := range lower {
        if ch == 'e' || ch == 'i' || ch =='o' || ch =='u' {
            ans[i] = 'a'
        }
    }
    return string(ans)
}