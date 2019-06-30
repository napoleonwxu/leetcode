func findLongestChain(pairs [][]int) int {
    /* DP, O(n*n)
    n := len(pairs)
    sort.Slice(pairs, func (i, j int) bool {
        return pairs[i][0] < pairs[j][0]
    })
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    ans := 1
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            if pairs[j][1] < pairs[i][0] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        ans = max(ans, dp[i])
    }
    return ans
    */
    // Greedy, O(nlogn)
    sort.Slice(pairs, func (i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
    cnt, cur := 0, -1<<31
    for _, pair := range pairs {
        if cur < pair[0] {
            cur = pair[1]
            cnt++
        }
    }
    return cnt
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}