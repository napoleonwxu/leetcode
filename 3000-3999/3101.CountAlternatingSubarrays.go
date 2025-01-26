func countAlternatingSubarrays(nums []int) int64 {
    ans, cur := int64(0), int64(0)
    for i := range nums {
        if i == 0 || nums[i] != nums[i-1] {
            cur++
        } else {
            cur = 1
        }
        ans += cur
    }
    return ans
}