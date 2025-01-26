func minimumOperations(nums []int) int {
    n := len(nums)
    m := make(map[int]bool, n)
    for i := n-1; i >= 0; i-- {
        if m[nums[i]] {
            return (i+3)/3
        }
        m[nums[i]] = true
    }
    return 0
}