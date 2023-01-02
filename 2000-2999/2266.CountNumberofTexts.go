const MOD = 1e9+7

func countTexts(pressedKeys string) int {
    n := len(pressedKeys)
    dp := make([]int, n+1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        cnt := 3
        if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
            cnt = 4
        }
        dp[i] = dp[i-1]
        for j := i-1; j > i-cnt && j > 0 && pressedKeys[i-1] == pressedKeys[j-1]; j-- {
            dp[i] = (dp[i]+dp[j-1]) % MOD
        }
    }
    return dp[n]
}