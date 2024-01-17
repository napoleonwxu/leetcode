func maxArrayValue(nums []int) int64 {
    ans := int64(0)
    for i := len(nums) - 1; i >= 0; i-- {
        if ans >= int64(nums[i]) {
            ans += int64(nums[i])
        } else {
            ans = int64(nums[i])
        }
    }
    return ans
}