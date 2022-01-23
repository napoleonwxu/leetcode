func rankTeams(votes []string) string {
    n := len(votes[0])
    Map := make(map[byte][]int)
    for _, vote := range votes {
        for i := range vote {
            if _, ok := Map[vote[i]]; !ok {
                Map[vote[i]] = make([]int, n)
            }
            Map[vote[i]][i]++
        }
    }
    ans := make([]byte, n)
    for i := range votes[0] {
        ans[i] = votes[0][i]
    }
    sort.Slice(ans, func(i, j int) bool {
        for k := 0; k < n; k++ {
            if Map[ans[i]][k] != Map[ans[j]][k] {
                return Map[ans[i]][k] > Map[ans[j]][k]
            }
        }
        return ans[i] < ans[j]
    })
    return string(ans)
}