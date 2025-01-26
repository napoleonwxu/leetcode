func maximumCoins(coins [][]int, k int) int64 {
    sort.Slice(coins, func(i, j int) bool {
        return coins[i][0] < coins[j][0]
    })
    n := len(coins)
    ans, cur := int64(0), int64(0)
    for i, j := 0, 0; i < n; i++ {
        for j < n && coins[j][1] <= coins[i][0] + k - 1 {
            cur += int64(coins[j][1] - coins[j][0] + 1) * int64(coins[j][2])
            j++
        }
        if j < n {
            part := int64(max(0, coins[i][0] + k - 1 - coins[j][0] + 1)) * int64(coins[j][2])
            ans = max(ans, cur + part)
        }
        cur -= int64(coins[i][1] - coins[i][0] + 1) * int64(coins[i][2])
    }
    cur = int64(0)
    for i, j := 0, 0; i < n; i++ {
        cur += int64(coins[i][1] - coins[i][0] + 1) * int64(coins[i][2])
        for coins[j][1] < coins[i][1] - k + 1 {
            cur -= int64(coins[j][1] - coins[j][0] +1) * int64(coins[j][2])
            j++
        }
        part := int64(max(0, coins[i][1] - k - coins[j][0] + 1)) * int64(coins[j][2])
        ans = max(ans, cur - part)
    }
    return ans
}