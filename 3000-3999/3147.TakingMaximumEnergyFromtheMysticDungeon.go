func maximumEnergy(energy []int, k int) int {
    ans, n := math.MinInt32, len(energy)
    for i := n-1; i >= n-k; i-- {
        tmp := 0
        for j := i; j >= 0; j -= k {
            tmp += energy[j]
            ans = max(ans, tmp)
        }
    }
    return ans
}