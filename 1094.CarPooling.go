//Same with 253. Meeting Rooms II.
func carPooling(trips [][]int, capacity int) bool {
    // O(n)
    stops := [1001]int{}
    for _, trip := range trips {
        stops[trip[1]] += trip[0]
        stops[trip[2]] -= trip[0]
    }
    for _, p := range stops {
        capacity -= p
        if capacity < 0 {
            return false
        }
    }
    /* O(nlgn) + O(n*n)
    sort.Slice(trips, func(i, j int) bool {
        return trips[i][1] < trips[j][1]
    })
    cnts := [1001]int{}
    for _, trip := range trips {
        cnt, from, to := trip[0], trip[1], trip[2]
        for loc := from; loc < to; loc++ {
            cnts[loc] += cnt
        }
        if cnts[from] > capacity {
            return false
        }
    }
    */
    return true
}