func maximumTotalDamage(power []int) int64 {
    cnt := make(map[int]int, len(power))
    for _, p := range power {
        cnt[p]++
    }
    uniq := make([]int, 0, len(cnt))
    for p := range cnt {
        uniq = append(uniq, p)
    }
    sort.Ints(uniq)
    n := len(uniq)
    dp := make([]int64, n)
    dp[0] = int64(uniq[0]) * int64(cnt[uniq[0]])
    for i := 1; i < n; i++ {
        cur := int64(uniq[i]) * int64(cnt[uniq[i]])
        dp[i] = max(dp[i-1], cur)
        pre := i - 1
        for ; pre >= 0 && uniq[pre] >= uniq[i]-2; pre-- {}
        if pre >= 0 {
            dp[i] = max(dp[i], dp[pre]+cur)
        }
    }
    return dp[n-1]
}
