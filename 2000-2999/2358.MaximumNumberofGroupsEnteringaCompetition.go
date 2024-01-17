func maximumGroups(grades []int) int {
    // O(sqrt(n))
    sum, cnt := 0, 0
    for sum+cnt < len(grades) {
        cnt++
        sum += cnt
    }
    return cnt
}