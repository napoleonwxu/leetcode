func countCompleteDayPairs(hours []int) int64 {
    ans := int64(0)
    cnts := make([]int, 24)
    for _, h := range hours {
        ans += int64(cnts[(24-h%24)%24])
        cnts[h%24]++
    }
    return ans
}