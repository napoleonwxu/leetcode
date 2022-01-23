const MOD = 1e9 + 7

func countVowelPermutation(n int) int {
    cnt := [5]int{}
    for i := 0; i < 5; i++ {
        cnt[i] = 1
    }
    for k := 1; k < n; k++ {
        nxt := [5]int{}
        nxt[0] = (cnt[1]+cnt[2]+cnt[4]) % MOD
        nxt[1] = (cnt[0]+cnt[2]) % MOD
        nxt[2] = (cnt[1]+cnt[3]) % MOD
        nxt[3] = cnt[2] % MOD
        nxt[4] = (cnt[2]+cnt[3]) % MOD
        cnt = nxt
    }
    sum := 0
    for _, c := range cnt {
        sum = (sum+c) % MOD
    }
    return sum
}