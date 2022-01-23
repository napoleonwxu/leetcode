const MOD = 1e9 + 7

func knightDialer(N int) int {
    hop := [][]int{{4,6},{6,8},{7,9},{4,8},{0,3,9},{},{0,1,7},{2,6},{1,3},{2,4}}
    now := make([]int, 10)
    for j := 0; j <= 9; j++ {
        now[j] = 1
    }
    for i := 2; i <= N; i++ {
        nxt := make([]int, 10)
        for j := 0; j <= 9; j++ {
            for _, n := range hop[j] {
                nxt[j] = (nxt[j] + now[n]) % MOD
            }
        }
        now = nxt
    }
    cnt := 0
    for j := 0; j <= 9; j++ {
        cnt = (cnt + now[j]) % MOD
    }
    return cnt
}