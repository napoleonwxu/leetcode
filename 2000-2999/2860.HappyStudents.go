import "sort"

func countWays(nums []int) int {
    n, cnt := len(nums), 0
    sort.Ints(nums)
    if nums[0] > 0 {
        cnt++
    }
    for i := 0; i < n-1; i++ {
        if i+1 > nums[i] && i+1 < nums[i+1] {
            cnt++
        }
    }
    if n > nums[n-1] {
        cnt++
    }
    return cnt
}