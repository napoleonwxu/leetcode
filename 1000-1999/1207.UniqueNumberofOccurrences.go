func uniqueOccurrences(arr []int) bool {
    cnt := make(map[int]int)
    for _, n := range arr {
        cnt[n]++
    }
    exist := make(map[int]bool)
    for _, c := range cnt {
        if exist[c] {
            return false
        }
        exist[c] = true
    }
    return true
}