const MOD = 1e9 + 7

func minimumPossibleSum(n int, target int) int {
    sum := 0
    for i := 1; n > 0; i++ {
        if 2*i <= target || i >= target {
            sum = (sum + i) % MOD
            n--
        }
    }
    return sum
}