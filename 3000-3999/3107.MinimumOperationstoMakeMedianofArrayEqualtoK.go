func minOperationsToMakeMedianK(nums []int, k int) int64 {
    sort.Ints(nums)
    ans, n := int64(0), len(nums)
    for i := 0; i <= n/2; i++ {
        ans += int64(max(0, nums[i]-k))
    }
    for i := n/2; i < n; i++ {
        ans += int64(max(0, k-nums[i]))
    }
    return ans
}