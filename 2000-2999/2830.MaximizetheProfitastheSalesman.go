func maximizeTheProfit(n int, offers [][]int) int {
    dp := make([]int, n+1)
    m := make(map[int][][]int, n)
    for _, offer := range offers {
        m[offer[1]] = append(m[offer[1]], []int{offer[0], offer[2]})
    }
    for end := 1; end <= n; end++ {
        dp[end] = dp[end-1]
        for _, tmp := range m[end-1] {
            dp[end] = max(dp[end], dp[tmp[0]]+tmp[1])
        }
    }
    return dp[n]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}