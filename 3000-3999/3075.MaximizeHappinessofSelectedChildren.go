func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Ints(happiness)
    turn := 0
    ans := int64(0)
    for i := len(happiness)-1; i >= 0 && happiness[i]-turn > 0 && k > 0; i-- {
        ans += int64(happiness[i] - turn)
        turn++
        k--
    }
    return ans
}