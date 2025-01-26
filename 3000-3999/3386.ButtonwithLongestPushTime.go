func buttonWithLongestTime(events [][]int) int {
    idx, pre := events[0][0], events[0][1]
    for i := 1; i < len(events); i++ {
        cur := events[i][1] - events[i-1][1]
        if cur > pre || (cur == pre && events[i][0] < idx) {
                idx, pre = events[i][0], cur
            }
    }
    return idx
}