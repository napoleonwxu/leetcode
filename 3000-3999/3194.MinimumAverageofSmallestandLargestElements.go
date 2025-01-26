func minimumAverage(nums []int) float64 {
    sort.Ints(nums)
    n := len(nums)
    ans := float64(nums[n-1])
    for i := 0; i < n/2; i++ {
        ans = min(ans, float64(nums[i]+nums[n-1-i])/2)
    }
    return ans
}