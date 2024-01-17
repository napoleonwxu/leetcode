func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
    ans := int64(0)
    cnt := make(map[int]int)
    cnt[0] = 1
    prefix := 0
    for _, num := range nums {
        if num%modulo == k {
            prefix = (prefix + 1) % modulo
        }
        ans += int64(cnt[(prefix-k+modulo)%modulo])
        cnt[prefix]++
    }
    return ans
}