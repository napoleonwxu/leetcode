func countPartitions(nums []int) int {
    sum, left := 0, 0
    for _, num := range nums {
        sum += num
    }
    cnt := 0
    for i := 0; i < len(nums)-1; i++ {
        left += nums[i]
        if left%2 == (sum-left)%2 {
            cnt++
        }
    }
    return cnt
}
