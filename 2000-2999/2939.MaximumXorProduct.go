const MOD = int64(1e9 + 7)

func maximumXorProduct(a int64, b int64, n int) int {
    if n > 0 {
        for bit := int64(1 << (n - 1)); bit > 0; bit >>= 1 {
            if min(a, b)&bit == 0 {
                a ^= bit
                b ^= bit
            }
        }
    }
    return int((a % MOD) * (b % MOD) % MOD)
}

func min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}