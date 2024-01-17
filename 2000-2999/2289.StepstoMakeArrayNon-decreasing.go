func totalSteps(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    stack := make([]int, n)
    step, j := 0, -1
    for i := n - 1; i >= 0; i-- {
        for j >= 0 && nums[i] > nums[stack[j]] {
            dp[i] = max(dp[i]+1, dp[stack[j]])
            j--
        }
        step = max(step, dp[i])
        j++
        stack[j] = i
    }
    return step
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
