func lenLongestFibSubseq(A []int) int {
    n := len(A)
    ans := 0
    // O(n^2) + O(n^2)
    idx := make(map[int]int, n)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for k := 0; k < n; k++ {
        idx[A[k]] = k
        for j := 0; j <k; j++ {
            i, ok := idx[A[k]-A[j]]
            if ok && i < j {
                if dp[i][j] >= 3 {
                    dp[j][k] = dp[i][j] + 1
                } else {
                    dp[j][k] = 3
                }
                if dp[j][k] > ans {
                    ans = dp[j][k]
                }
            }
        }
    }
    /* O(n^2*lg(max(A))) + O(n)
    Map := make(map[int]bool, n)
    for _, num := range A {
        Map[num] = true
    }
    for i := 0; i < n-2; i++ {
        for j := i+1; j < n-1; j++ {
            a, b := A[i], A[j]
            cnt := 2
            for Map[a+b] {
                a, b = b, a+b
                cnt++
            }
            if cnt >= 3 && cnt > ans {
                ans = cnt
            }
        }
    } */
    return ans
}