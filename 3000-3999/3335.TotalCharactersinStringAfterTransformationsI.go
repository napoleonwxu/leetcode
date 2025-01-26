const MOD = int(1e9 + 7)

func lengthAfterTransformations(s string, t int) int {
    cnt := make([]int, 26)
    for _, ch := range s {
        cnt[ch-'a']++
    }
    for ; t > 0; t-- {
        tmp := make([]int, 26)
        for i := 0; i < 26; i++ {
            if i == 25 {
                tmp[0] = (tmp[0] + cnt[i]) % MOD
                tmp[1] = (tmp[1] + cnt[i]) % MOD
            } else {
                tmp[i+1] = (tmp[i+1] + cnt[i]) % MOD
            }
        }
        cnt = tmp
    }
    ans := 0
    for _, c := range cnt {
        ans = (ans + c) % MOD
    }
    return ans
}