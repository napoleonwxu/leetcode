// https://discuss.leetcode.com/topic/86696/share-my-o-n-c-dp-solution-with-thinking-process-and-explanation/3
const MOD  = 1e9 + 7

func checkRecord(n int) int {
    if n == 1 {
        return 3
    }
    if n == 2 {
        return 3*3 - 1
    }
    P := make([]int, n+1)
    L := make([]int, n+1)
    A := make([]int, n+1)
    P[1], P[2] = 1, 3
    L[1], L[2] = 1, 3
    A[1], A[2], A[3] = 1, 2, 4
    for i := 3; i <= n; i++ {
        P[i] = (P[i-1] + L[i-1] + A[i-1]) % MOD
        L[i] = (P[i-1] + A[i-1] + P[i-2] + A[i-2]) % MOD
        if i > 3 {
            A[i] = (A[i-1] + A[i-2] + A[i-3]) % MOD
        }
    }
    return (P[n] + L[n] + A[n]) % MOD
}