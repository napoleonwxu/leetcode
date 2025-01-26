func maxDistinctElements(nums []int, k int) int {
    sort.Ints(nums)
    cnt, pre := 0, math.MinInt32
    for _, num := range nums {
        mn, mx := num-k, num+k
        if pre < mn {
            pre = mn
            cnt++
        } else if pre < mx {
            pre++
            cnt++
        }
    }
    return cnt
}