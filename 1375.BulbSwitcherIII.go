func numTimesAllBlue(light []int) int {
    cnt, right := 0, 0
    // O(n) + O(1)
    for i, bulb := range light {
        right = max(right, bulb)
        if right == i+1 {
            cnt++
        }
    }
    /* O(n) + O(n)
    Map := make(map[int]bool, len(light))
    for _, bulb := range light {
        right = max(right, bulb)
        Map[bulb] = true
        if len(Map) == right {
            cnt++
        }
    } */
    return cnt
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}