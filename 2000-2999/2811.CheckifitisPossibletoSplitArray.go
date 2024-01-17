func canSplitArray(nums []int, m int) bool {
    if len(nums) <= 2 {
        return true
    }
    for i := 1; i < len(nums); i++ {
        if nums[i-1]+nums[i] >= m {
            return true
        }
    }
    return false
}
