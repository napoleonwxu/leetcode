type Intervals [][]int

func (intervals Intervals) Len() int {
    return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
    /*
    if intervals[i][1] == intervals[j][1] {
        return intervals[i][0] < intervals[j][0]
    }
    */
    return intervals[i][1] < intervals[j][1]
}

func (intervals Intervals) Swap(i, j int) {
    intervals[i], intervals[j] = intervals[j], intervals[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) <= 1 {
        return 0
    }
    sort.Sort(Intervals(intervals))
    cnt, end := 0, intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < end {
            cnt++
        } else {
            end = intervals[i][1]
        }
    }
    return cnt
}