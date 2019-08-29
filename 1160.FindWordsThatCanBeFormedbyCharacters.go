func countCharacters(words []string, chars string) int {
    cnt := make([]int, 26)
    for _, ch := range chars {
        cnt[ch-'a']++
    }
    ans := 0
    for _, word := range words {
        flag := true
        cnt1 := make([]int, 26)
        for _, ch := range word {
            cnt1[ch-'a']++
            if cnt1[ch-'a'] > cnt[ch-'a'] {
                flag = false
                break
            }
        }
        if flag {
            ans += len(word)
        }
    }
    return ans
}