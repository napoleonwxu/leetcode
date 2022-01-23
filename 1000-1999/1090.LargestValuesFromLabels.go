func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
    n := len(values)
    vls := make([][]int, n)
    for i := 0; i < n; i++ {
        vls[i] = []int{values[i], labels[i]}
    }
    sort.Slice(vls, func(i, j int) bool {
        return vls[i][0] > vls[j][0]
    })
    sum, cnt := 0, 0
    Map := make(map[int]int)
    for _, vl := range vls {
        if cnt >= num_wanted {
            break
        }
        if Map[vl[1]] < use_limit {
            sum += vl[0]
            cnt++
            Map[vl[1]]++
        }
    }
    return sum
}