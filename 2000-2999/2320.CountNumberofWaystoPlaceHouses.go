const MOD = 1e9 + 7

func countHousePlacements(n int) int {
    pre, cur := 1, 2
    for ; n > 1; n-- {
        pre, cur = cur, (pre+cur)%MOD
    }
    return (cur * cur) % MOD
}