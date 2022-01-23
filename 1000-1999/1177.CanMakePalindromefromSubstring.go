func canMakePaliQueries(s string, queries [][]int) []bool {
    cnt := make([][26]int, len(s)+1)
    for i, ch := range s {
        for j := 0; j < 26; j++ {
            cnt[i+1][j] = cnt[i][j]
        }
        cnt[i+1][ch-'a']++
    }
    n := len(queries)
    ans := make([]bool, n)
    for i, q := range queries {
        sum := 0
        for j := 0; j < 26; j++ {
            sum += (cnt[q[1]+1][j] - cnt[q[0]][j]) & 1
        }
        ans[i] = sum/2 <= q[2]
    }
    return ans
}
