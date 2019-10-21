func subarraysWithKDistinct(A []int, K int) int {
    return subarraysMostKDistinct(A, K) - subarraysMostKDistinct(A, K-1)
}

// 340:Longest Substring with At Most K Distinct Characters 
func subarraysMostKDistinct(A []int, K int) int {
	// O(n) + O(len(set(A)))
    cnt := make(map[int]int)
    ans := 0
    for i, j := 0, 0; j < len(A); j++ {
        if cnt[A[j]] == 0 {
            K--
        }
        cnt[A[j]]++
        for ; K < 0; i++ {
            cnt[A[i]]--
            if cnt[A[i]] == 0 {
                K++
            }
        }
        ans += j - i + 1
    }
    return ans
}