func partition(s string) [][]string {
    ans := [][]string{}
    dfs(s, 0, []string{}, &ans)
    return ans
}

func dfs(s string, i int, path []string, ans *[][]string) {
    if i >= len(s) {
        tmp := make([]string, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
    }
    for j := i+1; j <= len(s); j++ {
        if isPalindrome(s[i:j]) {
            dfs(s, j, append(path, s[i:j]), ans)
        }
    }
}

func isPalindrome(s string) bool {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}