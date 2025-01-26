func countValidSelections(nums []int) int {
    n := len(nums)
    prefix := make([]int, n)
    prefix[0] = nums[0]
    for i := 1; i < n; i++ {
        prefix[i] = prefix[i-1] + nums[i]
    }
    sum := 0
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            if prefix[n-1] == 2*prefix[i] {
                sum += 2
            } else if abs(prefix[n-1]-2*prefix[i]) == 1 {
                sum += 1
            }
        }
    }
    return sum
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
