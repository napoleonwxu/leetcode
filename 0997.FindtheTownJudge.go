func findJudge(N int, trust [][]int) int {
    cnt := make([]int, N+1)
    for _, p := range trust {
        cnt[p[1]]++
        cnt[p[0]]--
    }
    for i := 1; i <= N; i++ {
        if cnt[i] == N-1 {
            return i
        }
    }
    return -1
}