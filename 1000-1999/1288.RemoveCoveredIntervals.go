func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    // O(nlogn) + O(1)
    cnt, right := 0, 0
    for _, interval := range intervals {
        if interval[1] > right {
            cnt++
            right = interval[1]
        }
    }
    return cnt
    /* O(nlogn) + O(n)
    ans := [][]int{}
    for _, interval := range intervals {
        n := len(ans)
        if n > 0 && ans[n-1][1] >= interval[1] {
            continue
        }
        ans = append(ans, interval)
    }
    return len(ans)
    */
}