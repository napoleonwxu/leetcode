type Intervals [][]int

func (intervals Intervals) Len() int {
    return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
}

func (intervals Intervals) Swap(i, j int) {
    intervals[i], intervals[j] = intervals[j], intervals[i]
}

func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    sort.Sort(Intervals(intervals))
    var ans [][]int
    tmp := intervals[0]
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= tmp[1] {
            if intervals[i][1] > tmp[1] {
                tmp[1] = intervals[i][1]
            }
        } else {
            ans = append(ans, tmp)
            tmp = intervals[i]
        }
    }
    ans = append(ans, tmp)
    return ans
}