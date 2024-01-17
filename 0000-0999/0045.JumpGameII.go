func jump(nums []int) int {
    // O(n)
    curFar, curEnd := 0, 0
    jumps := 0
    for i := 0; i < len(nums)-1; i++ {
        curFar = max(curFar, i+nums[i])
        if i == curEnd {
            jumps++
            curEnd = curFar
        }
    }
    return jumps
    /* O(n^2)
    n := len(nums)
    dp := make([]int, n)
    dp[0] = 0
    for i := 1; i < n; i++ {
        dp[i] = n
    }
    for i, num := range nums {
        for j := i+1; j <= i+num && j < n; j++ {
            dp[j] = min(dp[j], dp[i]+1)
        }
    }
    return dp[n-1]
    */
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}