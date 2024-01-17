import "sort"

func minOperations(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum < target {
        return -1
    }
    sort.Ints(nums)
    cnt := 0
    for target > 0 {
        n := len(nums)
        max := nums[n-1]
        nums = nums[:n-1]
        if sum-max >= target {
            sum -= max
        } else if target >= max {
            sum -= max
            target -= max
        } else {
            nums = append(nums, max/2)
            nums = append(nums, max/2)
            cnt++
        }
    }
    return cnt
}