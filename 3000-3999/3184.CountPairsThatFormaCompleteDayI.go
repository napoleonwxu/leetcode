func countCompleteDayPairs(hours []int) int {
    cnt, n := 0, len(hours)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if (hours[i]+hours[j]) % 24 == 0 {
                cnt++
            }
        }
    }
    return cnt
}