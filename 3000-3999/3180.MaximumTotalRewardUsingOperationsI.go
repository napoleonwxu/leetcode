func maxTotalReward(rewardValues []int) int {
    sort.Ints(rewardValues)
    n := len(rewardValues)
    ma := rewardValues[n-1] * 2
    preDP := make([]bool, ma)
    curDP := make([]bool, ma)
    preDP[0] = true
    for _, v := range rewardValues {
        for i := 0; i < v; i++ {
            curDP[i+v] = preDP[i+v] || preDP[i]
        }
        for i := 0; i < ma; i++ {
            preDP[i] = preDP[i] || curDP[i]
            curDP[i] = false
        }
    }
    for i := ma-1; i >= 0; i-- {
        if preDP[i] {
            return i
        }
    }
    return 0
}