import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    n := len(capacity)
    sub := make([]int, n)
    for i := 0; i < n; i++ {
        sub[i] = capacity[i] - rocks[i]
    }
    sort.Ints(sub)
    i := 0
    for ; i < n && additionalRocks >= sub[i]; i++ {
        additionalRocks -= sub[i]
    }
    return i
}