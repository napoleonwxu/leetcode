func triangleNumber(nums []int) int {
    n := len(nums)
    cnt := 0
    // O(n^2), cool
    sort.Ints(nums)
    for k := 2; k < n; k++ {
        for i, j := 0, k-1; i < j; {
            if nums[k] < nums[i] + nums[j] {
                cnt += j - i
                j--
            } else {
                i++
            }
        }
    }
    /* O(n^3)
    for i := 0; i < n-2; i++ {
        for j := i+1; j < n-1; j++ {
            for k := j+1; k < n; k++ {
                if nums[i] < nums[j]+nums[k] && nums[j] < nums[i]+nums[k] && nums[k] < nums[i]+nums[j] {
                    cnt++
                }
            }
        }
    }
    */
    return cnt
}