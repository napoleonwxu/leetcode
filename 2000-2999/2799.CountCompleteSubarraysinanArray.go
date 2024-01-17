func countCompleteSubarrays(nums []int) int {
    n, ans := len(nums), 0
    cnt := make(map[int]int, n)
    for _, num := range nums {
        cnt[num]++
    }
    k, l := len(cnt), 0
    cnt = make(map[int]int, n)
    for r := 0; r < n; r++ {
        if cnt[nums[r]] == 0 {
            k--
        }
        cnt[nums[r]]++
        for ; k == 0; l++ {
            cnt[nums[l]]--
            if cnt[nums[l]] == 0 {
                k++
            }
        }
        ans += l
    }
    return ans
}