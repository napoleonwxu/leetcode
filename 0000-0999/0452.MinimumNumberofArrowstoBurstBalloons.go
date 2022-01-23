type Intervals [][]int

func (interval Intervals) Len() int{
    return len(interval)
}

func (interval Intervals) Less(i, j int) bool {
    return interval[i][1] < interval[j][1]
}

func (interval Intervals) Swap(i, j int) {
    interval[i], interval[j] = interval[j], interval[i]
}

func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    sort.Sort(Intervals(points))
    cnt, end := 1, points[0][1]
    for i := 1; i < len(points); i++ {
        if end < points[i][0] {
            cnt++
            end = points[i][1]
        }
    }
    return cnt
}