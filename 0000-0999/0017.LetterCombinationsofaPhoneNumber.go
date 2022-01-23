func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    key := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    // solution2
    ans := []string{""}
    for _, d := range digits {
        nxt := []string{}
        for _, path := range ans {
            for _, ch := range key[int(d-'0')] {
                nxt = append(nxt, path+string(ch))
            }
        }
        ans = nxt
    }
    /* solution1
    ans := []string{}
    dfs(digits, "", key, &ans)
    */
    return ans
}

func dfs(digits, path string, key []string, ans *[]string) {
    if len(digits) == 0 {
        *ans = append(*ans, path)
        return
    }
    for _, ch := range key[int(digits[0]-'0')] {
        dfs(digits[1:], path+string(ch), key, ans)
    }
}