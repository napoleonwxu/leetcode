func longestMonotonicSubarray(nums []int) int {
    ans := 1
    inc, dec := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            inc++
            dec = 1
        } else if nums[i] < nums[i-1] {
            dec++
            inc = 1
        } else {
            inc, dec = 1, 1
        }
        ans = max(ans, max(inc, dec))
    }
    return ans
}