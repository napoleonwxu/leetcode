func restoreIpAddresses(s string) []string {
    if len(s) < 4 || len(s) > 12 {
        return nil
    }
    ans := []string{}
    dfs(s, 0, 0, []string{}, &ans)
    return ans
}

func dfs(s string, i, depth int, path[]string, ans *[]string) {
    if depth == 4 {
        if i >= len(s) {
            *ans = append(*ans, strings.Join(path, "."))
        }
        return
    }
    for j := i+1; j <= i+3 && j <= len(s); j++ {
        ip, _ := strconv.Atoi(s[i:j])
        if ip > 255 || (s[i] == '0' && j > i+1) {
            return
        }
        dfs(s, j, depth+1, append(path, s[i:j]), ans)
    }
}