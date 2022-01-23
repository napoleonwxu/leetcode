func findMaxAverage(nums []int, k int) float64 {
    // O(n) + O(1)
    sum, i := 0, 0
    for ; i < k; i++ {
        sum += nums[i]
    }
    max := sum
    for ; i < len(nums); i++ {
        sum += nums[i] - nums[i-k]
        if sum > max {
            max = sum
        }
    }
    return float64(max)/float64(k)
}