func maxSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    preSum := make([]int64, n+1)
    for i := range nums {
        preSum[i+1] = preSum[i] + int64(nums[i])
    }
    ans := preSum[k]
    for i := 0; i < k; i++ {
        sum := int64(0)
        for j := i; j+k <= n; j += k {
            cur := preSum[j+k] - preSum[j]
            sum = max(sum+cur, cur)
            ans = max(ans, sum)
        }
    }
    return ans
}