func maxScoreWords(words []string, letters []byte, score []int) int {
    cnt := make([]int, len(score))
    for _, ch := range letters {
        cnt[ch-'a']++
    }
    return dfs(words, cnt, score, 0)
}

func dfs(words []string, cnt, score []int, idx int) int {
    ans := 0
    for i := idx; i < len(words); i++ {
        cur, valid := 0, true
        for j := 0; j < len(words[i]); j++ {
            cnt[words[i][j]-'a']--
            if cnt[words[i][j]-'a'] < 0 {
                valid = false
            }
            cur += score[words[i][j]-'a']
        }
        if valid {
            cur += dfs(words, cnt, score, i+1)
            ans = max(ans, cur)
        }
        for j := 0; j < len(words[i]); j++ {
            cnt[words[i][j]-'a']++
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}