func isZeroArray(nums []int, queries [][]int) bool {
    n := len(nums)
    prefix := make([]int, n+1)
    for _, q := range queries {
        prefix[q[0]]++
        prefix[q[1]+1]--
    }
    for i := 0; i < n; i++ {
        if i > 0 {
            prefix[i] += prefix[i-1]
        }
        if nums[i] > prefix[i] {
            return false
        }
    }
    return true
}