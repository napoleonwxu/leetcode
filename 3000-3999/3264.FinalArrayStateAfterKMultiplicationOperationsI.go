func getFinalState(nums []int, k int, multiplier int) []int {
    for ; k > 0; k-- {
        idx := 0
        for i := range nums {
            if nums[i] < nums[idx] {
                idx = i
            }
        }
        nums[idx] *= multiplier
    }
    return nums
}