func countElements(nums []int) int {
    mi, ma := math.MaxInt32, math.MinInt32
    for _, num := range nums {
        mi = min(mi, num)
        ma = max(ma, num)
    }
    cnt := 0
    for _, num := range nums {
        if num > mi && num < ma {
            cnt++
        }
    }
    return cnt
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}