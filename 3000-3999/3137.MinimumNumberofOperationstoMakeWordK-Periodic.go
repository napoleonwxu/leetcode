func minimumOperationsToMakeKPeriodic(word string, k int) int {
    n := len(word)
    cnt := make(map[string]int, n/k)
    ma := 0
    for i := 0; i < n; i += k {
        cnt[word[i:i+k]]++
        ma = max(ma, cnt[word[i:i+k]])
    }
    return n/k - ma
}