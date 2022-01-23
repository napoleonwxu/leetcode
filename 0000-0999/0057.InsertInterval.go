func insert(intervals [][]int, newInterval []int) [][]int {
    ans := [][]int{}
    for i, interval := range intervals {
        if interval[1] < newInterval[0] {
            ans = append(ans, interval)
        } else if newInterval[1] < interval[0] {
            ans = append(ans, newInterval)
            ans = append(ans, intervals[i:]...)
            break
        } else {
            newInterval[0] = min(newInterval[0], interval[0])
            newInterval[1] = max(newInterval[1], interval[1])
        }
    }
    if len(ans) == 0 || ans[len(ans)-1][1] < newInterval[0] {
        ans = append(ans, newInterval)
    }
    return ans
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