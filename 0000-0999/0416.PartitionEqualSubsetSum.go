func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum&1 == 1 {
        return false
    }
    half := sum >> 1
    dp := make([]bool, half+1)
    dp[0] = true
    for i := 0; i < len(nums); i++ {
        for j := half; j >= nums[i]; j-- {
            dp[j] = dp[j] || dp[j-nums[i]]
        }
    }
    return dp[half]
}