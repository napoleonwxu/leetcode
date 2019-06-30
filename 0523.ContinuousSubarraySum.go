func checkSubarraySum(nums []int, k int) bool {
    n := len(nums)
    // Since the size of subarray is at least 2.
    if n < 2 {
        return false
    }
    // Two continuous "0" will form a subarray which has sum = 0. 0 * k == 0 will always be true.
    for i := 0; i < n-1; i++ {
        if nums[i] == 0 && nums[i+1] == 0 {
            return true
        }
    }
    // At this point, k can't be "0" any longer.
    if k == 0 {
        return false
    }
    
    Map := make(map[int]int)
    Map[0] = -1
    s := 0
    for i := 0; i < n; i++ {
        s = (s + nums[i])%k
        j, ok := Map[s]
        if ok {
            if i-j >= 2 {
                return true
            }
        } else {
            Map[s] = i
        }
    }
    return false
}