func removeDuplicates(nums []int) int {
    i := 0
    for j := range nums {
        if i < 2 || nums[j] > nums[i-2] {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}