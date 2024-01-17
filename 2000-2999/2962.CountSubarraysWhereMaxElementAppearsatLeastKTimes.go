func countSubarrays(nums []int, k int) int64 {
    max := 0
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    cnt, ans := 0, int64(0)
    for l, r := 0, 0; r < len(nums); r++ {
        if nums[r] == max {
            cnt++
        }
        for ; cnt >= k; l++ {
            if nums[l] == max {
                cnt--
            }
        }
        ans += int64(l)
    }
    return ans
}