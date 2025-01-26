func minArraySum(nums []int, k int, op1 int, op2 int) int {
    dp := make([][][]int, len(nums))
    for i := range dp {
        dp[i] = make([][]int, op1+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, op2+1)
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }
    return solve(dp, nums, 0, k, op1, op2)
}

func solve(dp [][][]int, nums []int, i, k, op1, op2 int) int {
    if i >= len(nums) {
        return 0
    }
    if dp[i][op1][op2] != -1 {
        return dp[i][op1][op2]
    }
    ans := nums[i] + solve(dp, nums, i+1, k, op1, op2)
    if op1 > 0 {
        ans = min(ans, (nums[i]+1)/2 + solve(dp, nums, i+1, k, op1-1, op2))
    }
    if op2 > 0 && nums[i] >= k {
        ans = min(ans, nums[i] - k + solve(dp, nums, i+1, k, op1, op2-1))
    }
    if op1 > 0 && op2 > 0 {
        left := (nums[i]+1) / 2
        if left >= k {
            ans = min(ans, left - k + solve(dp, nums, i+1, k, op1-1, op2-1))
        }
    }
    if op1 > 0 && op2 > 0 && nums[i] >= k {
        ans = min(ans, (nums[i]-k+1)/2 + solve(dp, nums, i+1, k, op1-1, op2-1))
    }
    dp[i][op1][op2] = ans
    return ans
}