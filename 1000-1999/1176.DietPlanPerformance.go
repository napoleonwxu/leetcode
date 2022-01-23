func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
    n := len(calories)
    point, T := 0, 0
    i := 0
    for ; i < k; i++ {
        T += calories[i]
    }
    if T < lower {
        point--
    } else if T > upper {
        point++
    }
    for ; i < n; i++ {
        T += calories[i] - calories[i-k]
        if T < lower {
            point--
        } else if T > upper {
            point++
        }
    }
    return point
}