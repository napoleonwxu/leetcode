func countDays(days int, meetings [][]int) int {
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][0] < meetings[j][0]
    })
    merge := make([][]int, 0, len(meetings))
    for _, m := range meetings {
        n := len(merge)
        if n > 0 && m[0] <= merge[n-1][1] {
            merge[n-1][1] = max(merge[n-1][1], m[1])
        } else {
            merge = append(merge, m)
        }
    }
    sum := 0
    for _, m := range merge {
        sum += m[1] - m[0] + 1
    }
    return days - sum
}