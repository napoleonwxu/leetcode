func minimumSumSubarray(nums []int, l int, r int) int {
    n, ans := len(nums), math.MaxInt64
    for i := 0; i < n; i++ {
        j, sum := i, 0
        for ; j < n && j-i+1 < l; j++ {
            sum += nums[j]
        }
        for ; j < n && j-i+1 <= r; j++ {
            sum += nums[j]
            if sum > 0 {
                ans = min(ans, sum)
            }
        }
    }
    if ans == math.MaxInt64 {
        return -1
    }
    return ans
}