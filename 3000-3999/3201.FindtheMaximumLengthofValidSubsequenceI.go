func maximumLength(nums []int) int {
    odd, even := 0, 0
    same := 0
    for i := range nums {
        if nums[i]%2 == 1 {
            odd++
        } else {
            even++
        }
        if i > 0 && nums[i]%2 == nums[i-1]%2 {
            same++
        }
    }
    return max(max(odd, even), len(nums)-same)
}