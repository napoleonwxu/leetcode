func numberOfSubarrays(nums []int, k int) int {
    return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
    cnt := 0
    i := 0
    for j := 0; j < len(nums); j++ {
        k -= nums[j] & 1
        for k < 0 {
            k += nums[i] & 1
            i++
        }
        cnt += j-i+1
    }
    return cnt
}