func numberOfSubarrays(nums []int, k int) int {
    return atMostKOdds(nums, k) - atMostKOdds(nums, k-1)
}

func atMostKOdds(nums []int, k int) int {
    cnt := 0
    i := 0
    for j := 0; j < len(nums); j++ {
        k -= nums[j] & 1
        for ; k < 0; i++ {
            k += nums[i] & 1
        }
        cnt += j - i + 1
    }
    return cnt
}