func rob(nums []int) int {
    pre, cur := 0, 0
    for _, num := range nums {
        pre, cur = cur, max(pre+num, cur)
    }
    return cur
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}