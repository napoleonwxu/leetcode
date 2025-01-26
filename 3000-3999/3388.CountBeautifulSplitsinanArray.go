const MOD = int64(1e9 + 7)

func beautifulSplits(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    base := int64(31)
    preHash := make([]int64, n+1)
    pow := make([]int64, n+1)
    pow[0] = 1
    for i := 0; i < n; i++ {
        pow[i+1] = (pow[i] * base) % MOD
        preHash[i+1] = (preHash[i] * base + int64(nums[i])) % MOD
    }
    cnt := 0
    for i := 1; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            if isPrefix(preHash, pow, 0, i, j) || isPrefix(preHash, pow, i, j, n) {
                cnt++
            }
        }
    }
    return cnt
}

func isPrefix(preHash, pow []int64, i, j, k int) bool {
    l1, l2 := j-i, k-j
    if l1 > l2 {
        return false
    }
    hash1 := (preHash[j] - (preHash[i]*pow[l1])%MOD + MOD) % MOD
    hash2 := (preHash[j+l1] - (preHash[j]*pow[l1])%MOD + MOD) % MOD
    return hash1 == hash2
}