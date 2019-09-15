const MOD = 1e9 + 7

func kConcatenationMaxSum(arr []int, k int) int {
    n := len(arr)
    maxSub, cur := math.MinInt32, 0
    for _, num := range arr {
        cur = max(cur+num, num)
        maxSub = max(maxSub, cur)
    }
    if maxSub <= 0 {
        return 0
    }
    preMax, sum := arr[0], arr[0]
    for i := 1; i < n; i++ {
        sum += arr[i]
        preMax = max(preMax, sum)
    }
    sufMax, suf := arr[n-1], arr[n-1]
    for i := n-2; i >= 0; i-- {
        suf += arr[i]
        sufMax = max(sufMax, suf)
    }
    if sum > 0 {
        return max((sufMax + sum*(k-2)%MOD + preMax)%MOD, maxSub)
    }
    return max((sufMax + preMax)%MOD, maxSub)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}