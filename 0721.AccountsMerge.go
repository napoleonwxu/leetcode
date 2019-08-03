func accountsMerge(accounts [][]string) [][]string {
    Map := make(map[string][]int)
    for num, account := range accounts {
        for i := 1; i < len(account); i++ {
            email := account[i]
            Map[email] = append(Map[email], num)
        }
    }
    visited := make([]bool, len(accounts))
    ans :=  [][]string{}
    var dfs func(int, map[string]bool)
    dfs = func(num int, emails_map map[string]bool) {
        if visited[num] {
            return
        }
        visited[num] = true
        account := accounts[num]
        for i := 1; i < len(account); i++ {
            emails_map[account[i]] = true
            for _, nei := range Map[account[i]] {
                dfs(nei, emails_map)
            }
        }
    }
    for num, account := range accounts {
        if visited[num] {
            continue
        }
        emails_map := make(map[string]bool)
        dfs(num, emails_map)
        emails := []string{}
        for email, _ := range emails_map {
            emails = append(emails, email)
        }
        sort.Strings(emails)
        new_account := []string{account[0]}
        new_account = append(new_account, emails...)
        ans = append(ans, new_account)
    }
    return ans
}
