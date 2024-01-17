func minIncrementOperations(nums []int, k int) int64 {
    var dp1, dp2, dp3, dp int64
    for _, num := range nums {
        dp = min(dp1, min(dp2, dp3)) + max(int64(k-num), 0)
        dp1 = dp2
        dp2 = dp3
        dp3 = dp
    }
    return min(dp1, min(dp2, dp3))
}

func min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}

func max(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}