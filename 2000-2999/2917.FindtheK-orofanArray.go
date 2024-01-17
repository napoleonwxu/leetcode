func findKOr(nums []int, k int) int {
    ans := 0
    for i := 0; i <= 31; i++ {
        cnt := 0
        for _, num := range nums {
            if (1<<i)&num > 0 {
                cnt++
            }
            if cnt >= k {
                break
            }
        }
        if cnt >= k {
            ans += 1 << i
        }
    }
    return ans
}