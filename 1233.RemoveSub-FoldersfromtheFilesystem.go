func removeSubfolders(folder []string) []string {
    sort.Strings(folder)
    Map := make(map[string]bool)
    ans := []string{}
    for _, f := range folder {
        flag := true
        for i := 1; i < len(f); i++ {
            if f[i] == '/' && Map[f[:i]] {
                flag = false
                break
            }
        }
        if flag {
            Map[f] = true
            ans = append(ans, f)
        }
    }
    return ans
}