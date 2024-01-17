import "sort"

func maxScore(nums []int) int {
    sort.Ints(nums)
    cnt := 0
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] <= 0 {
            break
        }
        cnt++
        if i > 0 {
            nums[i-1] += nums[i]
        }
    }
    return cnt
}